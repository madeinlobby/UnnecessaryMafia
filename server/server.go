package server

import (
	"encoding/json"
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
	server.router.HandleFunc("/helloworld", helloWorldHandler).Methods("GET")
	server.router.HandleFunc("/register", RegisterHandler).Methods("POST")
	http.Handle("/", server.router)
	if err := http.ListenAndServe(server.port, server.router); err != nil {
		log.Fatalf("failed to listen")
	}
}

func helloWorldHandler(writer http.ResponseWriter, request *http.Request) {
	type singleResponse struct {
		Message string `json:"message"`
	}
	resp := singleResponse{Message: "Hello World"}
	jsonResp, err := json.MarshalIndent(resp, "", "")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = writer.Write(jsonResp)
	if err != nil {
		log.Printf("could not write response: %s", request.RequestURI)
	}
	return
}
