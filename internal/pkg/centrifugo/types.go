package centrifugo

type StatsCommand struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

type StatsResponse struct {
	Body Body `json:"body"`
}

type Body struct {
	Data Data `json:"data"`
}

type Data struct {
	Nodes           []Node  `json:"nodes"`
	MetricsInterval float64 `json:"metrics_interval"`
}

type Node struct {
	UID       string `json:"uid"`
	Name      string `json:"name"`
	StartedAt int64  `json:"started_at"`

	// Metrics
	ClientBytesIn         float64 `json:"bytes_client_in"`
	ClientBytesOut        float64 `json:"bytes_client_out"`
	ClientNumMsgPublished float64 `json:"num_msg_published"`
	ClientNumMsgQueued    float64 `json:"num_msg_queued"`
	ClientNumMsgSent      float64 `json:"num_msg_sent"`
	ClientAPINumRequests  float64 `json:"num_client_requests"`

	HTTPAPINumRequests float64 `json:"num_api_requests"`

	NodeNumClients       float64 `json:"num_clients"`
	NodeNumUniqueClients float64 `json:"num_unique_clients"`
	NodeNumChannels      float64 `json:"num_channels"`

	NodeMemUsage float64 `json:"memory_sys"`
	NodeCPUUsage float64 `json:"cpu_usage"`
}
