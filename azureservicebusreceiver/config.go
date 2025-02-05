package azureservicebusreceiver

import (
	"errors"
	"github.com/Integrio/azureservicebusreceiver/azureservicebusreceiver/internal/metadata"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	ConnectionString               string `mapstructure:"connection_string"`
	NamespaceFqdn                  string `mapstructure:"namespace_fqdn"`
}

func (cfg *Config) Validate() error {
	if cfg.NamespaceFqdn == "" {
		return errors.New("namespace fqdn is required")
	}

	return nil
}
