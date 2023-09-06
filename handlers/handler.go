package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alperen10/Kandilli-API/client"
)

func EarthquakeHandler(c client.Client) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c.Fetch()
		res, _ := c.ParseToLines()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Fatal("unable to encode to json")
		}
	}
}
