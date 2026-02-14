package server

type MainStruct struct {
	UserText string `json:"userText"`
}

type MainResponse struct {
	TaskID  string `json:"taskID,omitempty"`
	Status  string `json:"status,omitempty"`
	Success bool   `json:"success"`
}
