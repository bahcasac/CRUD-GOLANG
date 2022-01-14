package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go-banco-de-dados/service"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", service.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", service.SearchUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", service.SearchUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", service.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", service.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Listen on port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
