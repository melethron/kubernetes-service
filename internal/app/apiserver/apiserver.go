package apiserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	config *Config
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/", s.helloWorld())
}

func (s *APIServer) helloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	}
}
