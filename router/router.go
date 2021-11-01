package router

import (
	"RestApi/users"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/users", users.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", users.GetUser).Methods("GET")
	router.HandleFunc("/users", users.CreateUsers).Methods("POST")
	router.HandleFunc("/users/{id}", users.UpdateUsers).Methods("PUT")
	router.HandleFunc("/users/{id}", users.DeleteUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
}
