// pkg/middleware/auth_test.go
package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Valid Token", func(t *testing.T) {
		// Create a valid token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": "test-user-id",
			"email":   "test@example.com",
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString([]byte(jwtSecret))
		assert.NoError(t, err)

		// Setup request
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		r.Use(AuthMiddleware())
		r.GET("/test", func(c *gin.Context) {
			userId, exists := c.Get("user_id")
			assert.True(t, exists)
			assert.Equal(t, "test-user-id", userId)
			c.Status(http.StatusOK)
		})

		// Make request
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		c.Request = req
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Missing Authorization Header", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		r.Use(AuthMiddleware())
		r.GET("/test", func(c *gin.Context) {
			t.Error("Handler should not be called")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		c.Request = req
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Invalid Token Format", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		r.Use(AuthMiddleware())
		r.GET("/test", func(c *gin.Context) {
			t.Error("Handler should not be called")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "InvalidFormat token")
		c.Request = req
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		r.Use(AuthMiddleware())
		r.GET("/test", func(c *gin.Context) {
			t.Error("Handler should not be called")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer invalid.token.string")
		c.Request = req
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Expired Token", func(t *testing.T) {
		// Create an expired token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": "test-user-id",
			"email":   "test@example.com",
			"exp":     time.Now().Add(-time.Hour).Unix(), // expired 1 hour ago
		})
		tokenString, err := token.SignedString([]byte(jwtSecret))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		r.Use(AuthMiddleware())
		r.GET("/test", func(c *gin.Context) {
			t.Error("Handler should not be called")
		})

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		c.Request = req
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
