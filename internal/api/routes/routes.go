package routes

import (
	health_controller "go-boilerplate/internal/api/controllers/health"
	"go-boilerplate/pkg/httpserver"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func InitialWorkerRoutes(route httpserver.Router, db *gorm.DB, redis *redis.Client) {
	healthCtrl := health_controller.NewHealthCtrl()

	route.GET("/health", healthCtrl.HealthCrl)
}

func InitialServerRoutes(route httpserver.Router, db *gorm.DB, redis *redis.Client) {
	healthCtrl := health_controller.NewHealthCtrl()
	system := route.Group("/system")
	system.GET("/health", healthCtrl.HealthCrl)
}
