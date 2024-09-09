package routes

import (
	customer_controller "go-boilerplate/internal/api/controllers/customer"
	health_controller "go-boilerplate/internal/api/controllers/health"
	repository_customer "go-boilerplate/internal/repository/customer"
	service_customer "go-boilerplate/internal/services"
	"go-boilerplate/pkg/httpserver"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func InitialWorkerRoutes(route httpserver.Router, db *gorm.DB, redis *redis.Client) {
	healthCtrl := health_controller.NewHealthCtrl()

	route.GET("/health", healthCtrl.HealthCrl)
}

func InitialServerRoutes(route httpserver.Router, db *gorm.DB, redis *redis.Client) {
	// initial repository
	customerRepo := repository_customer.NewCustomerRepository(db)

	// initial service
	customerSvc := service_customer.NewCustomerService(customerRepo)

	// initial controller
	healthCtrl := health_controller.NewHealthCtrl()
	customerCtrl := customer_controller.NewCustomerCtrl(customerSvc)
	system := route.Group("/system")
	system.GET("/health", healthCtrl.HealthCrl)

	v1 := route.Group("/v1")
	customer := v1.Group("/customer")
	customer.POST("/create-customer", httpserver.DbTransactionMiddleware(customerCtrl.CreateCustomer, db))
	customer.GET("/:cus_id", customerCtrl.GetCustomerData)

}
