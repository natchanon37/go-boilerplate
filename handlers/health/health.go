package health_handler

import "go-boilerplate/pkg/http"

type HealthCtrl struct{}

func NewHealthCtrl() *HealthCtrl {
	return &HealthCtrl{}
}

func (ctrl *HealthCtrl) Health(ctx http.Context) {
	http.Success(ctx, &http.SuccessResponse{})
}
