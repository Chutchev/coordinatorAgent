package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (s *Server) Do(w http.ResponseWriter, r *http.Request) {
	var req MainStruct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(MainResponse{Success: false, Status: err.Error()})
	}
	err = s.taskStore.Set("123", req.UserText)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MainResponse{Success: true, Status: "Task Created", TaskID: "123"})
}
