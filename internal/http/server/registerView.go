package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) RegisterView(w http.ResponseWriter, r *http.Request) {
	var data Register
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(RegisterResponse{Success: false, Error: "Decoding data failed"})
		return
	}
	err = s.agentStore.Set(data.AgentName, data.Credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(RegisterResponse{Success: false, Error: "Store data failed(save operations)"})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(RegisterResponse{Success: true, AgentName: data.AgentName})
}
