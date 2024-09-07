package health_controller

import "go-boilerplate/pkg/httpserver"

type healthCtrl struct{}

type HealthCtrl interface {
	HealthCrl(ctx httpserver.Context)
}

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
	httpserver.Data(ctx, "OK")
}

func NewHealthCtrl() HealthCtrl {
	return &healthCtrl{}
}
