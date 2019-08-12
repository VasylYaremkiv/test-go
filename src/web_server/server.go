package web_server

import (
	"fmt"
	"net/http"
	"strconv"

	// "os"

	"github.com/gorilla/mux"
)

type Server struct {
	port string
}

func CreateServer(portNumber int) (*Server, error) {
	return &Server{
		port: strconv.Itoa(portNumber),
	}, nil
}

func (s *Server) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/", rootHandler).Methods("GET")

	http.Handle("/", router)

	fmt.Println("Server started, listening on port " + s.port + "...")
	fmt.Println("CTL-C to break out of broker")
	http.ListenAndServe(":"+s.port, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Root!\n"))
}
