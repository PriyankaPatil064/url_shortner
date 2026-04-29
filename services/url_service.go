package services

import (
	"context"
	"url_shorterner_m/models"
	"url_shorterner_m/storage"
	"url_shorterner_m/utils"
)

// CREATE
func CreateShortURL(longURL string) string {
	shortCode := utils.GenerateShortCode(8)

	url := models.URL{
		ShortCode: shortCode,
		LongURL:   longURL,
	}

	// Save in MongoDB
	storage.URLCollection.InsertOne(context.Background(), url)

	// Save in Redis (cache)
	storage.RedisClient.Set(storage.Ctx, shortCode, longURL, 0)

	return shortCode
}

// READ
func GetLongURL(shortCode string) (string, bool) {

	// 🔥 1. Check Redis first
	longURL, err := storage.RedisClient.Get(storage.Ctx, shortCode).Result()
	if err == nil {
		return longURL, true
	}

	// 🔥 2. If not in Redis → check MongoDB
	var result models.URL
	err = storage.URLCollection.FindOne(
		context.Background(),
		map[string]string{"short_code": shortCode},
	).Decode(&result)

	if err != nil {
		return "", false
	}

	// 🔥 3. Store in Redis for next time
	storage.RedisClient.Set(storage.Ctx, shortCode, result.LongURL, 0)

	return result.LongURL, true
}