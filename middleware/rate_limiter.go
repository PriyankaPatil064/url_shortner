package middleware

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

// Token Bucket structure
type TokenBucket struct {
	tokens         float64
	capacity       float64
	refillRate     float64
	lastRefillTime time.Time
	mutex          sync.Mutex
}

// Create new bucket
func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		tokens:         capacity,
		capacity:       capacity,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

// Allow request or not
func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefillTime).Seconds()

	// refill tokens
	tb.tokens += elapsed * tb.refillRate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	tb.lastRefillTime = now

	if tb.tokens >= 1 {
		tb.tokens--
		return true
	}

	return false
}

// Store buckets per IP
var buckets = make(map[string]*TokenBucket)
var mu sync.Mutex

// Get client IP
func getIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

// Middleware
func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := getIP(r)

		mu.Lock()
		bucket, exists := buckets[ip]
		if !exists {
			bucket = NewTokenBucket(4, 0)
			buckets[ip] = bucket
		}
		mu.Unlock()

		// ✅ call outside lock
		if !bucket.Allow() {
			fmt.Println("❌ Rate limit exceeded for:", ip)
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		fmt.Println("✅ Request allowed for:", ip)
		next.ServeHTTP(w, r)
	})
}

