package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	_"github.com/go-sql-driver/mysql"
	"net/http"
)

func main(){

	router := mux.NewRouter()

	fmt.Println("Listen on port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
