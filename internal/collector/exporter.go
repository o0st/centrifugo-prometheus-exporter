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
	httpServerAPISubsystem = "http_server_api"
	nodeSubsystem          = "node"
)

type Exporter struct {
	nodeName          string
	client            *centrifugo.Client
	up                *prometheus.Desc
	clientAPIRequests *prometheus.Desc
	clientAPIBytes    *prometheus.Desc
	clientAPIMessages *prometheus.Desc
	httpAPIRequests   *prometheus.Desc
	nodeChannels      *prometheus.Desc
	nodeClients       *prometheus.Desc
	nodeUniqueClients *prometheus.Desc
	nodeMemUsage      *prometheus.Desc
	nodeCPUUsage      *prometheus.Desc
	nodeUptime        *prometheus.Desc
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
		clientAPIMessages: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, clientAPISubsystem, "messages"),
			"Number client messages by state.",
			[]string{"state"},
			nil,
		),
		httpAPIRequests: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, httpServerAPISubsystem, "requests"),
			"Number of server http api requests.",
			nil,
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
		nodeMemUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, nodeSubsystem, "mem_usage"),
			"Node memory usage in in bytes.",
			nil,
			nil,
		),
		nodeCPUUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, nodeSubsystem, "cpu_usage"),
			"Node CPU usage in in percents.",
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
	ch <- e.clientAPIMessages
	ch <- e.httpAPIRequests
	ch <- e.nodeChannels
	ch <- e.nodeClients
	ch <- e.nodeUniqueClients
	ch <- e.nodeMemUsage
	ch <- e.nodeCPUUsage
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
	ch <- prometheus.MustNewConstMetric(e.clientAPIRequests, prometheus.GaugeValue, node.ClientAPINumRequests)
	ch <- prometheus.MustNewConstMetric(e.clientAPIBytes, prometheus.GaugeValue, node.ClientBytesIn, "in")
	ch <- prometheus.MustNewConstMetric(e.clientAPIBytes, prometheus.GaugeValue, node.ClientBytesOut, "out")
	ch <- prometheus.MustNewConstMetric(e.clientAPIMessages, prometheus.GaugeValue, node.ClientNumMsgPublished, "published")
	ch <- prometheus.MustNewConstMetric(e.clientAPIMessages, prometheus.GaugeValue, node.ClientNumMsgSent, "sent")
	ch <- prometheus.MustNewConstMetric(e.clientAPIMessages, prometheus.GaugeValue, node.ClientNumMsgQueued, "queued")
	ch <- prometheus.MustNewConstMetric(e.httpAPIRequests, prometheus.GaugeValue, node.HTTPAPINumRequests)
	ch <- prometheus.MustNewConstMetric(e.nodeChannels, prometheus.GaugeValue, node.NodeNumChannels)
	ch <- prometheus.MustNewConstMetric(e.nodeClients, prometheus.GaugeValue, node.NodeNumClients)
	ch <- prometheus.MustNewConstMetric(e.nodeUniqueClients, prometheus.GaugeValue, node.NodeNumUniqueClients)
	ch <- prometheus.MustNewConstMetric(e.nodeMemUsage, prometheus.GaugeValue, node.NodeMemUsage)
	ch <- prometheus.MustNewConstMetric(e.nodeCPUUsage, prometheus.GaugeValue, node.NodeCPUUsage)
	ch <- prometheus.MustNewConstMetric(e.nodeUptime, prometheus.GaugeValue, float64(time.Now().Unix()-node.StartedAt))
}
