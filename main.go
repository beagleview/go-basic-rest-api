package main

import (
	"basic-rest/handles"
	"log"
	"net/http"
)

func handleRequest() {
	log.Println("Start Service")
	router := handles.Router()

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequest()
}
