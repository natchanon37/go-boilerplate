package controller_response

import (
	service_customer "go-boilerplate/internal/services"
	"go-boilerplate/pkg/httpserver"
)

type CreateCustomerResponse struct {
	httpserver.SuccessResponse
}

type GetCustomerResponse struct {
	httpserver.SuccessResponse
	Data service_customer.GetCustomerData `json:"data"`
}
