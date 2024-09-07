package main

import (
	"go-boilerplate/configs"
	"go-boilerplate/internal/api/routes"
	"log"
	"os"

	"go-boilerplate/pkg/httpserver"
	"go-boilerplate/pkg/utils"
)

func main() {
	// Liveness Probe
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")
	// Load Configs
	configs.LoadConfigs(&configs.Configs)

	// Start server
	router := httpserver.NewRouter(true)
	routes.InitialServerRoutes(router, nil, nil)
	host := configs.Configs.Server.Host
	port := utils.StringToInt(configs.Configs.Server.Port)
	server := httpserver.NewRestAPI(host, port, router)
	if err := server.Start(); err != nil {
		panic("Failed to start http server: " + err.Error())
	}
}
