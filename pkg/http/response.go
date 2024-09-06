package http

import "net/http"

const (
	SuccessStatus = "success"
	FailStatus    = "fail"
)

type BindError struct {
	Message string
}

func (e *BindError) Error() string {
	return e.Message
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	// Meta    *paginate.Meta `json:"meta,omitempty"`
}

func Success(ctx Context, r *SuccessResponse) {
	r.Status = SuccessStatus
	if r.Code == 0 {
		r.Code = http.StatusOK
	}
	if r.Message == "" {
		r.Message = ""
	}

	ctx.JSON(r.Code, r)
}
