package main

import (
	"log"
	"net/http"

	"github.com/beerus/manager"
)

func main() {
	log.Printf("Server started")

	router := manager.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
