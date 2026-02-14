package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetAllAgents(w http.ResponseWriter, r *http.Request) {
	var data []string
	data = s.agentStore.AllKeys()
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
}
