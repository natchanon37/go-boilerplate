package main

import (
	"go-boilerplate/configs"
	"go-boilerplate/internal/api/routes"
	"go-boilerplate/pkg/http"
	"go-boilerplate/pkg/utils"
)

func main() {
	// Load Configs
	configs.LoadConfigs(&configs.Configs)

	// Start server
	router := http.NewRouter(true, nil)
	routes.InitialWorkerRoutes(router, nil, nil)
	host := configs.Configs.Server.Host
	port := utils.StringToInt(configs.Configs.Server.Port)
	server := http.NewRestAPI(host, port, router)
	if err := server.Start(); err != nil {
		panic("Failed to start http server: " + err.Error())
	}
}
