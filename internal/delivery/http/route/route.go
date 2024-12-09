// internal/delivery/http/route/route.go
package route

import (
	"github.com/lunadiotic/golang-travel-booking/internal/domain/usecase"

	"github.com/lunadiotic/golang-travel-booking/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	userUseCase usecase.UserUseCase
}

func NewRouter(userUseCase usecase.UserUseCase) *RouterConfig {
	return &RouterConfig{
		userUseCase: userUseCase,
	}
}

func (rc *RouterConfig) SetupRoutes(r *gin.Engine) {
	// Initialize handlers
	authHandler := handler.NewAuthHandler(rc.userUseCase)
	userHandler := handler.NewUserHandler(rc.userUseCase)

	// Public routes
	public := r.Group("/api/v1")
	{
		public.GET("/health", healthCheck)

		auth := public.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}

	// Protected routes
	protected := r.Group("/api/v1")
	// protected.Use(middleware.AuthMiddleware())
	{
		user := protected.Group("/users")
		{
			user.GET("/profile", userHandler.GetProfile)
			user.PUT("/profile", userHandler.UpdateProfile)
		}
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "Service is healthy",
	})
}
