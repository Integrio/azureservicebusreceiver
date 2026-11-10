module github.com/Integrio/azureservicebusreceiver

go 1.26.0

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.24.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.15.0
	github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus v1.11.0
	github.com/google/go-cmp v0.7.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/stretchr/testify v1.12.1
	go.opentelemetry.io/collector/component v1.71.0
	go.opentelemetry.io/collector/component/componenttest v0.165.0
	go.opentelemetry.io/collector/confmap v1.71.0
	go.opentelemetry.io/collector/consumer v1.71.0
	go.opentelemetry.io/collector/consumer/consumertest v0.165.0
	go.opentelemetry.io/collector/filter v0.165.0
	go.opentelemetry.io/collector/pdata v1.71.0
	go.opentelemetry.io/collector/pdata/xpdata v0.165.0
	go.opentelemetry.io/collector/receiver v1.71.0
	go.opentelemetry.io/collector/receiver/receivertest v0.165.0
	go.opentelemetry.io/collector/scraper v0.165.0
	go.opentelemetry.io/collector/scraper/scraperhelper v0.165.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.28.0
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.13.0 // indirect
	github.com/Azure/go-amqp v1.7.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.10.1 // indirect
	github.com/cenkalti/backoff/v7 v7.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.5.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/jsonschema-go v0.4.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-version v1.9.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/iimos/ucum v0.0.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.3 // indirect
	github.com/knadh/koanf/providers/confmap v1.0.1 // indirect
	github.com/knadh/koanf/v2 v2.4.1 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/collector/client v1.71.0 // indirect
	go.opentelemetry.io/collector/cmd/mdatagen v0.165.0 // indirect
	go.opentelemetry.io/collector/config/configoptional v1.71.0 // indirect
	go.opentelemetry.io/collector/config/configretry v1.71.0 // indirect
	go.opentelemetry.io/collector/confmap/provider/fileprovider v1.71.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.165.0 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.165.0 // indirect
	go.opentelemetry.io/collector/exporter v1.71.0 // indirect
	go.opentelemetry.io/collector/exporter/exporterhelper v0.165.0 // indirect
	go.opentelemetry.io/collector/extension v1.71.0 // indirect
	go.opentelemetry.io/collector/extension/xextension v0.165.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.71.0 // indirect
	go.opentelemetry.io/collector/internal/componentalias v0.165.0 // indirect
	go.opentelemetry.io/collector/internal/schemagen v0.165.0 // indirect
	go.opentelemetry.io/collector/internal/telemetry v0.165.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.165.0 // indirect
	go.opentelemetry.io/collector/pipeline v1.71.0 // indirect
	go.opentelemetry.io/collector/pipeline/xpipeline v0.165.0 // indirect
	go.opentelemetry.io/collector/receiver/receiverhelper v0.165.0 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.165.0 // indirect
	go.opentelemetry.io/otel v1.47.0 // indirect
	go.opentelemetry.io/otel/log v1.47.0 // indirect
	go.opentelemetry.io/otel/metric v1.47.0 // indirect
	go.opentelemetry.io/otel/sdk v1.47.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.47.0 // indirect
	go.opentelemetry.io/otel/trace v1.47.0 // indirect
	go.yaml.in/yaml/v3 v3.0.6 // indirect
	golang.org/x/crypto v0.58.0 // indirect
	golang.org/x/mod v0.42.0 // indirect
	golang.org/x/net v0.62.0 // indirect
	golang.org/x/sync v0.24.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/text v0.43.0 // indirect
	golang.org/x/tools v0.52.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260928230214-8a89bd6388cc // indirect
	google.golang.org/grpc v1.85.0 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)

tool go.opentelemetry.io/collector/cmd/mdatagen
