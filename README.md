
# OpenTelemetry Collector Azure Service Bus Receiver

The `azureservicebusreceiver` enables the collection of metrics from Azure Service Bus. For details about which metrics that are produced, see [documentation](./documentation.md).

## Configuration

Using default credentials:
```yaml
azureservicebus:
  namespace_fqdn: yournamespace.servicebus.windows.net
  collection_interval: 1m
  auth: default_credentials
```

Using managed identity:
```yaml
azureservicebus:
  namespace_fqdn: yournamespace.servicebus.windows.net
  collection_interval: 1m
  auth: managed_identity
  client_id: ${env:CLIENT_ID}
```

Using service principal authentication:
```yaml
azureservicebus:
  namespace_fqdn: yournamespace.servicebus.windows.net
  collection_interval: 1m
  auth: service_principal
  tenant_id: ${env:TENANT_ID}
  client_id: ${env:CLIENT_ID}
  client_secret: ${env:CLIENT_SECRET}
```

Using connection string (not recommended):
```yaml
azureservicebus:
  namespace_fqdn: yournamespace.servicebus.windows.net
  collection_interval: 1m
  auth: connection_string
  connection_string: ${env:CONNECTION_STRING}
```


## Contributing
### Developing
Clone the repository\
Make changes to the receiver\
Run `make build-local` to build a local collector to ./collector\
TBC..