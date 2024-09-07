package main

import (
	"go-boilerplate/configs"
	"go-boilerplate/internal/controller/routes"
	"go-boilerplate/pkg/http"
	"go-boilerplate/pkg/utils"
)

func main() {
	// Load Configs
	configs.LoadConfigs(&configs.Configs)
	host := configs.Configs.Server.Host
	port := utils.StringToInt(configs.Configs.Server.Port)

	// Start server
	router := http.NewRouter(true, nil)
	routes.InitialWorker(router)

	server := http.NewRestAPI(host, port, router)
	go func() {
		if err := server.Start(); err != nil {
			panic("Failed to start http server: " + err.Error())
		}
	}()

}
