package server

import (
	"balance/internal/balance/store"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	config *Config
	store  store.AccountStore
}

func New(config *Config, store store.AccountStore) *Server {
	server := Server{
		router: mux.NewRouter(),
		config: config,
		store:  store,
	}

	server.configHandlers()

	log.Println("Create server")

	return &server
}

func (s *Server) configHandlers() {
	s.router.HandleFunc("/transaction", s.handleTransaction()).Methods("POST")

	s.router.HandleFunc("/balance", s.handleGetBalance()).Methods("GET")
	s.router.HandleFunc("/balance", s.handleBalanceUpdate()).Methods("POST")
}

func (s *Server) Start() error {
	log.Println("Start server")
	return http.ListenAndServe(s.config.Address, s.router)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
