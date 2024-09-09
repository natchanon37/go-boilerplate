package main

import (
	"go-boilerplate/configs"
	"go-boilerplate/internal/api/routes"
	"go-boilerplate/internal/database"
	database_mysql "go-boilerplate/internal/database/mysql"
	"log"
	"os"

	_ "go-boilerplate/docs"
	"go-boilerplate/pkg/httpserver"
	"go-boilerplate/pkg/utils"

	_ "ariga.io/atlas-provider-gorm/gormschema"
)

// @title           Go Boilerplate API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9005

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Liveness Probe
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")
	// Load Configs
	configs.LoadConfigs(&configs.Configs)

	// mysql connection
	var db database.Database             // common interface for all databases
	mysqlDb := &database_mysql.MySqlDb{} // mysql database
	if err := mysqlDb.Connect(configs.Configs.Database); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	db = mysqlDb
	defer db.Close()

	// Start server
	router := httpserver.NewRouter(true)
	routes.InitialServerRoutes(router, db.GetDB(), nil)
	host := configs.Configs.Server.Host
	port := utils.StringToInt(configs.Configs.Server.Port)
	server := httpserver.NewRestAPI(host, port, router)
	if err := server.Start(); err != nil {
		panic("Failed to start http server: " + err.Error())
	}
}
