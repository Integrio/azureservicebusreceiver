package azureservicebusreceiver

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/admin"
	"github.com/Integrio/azureservicebusreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"
	"time"
)

const (
	bytesPerMegabyte int64 = 1024 * 1024
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
	if s.cfg.Authentication == connectionString {
		s.client, err = admin.NewClientFromConnectionString(s.cfg.ConnectionString, nil)
	} else {
		cred, err := getAzCredential(s.cfg)
		if err != nil {
			return err
		}
		s.client, err = admin.NewClient(s.cfg.NamespaceFqdn, cred, nil)
	}

	return nil
}

func getAzCredential(cfg *Config) (azcore.TokenCredential, error) {
	var cred azcore.TokenCredential
	var err error

	switch cfg.Authentication {
	case defaultCredentials:
		if cred, err = azidentity.NewDefaultAzureCredential(nil); err != nil {
			return nil, err
		}
	case servicePrincipal:
		if cred, err = azidentity.NewClientSecretCredential(cfg.TenantId, cfg.ClientId, cfg.ClientSecret, nil); err != nil {
			return nil, err
		}
	case managedIdentity:
		var options *azidentity.ManagedIdentityCredentialOptions
		if cfg.ClientId != "" {
			options = &azidentity.ManagedIdentityCredentialOptions{
				ID: azidentity.ClientID(cfg.ClientId),
			}
		}
		if cred, err = azidentity.NewManagedIdentityCredential(options); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown authentication %v", cfg.Authentication)
	}

	return cred, nil
}

func (s *serviceBusScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	if s.client == nil {
		return pmetric.NewMetrics(), errors.New("client not initialized")
	}

	scrapeErrors := scrapererror.ScrapeErrors{}

	now := pcommon.NewTimestampFromTime(time.Now())

	s.scrapeQueues(ctx, &scrapeErrors, now)
	topics := s.scrapeTopics(ctx, &scrapeErrors, now)
	s.scrapeSubscriptions(ctx, topics, &scrapeErrors, now)

	rb := s.mb.NewResourceBuilder()
	rb.SetServicebusNamespaceName(s.cfg.NamespaceFqdn)
	return s.mb.Emit(metadata.WithResource(rb.Emit())), scrapeErrors.Combine()
}

func (s *serviceBusScraper) scrapeQueues(
	ctx context.Context,
	errors *scrapererror.ScrapeErrors,
	now pcommon.Timestamp,
) {
	pager := s.client.NewListQueuesRuntimePropertiesPager(nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			errors.AddPartial(1, err)
			continue
		}

		for _, queue := range page.QueueRuntimeProperties {
			queueName := queue.QueueName
			s.mb.RecordServicebusQueueCurrentSizeDataPoint(now, queue.SizeInBytes, queueName)
			s.mb.RecordServicebusQueueScheduledMessagesDataPoint(now, int64(queue.ScheduledMessageCount), queueName)
			s.mb.RecordServicebusQueueActiveMessagesDataPoint(now, int64(queue.ActiveMessageCount), queueName)
			s.mb.RecordServicebusQueueDeadletterMessagesDataPoint(now, int64(queue.DeadLetterMessageCount), queueName)
			s.scrapeQueueProperties(ctx, queueName, errors, now)
		}

	}
}

func (s *serviceBusScraper) scrapeQueueProperties(
	ctx context.Context,
	queue string,
	errors *scrapererror.ScrapeErrors,
	now pcommon.Timestamp,
) {
	queueProps, err := s.client.GetQueue(ctx, queue, nil)
	if err != nil {
		errors.AddPartial(1, err)
		return
	}

	maxSize := int64(*queueProps.MaxSizeInMegabytes) * bytesPerMegabyte
	s.mb.RecordServicebusQueueMaxSizeDataPoint(now, maxSize, queue)
}

func (s *serviceBusScraper) scrapeTopics(
	ctx context.Context,
	errors *scrapererror.ScrapeErrors,
	now pcommon.Timestamp,
) []string {
	pager := s.client.NewListTopicsRuntimePropertiesPager(nil)
	var topics []string

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			errors.AddPartial(1, err)
			continue
		}

		for _, topic := range page.TopicRuntimeProperties {
			topicName := topic.TopicName
			s.mb.RecordServicebusTopicCurrentSizeDataPoint(now, topic.SizeInBytes, topicName)
			s.mb.RecordServicebusTopicScheduledMessagesDataPoint(now, int64(topic.ScheduledMessageCount), topicName)
			s.scrapeTopicProperties(ctx, topicName, errors, now)
			topics = append(topics, topicName)
		}
	}

	return topics
}

func (s *serviceBusScraper) scrapeTopicProperties(
	ctx context.Context,
	topic string,
	errors *scrapererror.ScrapeErrors,
	now pcommon.Timestamp,
) {
	topicProps, err := s.client.GetTopic(ctx, topic, nil)
	if err != nil {
		errors.AddPartial(1, err)
		return
	}

	maxSize := int64(*topicProps.MaxSizeInMegabytes) * bytesPerMegabyte
	s.mb.RecordServicebusTopicMaxSizeDataPoint(now, maxSize, topic)
}

func (s *serviceBusScraper) scrapeSubscriptions(
	ctx context.Context,
	topics []string,
	errors *scrapererror.ScrapeErrors,
	now pcommon.Timestamp,
) {
	for _, topic := range topics {
		pager := s.client.NewListSubscriptionsRuntimePropertiesPager(topic, nil)

		for pager.More() {
			page, err := pager.NextPage(ctx)
			if err != nil {
				errors.AddPartial(1, err)
				continue
			}

			for _, sub := range page.SubscriptionRuntimeProperties {
				subName := sub.SubscriptionName
				s.mb.RecordServicebusTopicSubscriptionActiveMessagesDataPoint(now, int64(sub.ActiveMessageCount), topic, subName)
				s.mb.RecordServicebusTopicSubscriptionDeadletterMessagesDataPoint(now, int64(sub.DeadLetterMessageCount), topic, subName)
			}
		}
	}
}
