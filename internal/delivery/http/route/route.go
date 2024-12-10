// internal/delivery/http/route/route.go
package route

import (
	"github.com/gin-gonic/gin"
	"github.com/lunadiotic/golang-travel-booking/internal/delivery/http/handler"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/usecase"
	"github.com/lunadiotic/golang-travel-booking/pkg/middleware"
)

type RouterConfig struct {
	userUseCase        usecase.UserUseCase
	destinationUseCase usecase.DestinationUseCase
}

func NewRouter(userUseCase usecase.UserUseCase, destinationUseCase usecase.DestinationUseCase) *RouterConfig {
	return &RouterConfig{
		userUseCase:        userUseCase,
		destinationUseCase: destinationUseCase,
	}
}

func (rc *RouterConfig) SetupRoutes(r *gin.Engine, jwtSecret string) {
	r.Use(middleware.LoggingMiddleware())

	// Initialize handlers
	authHandler := handler.NewAuthHandler(rc.userUseCase)
	userHandler := handler.NewUserHandler(rc.userUseCase)
	destinationHandler := handler.NewDestinationHandler(rc.destinationUseCase)

	// Public routes
	public := r.Group("/api/v1")
	{
		public.GET("/health", healthCheck)

		auth := public.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		destinations := public.Group("/destinations")
		{
			destinations.GET("", destinationHandler.GetAll)
			destinations.GET("/:id", destinationHandler.GetByID)
		}
	}

	// Middleware
	authMiddleware := middleware.AuthMiddleware(jwtSecret)

	// Protected routes
	protected := r.Group("/api/v1")
	protected.Use(authMiddleware)
	{
		user := protected.Group("/users")
		{
			user.GET("/profile", userHandler.GetProfile)
			user.PUT("/profile", userHandler.UpdateProfile)
		}

		destinations := protected.Group("/destinations")
		{
			destinations.POST("", destinationHandler.Create)
			destinations.PUT("/:id", destinationHandler.Update)
			destinations.DELETE("/:id", destinationHandler.Delete)
		}
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "Service is healthy",
	})
}
