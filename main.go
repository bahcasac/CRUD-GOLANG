package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go-banco-de-dados/services"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", services.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", services.SearchUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", services.SearchUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", services.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", services.DeleteUser).Methods(http.MethodDelete)
	fmt.Println("Listen on port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))

}
