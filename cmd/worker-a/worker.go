package main

import (
	"go-boilerplate/configs"

	"go-boilerplate/internal/api/routes"
	"go-boilerplate/internal/kafka"
	"go-boilerplate/pkg/httpserver"
	"go-boilerplate/pkg/utils"
)

func main() {
	// Load Configs
	configs.LoadConfigs(&configs.Configs)

	// init gin router
	router := httpserver.NewRouter(true)

	// set up routes
	routes.InitialWorkerRoutes(router, nil, nil)

	// Start rest server
	host := configs.Configs.Server.Host
	port := utils.StringToInt(configs.Configs.Server.Port)
	server := httpserver.NewRestAPI(host, port, router)
	go func() {
		if err := server.Start(); err != nil {
			panic("Failed to start http server: " + err.Error())
		}
	}()

	// set up worker
	worker := kafka.NewWorkerA(configs.Configs.WorkerAConsumer)
	worker.Start()
}
