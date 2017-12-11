package main

import (
	"log"
	"net/http"
	"os"

	"github.com/beerus/manager"
)

func main() {
	log.Printf("Server started")

	router := manager.NewRouter()

	manager.InitDB(os.Getenv("CONN_STRING"))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
