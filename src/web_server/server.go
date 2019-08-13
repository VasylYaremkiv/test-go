package web_server

import (
	"fmt"
	"net/http"
	"strconv"
	"model"

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
	router.HandleFunc("/create", createHandler).Methods("GET")

	http.Handle("/", router)

	fmt.Println("Server started, listening on port " + s.port + "...")
	fmt.Println("CTL-C to break out of broker")
	http.ListenAndServe(":"+s.port, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All!\n"))
	users, err := model.FindAll()
	if err != nil {
		fmt.Println("Find all error: %s", err)
	}
	fmt.Println("Users: %s", users)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create!\n"))
	user, err := model.Create("Test")
	if err != nil {
		fmt.Println("Create user error: %s", err)
	}
	fmt.Println("User: %s", user)
}
