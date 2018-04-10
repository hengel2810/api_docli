package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/hengel2810/api_docli/api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/image", api.HandlePostImage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

