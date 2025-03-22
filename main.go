package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"playsee.co/interview/api/middleware"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.ApiKeyAuthMiddleware)
	router.HandleFunc("/test-1", Test1).Methods("POST")

	port := 8082
	fmt.Println("Playsee Backend Interview Test server is running at port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
