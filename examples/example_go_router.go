package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutes() {
	router := mux.NewRouter()

	// Este endpoint será detectado automaticamente pelo GoRouterExtractor
	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// implementação
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// implementação
	fmt.Fprintf(w, "teste, teste, teste")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// implementação
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// implementação
}
