package routes

import (
	health_controller "go-boilerplate/internal/api/controllers/health"
	"go-boilerplate/pkg/http"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func InitialWorkerRoutes(route http.Router, db *gorm.DB, redis *redis.Client) {
	healthCtrl := health_controller.NewHealthCtrl()

	route.GET("/health", healthCtrl.Health)
}

func InitialServerRoutes(route http.Router, db *gorm.DB, redis *redis.Client) {
	healthCtrl := health_controller.NewHealthCtrl()

	route.GET("/health", healthCtrl.Health)
}
