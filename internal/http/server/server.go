package server

import (
	"fmt"
	"github.com/Chutchev/coordinatorAgent/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Server struct {
	Host       string
	Port       int
	agentStore store.StoreInterface
	taskStore  store.StoreInterface
}

func NewServer(host string, port int, agentStore, taskStor store.StoreInterface) *Server {
	return &Server{
		Host:       host,
		Port:       port,
		agentStore: agentStore,
		taskStore:  taskStor,
	}
}

func (s *Server) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/agents", func(r chi.Router) {
		r.Post("/register", s.RegisterView)
		r.Get("/", s.GetAllAgents)
	})
	r.Post("/do", s.Do)
	r.Get("/health", s.HealthCheck)
	http.ListenAndServe(fmt.Sprintf("%v:%v", s.Host, s.Port), r)
}
