package main

import (
	"fmt"
	"net/http"
	"url_shorterner_m/handlers"
	"url_shorterner_m/storage"
)

func main() {
	// initialize MongoDB
	storage.InitMongo()
	storage.InitRedis()

	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
