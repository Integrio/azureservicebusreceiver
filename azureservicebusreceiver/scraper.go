package azureservicebusreceiver

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/admin"
	"github.com/Integrio/azureservicebusreceiver/azureservicebusreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"
	"time"
)

type serviceBusScraper struct {
	client   *admin.Client
	logger   *zap.Logger
	cfg      *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
}

// newScraper creates a new scraper
func newScraper(logger *zap.Logger, cfg *Config, settings receiver.Settings) *serviceBusScraper {
	return &serviceBusScraper{
		logger:   logger,
		cfg:      cfg,
		settings: settings.TelemetrySettings,
		mb:       metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings),
	}
}

// start starts the scraper by creating a new admin client on the scraper
func (s *serviceBusScraper) start(_ context.Context, _ component.Host) (err error) {
	if s.cfg.ConnectionString != "" {
		s.logger.Debug("Using connection string to create client")
		s.client, err = admin.NewClientFromConnectionString(s.cfg.ConnectionString, nil)
	} else {
		s.logger.Debug("Using managed identity to create client")
		credential, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			return err
		}

		s.client, err = admin.NewClient(s.cfg.NamespaceFqdn, credential, nil)
	}

	return
}

func (s *serviceBusScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	if s.client == nil {
		return pmetric.NewMetrics(), errors.New("client not initialized")
	}

	scrapeErrors := scrapererror.ScrapeErrors{}

	rb := s.mb.NewResourceBuilder()
	rb.SetServicebusNamespaceName(s.cfg.NamespaceFqdn)

	s.scrapeQueues(ctx, scrapeErrors)
	topics := s.scrapeTopics(ctx, scrapeErrors)
	s.scrapeSubscriptions(ctx, topics, scrapeErrors)

	return s.mb.Emit(metadata.WithResource(rb.Emit())), scrapeErrors.Combine()
}

func (s *serviceBusScraper) scrapeQueues(ctx context.Context, errors scrapererror.ScrapeErrors) {
	pager := s.client.NewListQueuesRuntimePropertiesPager(nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			errors.AddPartial(1, err)
			continue
		}

		now := pcommon.NewTimestampFromTime(time.Now())
		for _, queue := range page.QueueRuntimeProperties {
			queueName := queue.QueueName
			s.mb.RecordServicebusQueueCurrentSizeDataPoint(now, queue.SizeInBytes, queueName)
			s.mb.RecordServicebusQueueScheduledMessagesDataPoint(now, int64(queue.ScheduledMessageCount), queueName)
			s.mb.RecordServicebusQueueActiveMessagesDataPoint(now, int64(queue.ActiveMessageCount), queueName)
			s.mb.RecordServicebusQueueDeadletterMessagesDataPoint(now, int64(queue.DeadLetterMessageCount), queueName)
		}

	}
}

func (s *serviceBusScraper) scrapeTopics(ctx context.Context, errors scrapererror.ScrapeErrors) []string {
	pager := s.client.NewListTopicsRuntimePropertiesPager(nil)
	var topics []string

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			errors.AddPartial(1, err)
			continue
		}

		now := pcommon.NewTimestampFromTime(time.Now())
		for _, topic := range page.TopicRuntimeProperties {
			topicName := topic.TopicName
			s.mb.RecordServicebusTopicCurrentSizeDataPoint(now, topic.SizeInBytes, topicName)
			s.mb.RecordServicebusTopicScheduledMessagesDataPoint(now, int64(topic.ScheduledMessageCount), topicName)
			topics = append(topics, topicName)
		}
	}

	return topics
}

func (s *serviceBusScraper) scrapeSubscriptions(ctx context.Context, topics []string, errors scrapererror.ScrapeErrors) {
	for _, topic := range topics {
		pager := s.client.NewListSubscriptionsRuntimePropertiesPager(topic, nil)

		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				errors.AddPartial(1, err)
				continue
			}

			now := pcommon.NewTimestampFromTime(time.Now())
			for _, sub := range page.SubscriptionRuntimeProperties {
				subName := sub.SubscriptionName
				s.mb.RecordServicebusTopicSubscriptionActiveMessagesDataPoint(now, int64(sub.ActiveMessageCount), topic, subName)
				s.mb.RecordServicebusTopicSubscriptionDeadletterMessagesDataPoint(now, int64(sub.ActiveMessageCount), topic, subName)
			}
		}
	}
}
