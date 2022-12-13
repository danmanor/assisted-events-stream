package types

type Event struct {
	Name     string                 `json:"name"`
	Payload  interface{}            `json:"payload"`
	Metadata map[string]interface{} `json:"metadata"`
}

type State struct {
	Metadata map[string]interface{} `json:"metadata"`
	Payload  map[string]interface{} `json:"payload"`
}

type EnrichedEvent struct {
	// we'll be using a UUID based on event fields ID until we'll receive a proper UUID from the service
	ID        string                   `json:"id"`
	Message   string                   `json:"message"`
	Category  string                   `json:"category"`
	ClusterID string                   `json:"cluster_id"`
	EventTime string                   `json:"event_time"`
	Name      string                   `json:"name"`
	RequestID string                   `json:"request_id"`
	Severity  string                   `json:"severity"`
	Cluster   map[string]interface{}   `json:"cluster"`
	InfraEnvs []map[string]interface{} `json:"infra_envs"`
	Versions  map[string]string        `json:"versions"`
}

type ClusterState State

type HostState State

type InfraEnvState State
