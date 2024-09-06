package routes

import (
	health_handler "go-boilerplate/handlers/health"
	"go-boilerplate/pkg/http"
)

func InitialWorker(route http.Router) {
	healthCtrl := health_handler.NewHealthCtrl()

	route.GET("/health", healthCtrl.Health)
}

func InitialServer(route http.Router) {
	healthCtrl := health_handler.NewHealthCtrl()

	route.GET("/health", healthCtrl.Health)
}
