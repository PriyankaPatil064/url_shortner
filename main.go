package main

import (
	"fmt"
	"net/http"
	"url_shorterner_m/handlers"
	"url_shorterner_m/middleware"
	"url_shorterner_m/storage"
)

func main() {
	// Initialize databases
	storage.InitMongo()
	storage.InitRedis()

	// Routes with rate limiter
	http.Handle("/shorten",
		middleware.RateLimiter(http.HandlerFunc(handlers.ShortenHandler)),
	)

	http.Handle("/",
		middleware.RateLimiter(http.HandlerFunc(handlers.RedirectHandler)),
	)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}