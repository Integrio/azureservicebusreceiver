// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					ServicebusQueueActiveMessages:                 MetricConfig{Enabled: true},
					ServicebusQueueCurrentSize:                    MetricConfig{Enabled: true},
					ServicebusQueueDeadletterMessages:             MetricConfig{Enabled: true},
					ServicebusQueueMaxSize:                        MetricConfig{Enabled: true},
					ServicebusQueueScheduledMessages:              MetricConfig{Enabled: true},
					ServicebusTopicCurrentSize:                    MetricConfig{Enabled: true},
					ServicebusTopicMaxSize:                        MetricConfig{Enabled: true},
					ServicebusTopicScheduledMessages:              MetricConfig{Enabled: true},
					ServicebusTopicSubscriptionActiveMessages:     MetricConfig{Enabled: true},
					ServicebusTopicSubscriptionDeadletterMessages: MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ServicebusNamespaceName: ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					ServicebusQueueActiveMessages:                 MetricConfig{Enabled: false},
					ServicebusQueueCurrentSize:                    MetricConfig{Enabled: false},
					ServicebusQueueDeadletterMessages:             MetricConfig{Enabled: false},
					ServicebusQueueMaxSize:                        MetricConfig{Enabled: false},
					ServicebusQueueScheduledMessages:              MetricConfig{Enabled: false},
					ServicebusTopicCurrentSize:                    MetricConfig{Enabled: false},
					ServicebusTopicMaxSize:                        MetricConfig{Enabled: false},
					ServicebusTopicScheduledMessages:              MetricConfig{Enabled: false},
					ServicebusTopicSubscriptionActiveMessages:     MetricConfig{Enabled: false},
					ServicebusTopicSubscriptionDeadletterMessages: MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ServicebusNamespaceName: ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{}))
			require.Emptyf(t, diff, "Config mismatch (-expected +actual):\n%s", diff)
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				ServicebusNamespaceName: ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				ServicebusNamespaceName: ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{}))
			require.Emptyf(t, diff, "Config mismatch (-expected +actual):\n%s", diff)
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}
