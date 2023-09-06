package main

import (
	"log"
	"net/http"

	"github.com/Alperen10/Kandilli-API/client"
	"github.com/Alperen10/Kandilli-API/handlers"
)

func main() {
	c := client.NewHTTPClient()

	http.HandleFunc("/", handlers.EarthquakeHandler(c))

	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("unable to listen on port 8080")
	}
}
