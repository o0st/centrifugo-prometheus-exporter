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
	UID       string  `json:"uid"`
	Name      string  `json:"name"`
	StartedAt int64   `json:"started_at"`
	Metrics   Metrics `json:"metrics"`
}

type Metrics struct {
	ClientAPI15Count               float64 `json:"client_api_15_count"`
	ClientAPI15Microseconds50Ile   float64 `json:"client_api_15_microseconds_50%ile"`
	ClientAPI15Microseconds90Ile   float64 `json:"client_api_15_microseconds_90%ile"`
	ClientAPI15Microseconds99Ile   float64 `json:"client_api_15_microseconds_99%ile"`
	ClientAPI15Microseconds9999Ile float64 `json:"client_api_15_microseconds_99.99%ile"`
	ClientAPI15MicrosecondsMax     float64 `json:"client_api_15_microseconds_max"`
	ClientAPI15MicrosecondsMean    float64 `json:"client_api_15_microseconds_mean"`
	ClientAPI15MicrosecondsMin     float64 `json:"client_api_15_microseconds_min"`
	ClientAPI1Count                float64 `json:"client_api_1_count"`
	ClientAPI1Microseconds50Ile    float64 `json:"client_api_1_microseconds_50%ile"`
	ClientAPI1Microseconds90Ile    float64 `json:"client_api_1_microseconds_90%ile"`
	ClientAPI1Microseconds99Ile    float64 `json:"client_api_1_microseconds_99%ile"`
	ClientAPI1Microseconds9999Ile  float64 `json:"client_api_1_microseconds_99.99%ile"`
	ClientAPI1MicrosecondsMax      float64 `json:"client_api_1_microseconds_max"`
	ClientAPI1MicrosecondsMean     float64 `json:"client_api_1_microseconds_mean"`
	ClientAPI1MicrosecondsMin      float64 `json:"client_api_1_microseconds_min"`
	ClientAPINumRequests           float64 `json:"client_api_num_requests"`
	ClientBytesIn                  float64 `json:"client_bytes_in"`
	ClientBytesOut                 float64 `json:"client_bytes_out"`
	ClientNumConnect               float64 `json:"client_num_connect"`
	ClientNumMsgPublished          float64 `json:"client_num_msg_published"`
	ClientNumMsgQueued             float64 `json:"client_num_msg_queued"`
	ClientNumMsgSent               float64 `json:"client_num_msg_sent"`
	ClientNumSubscribe             float64 `json:"client_num_subscribe"`
	HTTPAPI15Count                 float64 `json:"http_api_15_count"`
	HTTPAPI15Microseconds50Ile     float64 `json:"http_api_15_microseconds_50%ile"`
	HTTPAPI15Microseconds90Ile     float64 `json:"http_api_15_microseconds_90%ile"`
	HTTPAPI15Microseconds99Ile     float64 `json:"http_api_15_microseconds_99%ile"`
	HTTPAPI15Microseconds9999Ile   float64 `json:"http_api_15_microseconds_99.99%ile"`
	HTTPAPI15MicrosecondsMax       float64 `json:"http_api_15_microseconds_max"`
	HTTPAPI15MicrosecondsMean      float64 `json:"http_api_15_microseconds_mean"`
	HTTPAPI15MicrosecondsMin       float64 `json:"http_api_15_microseconds_min"`
	HTTPAPI1Count                  float64 `json:"http_api_1_count"`
	HTTPAPI1Microseconds50Ile      float64 `json:"http_api_1_microseconds_50%ile"`
	HTTPAPI1Microseconds90Ile      float64 `json:"http_api_1_microseconds_90%ile"`
	HTTPAPI1Microseconds99Ile      float64 `json:"http_api_1_microseconds_99%ile"`
	HTTPAPI1Microseconds9999Ile    float64 `json:"http_api_1_microseconds_99.99%ile"`
	HTTPAPI1MicrosecondsMax        float64 `json:"http_api_1_microseconds_max"`
	HTTPAPI1MicrosecondsMean       float64 `json:"http_api_1_microseconds_mean"`
	HTTPAPI1MicrosecondsMin        float64 `json:"http_api_1_microseconds_min"`
	HTTPAPINumRequests             float64 `json:"http_api_num_requests"`
	HTTPRawWsNumRequests           float64 `json:"http_raw_ws_num_requests"`
	HTTPSockjsNumRequests          float64 `json:"http_sockjs_num_requests"`
	NodeCPUUsage                   float64 `json:"node_cpu_usage"`
	NodeMemorySys                  float64 `json:"node_memory_sys"`
	NodeNumAddClientConn           float64 `json:"node_num_add_client_conn"`
	NodeNumAddClientSub            float64 `json:"node_num_add_client_sub"`
	NodeNumAddPresence             float64 `json:"node_num_add_presence"`
	NodeNumAdminMsgPublished       float64 `json:"node_num_admin_msg_published"`
	NodeNumAdminMsgReceived        float64 `json:"node_num_admin_msg_received"`
	NodeNumChannels                float64 `json:"node_num_channels"`
	NodeNumClientMsgPublished      float64 `json:"node_num_client_msg_published"`
	NodeNumClientMsgReceived       float64 `json:"node_num_client_msg_received"`
	NodeNumClients                 float64 `json:"node_num_clients"`
	NodeNumControlMsgPublished     float64 `json:"node_num_control_msg_published"`
	NodeNumControlMsgReceived      float64 `json:"node_num_control_msg_received"`
	NodeNumGoroutine               float64 `json:"node_num_goroutine"`
	NodeNumHistory                 float64 `json:"node_num_history"`
	NodeNumJoinMsgPublished        float64 `json:"node_num_join_msg_published"`
	NodeNumJoinMsgReceived         float64 `json:"node_num_join_msg_received"`
	NodeNumLastMessageID           float64 `json:"node_num_last_message_id"`
	NodeNumLeaveMsgPublished       float64 `json:"node_num_leave_msg_published"`
	NodeNumLeaveMsgReceived        float64 `json:"node_num_leave_msg_received"`
	NodeNumPresence                float64 `json:"node_num_presence"`
	NodeNumRemoveClientConn        float64 `json:"node_num_remove_client_conn"`
	NodeNumRemoveClientSub         float64 `json:"node_num_remove_client_sub"`
	NodeNumRemovePresence          float64 `json:"node_num_remove_presence"`
	NodeNumUniqueClients           float64 `json:"node_num_unique_clients"`
	NodeUptimeSeconds              float64 `json:"node_uptime_seconds"`
}
