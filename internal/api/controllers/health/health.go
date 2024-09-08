package health_controller

import "go-boilerplate/pkg/httpserver"

type healthCtrl struct{}

type HealthCtrl interface {
	HealthCrl(ctx httpserver.Context)
}

// @Summary      system health check
// @Description  system health check
// @Tags         system
// @Router       /system/health [get]
// @Accept       json
// @Produce      json
// @Success      200  {object}  controller_response.HealthResponse
func (ctrl *healthCtrl) HealthCrl(ctx httpserver.Context) {
	if ctrl == nil {
		// Log this unexpected state
		ctx.JSON(500, map[string]string{"error": "Health controller is nil"})
		return
	}
	if ctx == nil {
		// This shouldn't happen, but let's be safe
		panic("Context is nil in HealthCrl")
	}
	httpserver.Success(ctx, &httpserver.SuccessResponse{
		Data: map[string]string{"status": "ok"},
	})
}

func NewHealthCtrl() HealthCtrl {
	return &healthCtrl{}
}
