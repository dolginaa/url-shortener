package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	mux  *mux.Router
}

func NewServer(mux *mux.Router) Server {
	return Server{
		addr: ":7000",
		mux:  mux,
	}
}

func (s *Server) Start() {
	log.Print("Starting...")
	if err := http.ListenAndServe(s.addr, s.mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
