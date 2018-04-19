package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/hengel2810/api_docli/middleware"
	"github.com/codegangsta/negroni"
	"github.com/hengel2810/api_docli/api"
	"fmt"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", testAPI).Methods("GET")
	jwtMiddleware := middleware.JWTMiddleware()
	fmt.Println("############## 111 ################")
	router.Handle("/image", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("############## 222 ################")
			api.HandlePostImage(w, r)
	}))))
	log.Fatal(http.ListenAndServe(":8000", router))
}

func testAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

