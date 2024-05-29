package main

import (
	"log"
	"net/http"

	"LoadingData/internal/db"
	"LoadingData/internal/handlers"
)

func main() {
	store, err := db.NewElasticsearchStore()
	if err != nil {
		log.Fatalf("Failed to create store: %v", err)
	}


	http.HandleFunc("/api/get_token", handlers.GetTokenHandler)
	http.HandleFunc("/api/protected", handlers.ProtectedHandler)
	
	http.HandleFunc("/api/places", handlers.GetPlacesHandler(store))
	http.HandleFunc("/api/recommend", handlers.GetNearestPlacesHandler(store))

	log.Println("Starting server on :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
