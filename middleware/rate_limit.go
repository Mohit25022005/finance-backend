package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	visitors = make(map[string]*visitor)
	mu       sync.RWMutex
)

// cleanupVisitors removes old entries to prevent memory leak
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(1, 5) // 1 req/sec, burst 5
		visitors[ip] = &visitor{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

// getClientIP extracts real client IP safely
func getClientIP(c *gin.Context) string {
	ip := c.ClientIP()
	if parsed := net.ParseIP(ip); parsed != nil {
		return ip
	}
	return "unknown"
}

func RateLimitMiddleware() gin.HandlerFunc {
	// start cleanup goroutine once
	go cleanupVisitors()

	return func(c *gin.Context) {

		// allow swagger without limits
		if len(c.Request.URL.Path) >= 8 && c.Request.URL.Path[:8] == "/swagger" {
			c.Next()
			return
		}

		ip := getClientIP(c)
		limiter := getLimiter(ip)

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}