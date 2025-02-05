package azureservicebusreceiver

import (
	"context"
	"errors"
	"github.com/Integrio/azureservicebusreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"time"
)

var (
	typeStr = component.MustNewType("azureservicebus")
)

const (
	defaultInterval = 5 * time.Minute
)

var errInvalidConfig = errors.New("config was not an azure servicebus receiver config")

// NewFactory creates a factory for azureservicebus receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability))
}

func createDefaultConfig() component.Config {
	cfg := scraperhelper.NewDefaultControllerConfig()
	cfg.CollectionInterval = defaultInterval
	cfg.InitialDelay = time.Second

	return &Config{
		ControllerConfig:     cfg,
		MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
	}
}

func createMetricsReceiver(_ context.Context, params receiver.Settings, conf component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	cfg, ok := conf.(*Config)
	if !ok {
		return nil, errInvalidConfig
	}

	serviceBusScraper := newScraper(params.Logger, cfg, params)
	s, err := scraper.NewMetrics(serviceBusScraper.scrape, scraper.WithStart(serviceBusScraper.start))
	if err != nil {
		return nil, err
	}

	return scraperhelper.NewMetricsController(&cfg.ControllerConfig, params, consumer, scraperhelper.AddScraper(metadata.Type, s))
}
