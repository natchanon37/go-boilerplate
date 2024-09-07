package routes

import (
	health_controller "go-boilerplate/internal/api/controller/health"
	"go-boilerplate/pkg/http"
)

func InitialWorker(route http.Router) {
	healthCtrl := health_controller.NewHealthCtrl()

	route.GET("/health", healthCtrl.Health)
}

func InitialServer(route http.Router) {
	healthCtrl := health_controller.NewHealthCtrl()

	route.GET("/health", healthCtrl.Health)
}
