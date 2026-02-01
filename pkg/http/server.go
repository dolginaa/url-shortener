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
	http.Handle("/", s.mux)
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		log.Fatal("failed to start server: %w", err)
	}
}
