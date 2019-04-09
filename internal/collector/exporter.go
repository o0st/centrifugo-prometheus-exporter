package collector

import (
	"context"
	"time"

	"github.com/kismia/centrifugo-prometheus-exporter/internal/pkg/centrifugo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

const (
	namespace              = "centrifugo"
	clientAPISubsystem     = "client_api"
	httpServerSubsystem    = "http_server"
	httpServerAPISubsystem = "http_server_api"
	nodeSubsystem          = "node"
)

type Exporter struct {
	nodeName               string
	client                 *centrifugo.Client
	up                     *prometheus.Desc
	clientAPIRequests      *prometheus.Desc
	clientAPIBytes         *prometheus.Desc
	clientAPIConnections   *prometheus.Desc
	clientAPIMessages      *prometheus.Desc
	clientAPISubscriptions *prometheus.Desc
	httpAPIRequests        *prometheus.Desc
	httpRequests           *prometheus.Desc
	nodeMessages           *prometheus.Desc
	nodeOperations         *prometheus.Desc
	nodeChannels           *prometheus.Desc
	nodeClients            *prometheus.Desc
	nodeUniqueClients      *prometheus.Desc
	nodeHistoryItems       *prometheus.Desc
	nodePresenceItems      *prometheus.Desc
	nodeUptime             *prometheus.Desc
}

func NewExporter(client *centrifugo.Client, nodeName string) *Exporter {
	return &Exporter{
		nodeName: nodeName,
		client:   client,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "up"),
			"Could the centrifugo server be reached.",
			nil,
			nil,
		),
		clientAPIRequests: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, clientAPISubsystem, "requests"),
			"Number of requests to client api.",
			nil,
			nil,
		),
		clientAPIBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, clientAPISubsystem, "bytes"),
			"Number of bytes coming to/from client api.",
			[]string{"direction"},
			nil,
		),
		clientAPIConnections: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, clientAPISubsystem, "connections"),
			"Number of connection to client api.",
			nil,
			nil,
		),
		clientAPIMessages: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, clientAPISubsystem, "messages"),
			"Number client messages by state.",
			[]string{"state"},
			nil,
		),
		clientAPISubscriptions: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, clientAPISubsystem, "subscriptions"),
			"Number of client api subscriptions.",
			nil,
			nil,
		),
		httpAPIRequests: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, httpServerAPISubsystem, "requests"),
			"Number of server http api requests.",
			nil,
			nil,
		),
		httpRequests: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, httpServerSubsystem, "requests"),
			"Number of server http requests by type.",
			[]string{"type"},
			nil,
		),
		nodeMessages: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, nodeSubsystem, "messages"),
			"Number of node messages by type and state.",
			[]string{"type", "state"},
			nil,
		),
		nodeOperations: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, nodeSubsystem, "operations"),
			"Number of node operations by type and entity.",
			[]string{"type", "entity"},
			nil,
		),
		nodeChannels: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, nodeSubsystem, "channels"),
			"Number of node channels.",
			nil,
			nil,
		),
		nodeClients: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, nodeSubsystem, "clients"),
			"Number of node clients.",
			nil,
			nil,
		),
		nodeUniqueClients: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, nodeSubsystem, "clients_unique"),
			"Number of node unique clients.",
			nil,
			nil,
		),
		nodeHistoryItems: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, nodeSubsystem, "history_items"),
			"Number of node history items.",
			nil,
			nil,
		),
		nodePresenceItems: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, nodeSubsystem, "presence_items"),
			"Number of node presence items.",
			nil,
			nil,
		),
		nodeUptime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, nodeSubsystem, "uptime_seconds"),
			"Node uptime in seconds.",
			nil,
			nil,
		),
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.up
	ch <- e.clientAPIRequests
	ch <- e.clientAPIBytes
	ch <- e.clientAPIConnections
	ch <- e.clientAPIMessages
	ch <- e.clientAPISubscriptions
	ch <- e.httpAPIRequests
	ch <- e.httpRequests
	ch <- e.nodeMessages
	ch <- e.nodeOperations
	ch <- e.nodeClients
	ch <- e.nodeUniqueClients
	ch <- e.nodeChannels
	ch <- e.nodeHistoryItems
	ch <- e.nodePresenceItems
	ch <- e.nodeUptime
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	response, err := e.client.GetStats(context.Background())

	if err != nil {
		ch <- prometheus.MustNewConstMetric(e.up, prometheus.GaugeValue, 0)

		logrus.Error(err)

		return
	}

	for _, node := range response.Body.Data.Nodes {
		if node.Name == e.nodeName {
			e.collectFromNode(ch, node)

			return
		}
	}

	logrus.Errorf("node name '%s' not found in server response", e.nodeName)
}

func (e *Exporter) collectFromNode(ch chan<- prometheus.Metric, node centrifugo.Node) {
	ch <- prometheus.MustNewConstMetric(e.up, prometheus.GaugeValue, 1)
	ch <- prometheus.MustNewConstMetric(e.clientAPIRequests, prometheus.GaugeValue, node.Metrics.ClientAPINumRequests)
	ch <- prometheus.MustNewConstMetric(e.clientAPIBytes, prometheus.GaugeValue, node.Metrics.ClientBytesIn, "in")
	ch <- prometheus.MustNewConstMetric(e.clientAPIBytes, prometheus.GaugeValue, node.Metrics.ClientBytesOut, "out")
	ch <- prometheus.MustNewConstMetric(e.clientAPIConnections, prometheus.GaugeValue, node.Metrics.ClientNumConnect)
	ch <- prometheus.MustNewConstMetric(e.clientAPIMessages, prometheus.GaugeValue, node.Metrics.ClientNumMsgPublished, "published")
	ch <- prometheus.MustNewConstMetric(e.clientAPIMessages, prometheus.GaugeValue, node.Metrics.ClientNumMsgSent, "sent")
	ch <- prometheus.MustNewConstMetric(e.clientAPIMessages, prometheus.GaugeValue, node.Metrics.ClientNumMsgQueued, "queued")
	ch <- prometheus.MustNewConstMetric(e.clientAPISubscriptions, prometheus.GaugeValue, node.Metrics.ClientNumSubscribe)
	ch <- prometheus.MustNewConstMetric(e.httpAPIRequests, prometheus.GaugeValue, node.Metrics.HTTPAPINumRequests)
	ch <- prometheus.MustNewConstMetric(e.httpRequests, prometheus.GaugeValue, node.Metrics.HTTPRawWsNumRequests, "ws")
	ch <- prometheus.MustNewConstMetric(e.httpRequests, prometheus.GaugeValue, node.Metrics.HTTPSockjsNumRequests, "sockjs")
	ch <- prometheus.MustNewConstMetric(e.nodeMessages, prometheus.GaugeValue, node.Metrics.NodeNumClientMsgPublished, "client", "published")
	ch <- prometheus.MustNewConstMetric(e.nodeMessages, prometheus.GaugeValue, node.Metrics.NodeNumClientMsgReceived, "client", "received")
	ch <- prometheus.MustNewConstMetric(e.nodeMessages, prometheus.GaugeValue, node.Metrics.NodeNumControlMsgPublished, "control", "published")
	ch <- prometheus.MustNewConstMetric(e.nodeMessages, prometheus.GaugeValue, node.Metrics.NodeNumControlMsgReceived, "control", "received")
	ch <- prometheus.MustNewConstMetric(e.nodeMessages, prometheus.GaugeValue, node.Metrics.NodeNumJoinMsgPublished, "join", "published")
	ch <- prometheus.MustNewConstMetric(e.nodeMessages, prometheus.GaugeValue, node.Metrics.NodeNumJoinMsgReceived, "join", "received")
	ch <- prometheus.MustNewConstMetric(e.nodeMessages, prometheus.GaugeValue, node.Metrics.NodeNumLeaveMsgPublished, "leave", "published")
	ch <- prometheus.MustNewConstMetric(e.nodeMessages, prometheus.GaugeValue, node.Metrics.NodeNumLeaveMsgReceived, "leave", "received")
	ch <- prometheus.MustNewConstMetric(e.nodeMessages, prometheus.GaugeValue, node.Metrics.NodeNumAdminMsgPublished, "admin", "published")
	ch <- prometheus.MustNewConstMetric(e.nodeMessages, prometheus.GaugeValue, node.Metrics.NodeNumAdminMsgReceived, "admin", "received")
	ch <- prometheus.MustNewConstMetric(e.nodeOperations, prometheus.GaugeValue, node.Metrics.NodeNumAddClientConn, "add", "conn")
	ch <- prometheus.MustNewConstMetric(e.nodeOperations, prometheus.GaugeValue, node.Metrics.NodeNumRemoveClientConn, "remove", "conn")
	ch <- prometheus.MustNewConstMetric(e.nodeOperations, prometheus.GaugeValue, node.Metrics.NodeNumAddClientSub, "add", "sub")
	ch <- prometheus.MustNewConstMetric(e.nodeOperations, prometheus.GaugeValue, node.Metrics.NodeNumRemoveClientSub, "remove", "sub")
	ch <- prometheus.MustNewConstMetric(e.nodeOperations, prometheus.GaugeValue, node.Metrics.NodeNumAddPresence, "add", "presence")
	ch <- prometheus.MustNewConstMetric(e.nodeOperations, prometheus.GaugeValue, node.Metrics.NodeNumRemovePresence, "remove", "presence")
	ch <- prometheus.MustNewConstMetric(e.nodeChannels, prometheus.GaugeValue, node.Metrics.NodeNumChannels)
	ch <- prometheus.MustNewConstMetric(e.nodeClients, prometheus.GaugeValue, node.Metrics.NodeNumClients)
	ch <- prometheus.MustNewConstMetric(e.nodeUniqueClients, prometheus.GaugeValue, node.Metrics.NodeNumUniqueClients)
	ch <- prometheus.MustNewConstMetric(e.nodeHistoryItems, prometheus.GaugeValue, node.Metrics.NodeNumHistory)
	ch <- prometheus.MustNewConstMetric(e.nodePresenceItems, prometheus.GaugeValue, node.Metrics.NodeNumPresence)
	ch <- prometheus.MustNewConstMetric(e.nodeUptime, prometheus.GaugeValue, float64(time.Now().Unix()-node.StartedAt))
}
