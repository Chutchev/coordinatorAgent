package server

type Credentials struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Register struct {
	AgentName   string      `json:"agent_name"`
	Credentials Credentials `json:"credentials"`
}

type RegisterResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error,omitempty"`
	AgentName string `json:"agent_name,omitempty"`
}
