package services

import (
	"context"
	"fmt"
	"url_shorterner_m/models"
	"url_shorterner_m/storage"
	"url_shorterner_m/utils"
)

// CREATE
func CreateShortURL(longURL string) string {

	// Step 1: Generate unique ID using Redis
	id, err := storage.RedisClient.Incr(storage.Ctx, "url_counter").Result()
	if err != nil {
		return ""
	}

	// Step 2: Convert ID → Base62
	shortCode := utils.EncodeBase62(id)

	url := models.URL{
		ShortCode: shortCode,
		LongURL:   longURL,
	}

	// Step 3: Store in MongoDB
	_, err = storage.URLCollection.InsertOne(context.Background(), &url)
	if err != nil {
		fmt.Printf("❌ MongoDB Insert Error: %v\n", err)
	}

	// Step 4: Cache in Redis
	storage.RedisClient.Set(storage.Ctx, shortCode, longURL, 0)

	return shortCode
}

// READ
func GetLongURL(shortCode string) (string, bool) {

	// 1. Check Redis
	longURL, err := storage.RedisClient.Get(storage.Ctx, shortCode).Result()
	if err == nil {
		return longURL, true
	}

	// 2. Check MongoDB
	var result models.URL
	err = storage.URLCollection.FindOne(
		context.Background(),
		map[string]string{"short_code": shortCode},
	).Decode(&result)

	if err != nil {
		return "", false
	}

	// 3. Cache in Redis
	storage.RedisClient.Set(storage.Ctx, shortCode, result.LongURL, 0)

	return result.LongURL, true
}
