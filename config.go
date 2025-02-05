package azureservicebusreceiver

import (
	"errors"
	"fmt"
	"github.com/Integrio/azureservicebusreceiver/internal/metadata"
	"github.com/hashicorp/go-multierror"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"go.uber.org/multierr"
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	NamespaceFqdn                  string `mapstructure:"namespace_fqdn"`
	Authentication                 string `mapstructure:"auth"`
	ConnectionString               string `mapstructure:"connection_string"`
	TenantId                       string `mapstructure:"tenant_id"`
	ClientId                       string `mapstructure:"client_id"`
	ClientSecret                   string `mapstructure:"client_secret"`
}

const (
	defaultCredentials = "default_credentials"
	servicePrincipal   = "service_principal"
	managedIdentity    = "managed_identity"
	connectionString   = "connection_string"
)

func (cfg *Config) Validate() (err error) {
	if cfg.NamespaceFqdn == "" {
		err = multierror.Append(err, errors.New("NamespaceFqdn is required"))
	}

	switch cfg.Authentication {
	case servicePrincipal:
		if cfg.TenantId == "" {
			err = multierr.Append(err, errors.New("TenantId is required"))
		}

		if cfg.ClientId == "" {
			err = multierr.Append(err, errors.New("ClientId is required"))
		}

		if cfg.ClientSecret == "" {
			err = multierr.Append(err, errors.New("ClientSecret is required"))
		}
	case connectionString:
		if cfg.ConnectionString == "" {
			err = multierror.Append(err, errors.New("ConnectionString is required"))
		}

	case managedIdentity:
	case defaultCredentials:
	default:
		return fmt.Errorf("authentication %v is not supported. supported authentications include [%v,%v,%v,%v]", cfg.Authentication, servicePrincipal, managedIdentity, defaultCredentials, connectionString)
	}

	return nil
}
