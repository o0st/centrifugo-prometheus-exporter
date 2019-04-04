# Prometheus Exporter for Centrifugo (v1)
A [centrifugo](https://github.com/centrifugal/centrifugo/tree/v1) exporter for prometheus.

## Configuration flags

```bash
$ centrifugo-prometheus-exporter -h

Usage:
  centrifugo-prometheus-exporter [flags]

Flags:
      --address string                address to listen on for web interface and telemetry (default ":9564")
      --centrifugo-api-key string     centrifugo api key (or use env CENTRIFUGO_API_KEY)
      --centrifugo-endpoint string    centrifugo server endpoint (default "http://localhost")
      --centrifugo-node-name string   target centrifugo node name
  -h, --help                          help for centrifugo-prometheus-exporter
      --metrics-path string           path under which to expose metrics (default "/metrics")
      
```
## Run in container
```bash
$ docker run -p 9564:9564 kismia/centrifugo-prometheus-exporter
```

## Collectors

```text
# HELP centrifugo_client_api_bytes Number of bytes coming to/from client api.
# TYPE centrifugo_client_api_bytes gauge
centrifugo_client_api_bytes{direction="in"} 2.64127e+06
centrifugo_client_api_bytes{direction="out"} 6.707046e+06
# HELP centrifugo_client_api_connections Number of connection to client api.
# TYPE centrifugo_client_api_connections gauge
centrifugo_client_api_connections 4127
# HELP centrifugo_client_api_messages Number client messages by state.
# TYPE centrifugo_client_api_messages gauge
centrifugo_client_api_messages{state="published"} 0
centrifugo_client_api_messages{state="queued"} 20838
centrifugo_client_api_messages{state="sent"} 20834
# HELP centrifugo_client_api_requests Number of requests to client api.
# TYPE centrifugo_client_api_requests gauge
centrifugo_client_api_requests 11741
# HELP centrifugo_client_api_subscriptions Number of client api subscriptions.
# TYPE centrifugo_client_api_subscriptions gauge
centrifugo_client_api_subscriptions 20500
# HELP centrifugo_http_server_api_requests Number of server http api requests.
# TYPE centrifugo_http_server_api_requests gauge
centrifugo_http_server_api_requests 11
# HELP centrifugo_http_server_requests Number of server http requests by type.
# TYPE centrifugo_http_server_requests gauge
centrifugo_http_server_requests{type="sockjs"} 4194
centrifugo_http_server_requests{type="ws"} 0
# HELP centrifugo_node_channels Number of node channels.
# TYPE centrifugo_node_channels gauge
centrifugo_node_channels 7396
# HELP centrifugo_node_clients Number of node clients.
# TYPE centrifugo_node_clients gauge
centrifugo_node_clients 2920
# HELP centrifugo_node_clients_unique Number of node unique clients.
# TYPE centrifugo_node_clients_unique gauge
centrifugo_node_clients_unique 2506
# HELP centrifugo_node_history_items Number of node history items.
# TYPE centrifugo_node_history_items gauge
centrifugo_node_history_items 996
# HELP centrifugo_node_messages Number of node messages by type and state.
# TYPE centrifugo_node_messages gauge
centrifugo_node_messages{state="",type=""} 20838
centrifugo_node_messages{state="published",type="admin"} 0
centrifugo_node_messages{state="published",type="client"} 19821
centrifugo_node_messages{state="published",type="control"} 405
centrifugo_node_messages{state="published",type="join"} 0
centrifugo_node_messages{state="published",type="leave"} 0
centrifugo_node_messages{state="received",type="admin"} 0
centrifugo_node_messages{state="received",type="client"} 7151
centrifugo_node_messages{state="received",type="control"} 595
centrifugo_node_messages{state="received",type="join"} 0
centrifugo_node_messages{state="received",type="leave"} 0
# HELP centrifugo_node_operations Number of node operations by type and entity.
# TYPE centrifugo_node_operations gauge
centrifugo_node_operations{entity="conn",type="add"} 4127
centrifugo_node_operations{entity="conn",type="remove"} 4222
centrifugo_node_operations{entity="presence",type="add"} 19123
centrifugo_node_operations{entity="presence",type="remove"} 20929
centrifugo_node_operations{entity="sub",type="add"} 20500
centrifugo_node_operations{entity="sub",type="remove"} 20929
# HELP centrifugo_node_presence_items Number of node presence items.
# TYPE centrifugo_node_presence_items gauge
centrifugo_node_presence_items 10
# HELP centrifugo_node_uptime_seconds Node uptime in seconds.
# TYPE centrifugo_node_uptime_seconds gauge
centrifugo_node_uptime_seconds 2.1273767e+07
# HELP centrifugo_up Could the centrifugo server be reached.
# TYPE centrifugo_up gauge
centrifugo_up 1
```
