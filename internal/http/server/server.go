package server

import (
	"fmt"
	"github.com/Chutchev/coordinatorAgent/internal/http/handlers/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Server struct {
	Host string
	Port int
}

func NewServer(host string, port int) *Server {
	return &Server{
		Host: host,
		Port: port,
	}
}

func (s Server) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/agents", func(r chi.Router) {
		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	})
	r.Get("/health", health.HealthCheck)
	http.ListenAndServe(fmt.Sprintf("%v:%v", s.Host, s.Port), r)
}
