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
}

func NewServer(host string, port int, agentStore store.StoreInterface) *Server {
	return &Server{
		Host:       host,
		Port:       port,
		agentStore: agentStore,
	}
}

func (s *Server) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/agents", func(r chi.Router) {
		r.Post("/register", s.RegisterView)
		r.Get("/", s.GetAllAgents)
	})
	r.Get("/health", s.HealthCheck)
	http.ListenAndServe(fmt.Sprintf("%v:%v", s.Host, s.Port), r)
}
