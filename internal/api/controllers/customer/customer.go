package customer_controller

import (
	service_customer "go-boilerplate/internal/services"
	"go-boilerplate/pkg/httpserver"
	"go-boilerplate/pkg/utils"

	"gorm.io/gorm"
)

type CustomerCtrl interface {
	CreateCustomer(ctx httpserver.Context)
	GetCustomerData(ctx httpserver.Context)
}
type customerCtrl struct {
	cusSvc service_customer.CustomerService
}

// @Summary Create Customer
// @Description Create Customer
// @Tags customer
// @Router /v1/customer/create-customer [post]
// @Accept json
// @Produce json
// @Param request body service_customer.CreateCustomerReq true "Create Customer Request"
// @Success 200 {object} controller_response.CreateCustomerResponse
func (ctrl *customerCtrl) CreateCustomer(ctx httpserver.Context) {
	var r service_customer.CreateCustomerReq

	if err := ctx.Bind(&r); err != nil {
		httpserver.AttachError(ctx, err)
		return
	}

	if err := utils.ValidateStruct(r); err != nil {
		httpserver.AttachError(ctx, err)
		return
	}

	reqCtx := ctx.GetRequestCtx()
	txHandle, _ := ctx.Get("db_tx")
	err := ctrl.cusSvc.WithTx(txHandle.(*gorm.DB)).CreateCustomer(reqCtx, r)
	if err != nil {
		httpserver.AttachError(ctx, err)
		return
	}
	httpserver.Success(ctx, &httpserver.SuccessResponse{
		Data: map[string]string{"create": "success"},
	})
}

// @Summary Get Customer Data
// @Description Get Customer Data
// @Tags customer
// @Router /v1/customer/{cus_id} [get]
// @Accept json
// @Produce json
// @Param        cus_id  path  string  true  "customer id"
// @Success 200 {object} controller_response.GetCustomerResponse
func (ctrl *customerCtrl) GetCustomerData(ctx httpserver.Context) {
	cusId := ctx.GetParam("cus_id")

	reqCtx := ctx.GetRequestCtx()
	cusData, err := ctrl.cusSvc.GetCustomerById(reqCtx, cusId)
	if err != nil {
		httpserver.AttachError(ctx, err)
		return
	}

	httpserver.Success(ctx, &httpserver.SuccessResponse{
		Data: cusData,
	})
}

func NewCustomerCtrl(
	cusSvc service_customer.CustomerService,
) CustomerCtrl {
	return &customerCtrl{
		cusSvc: cusSvc,
	}
}
