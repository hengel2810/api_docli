package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/image", api.HandlePostImage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

