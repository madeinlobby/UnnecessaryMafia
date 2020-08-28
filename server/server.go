package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
	port   string
}

func (server *Server) initServer() {
	server.router = mux.NewRouter()
	server.port = ":8080"
}

func (server *Server) Run() {
	server.initServer()
	//todo

	if err := http.ListenAndServe(server.port, server.router); err != nil {
		log.Fatalf("failed to listen")
	}
}
