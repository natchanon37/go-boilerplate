package httpserver

import (
	"fmt"
	"log"
	"net/http"
)

const (
	SuccessStatus = "success"
	FailStatus    = "fail"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	// Meta    *paginate.Meta `json:"meta,omitempty"`
}

type FailResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	ErrCode string      `json:"err_code"`
	Errors  interface{} `json:"errors"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

type BindError struct {
	Message string
}

func (e *BindError) Error() string {
	return e.Message
}

func (f *FailResponse) Error() string {
	return f.Message
}

func AttachError(ctx Context, err error) {
	ctx.AttachError(err)
}

func NewFail(message string, code int) error {
	return &FailResponse{
		Code:    code,
		Status:  FailStatus,
		Message: message,
		Errors:  nil,
	}
}

func NotContent(ctx Context) {
	Success(ctx, &SuccessResponse{
		Code: http.StatusNoContent,
	})
}

func Created(ctx Context) {
	Success(ctx, &SuccessResponse{
		Code: http.StatusCreated,
	})
}

func Data(ctx Context, data interface{}) {
	Success(ctx, &SuccessResponse{
		Code: http.StatusOK,
		Data: data,
	})
}

func Success(ctx Context, r *SuccessResponse) {
	if ctx == nil {
		log.Println("Context is nil in Success function")
		return
	}
	if r == nil {
		log.Println("SuccessResponse is nil in Success function")
		ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		return
	}

	r.Status = SuccessStatus
	if r.Code == 0 {
		r.Code = http.StatusOK
	}
	if r.Message == "" {
		r.Message = ""
	}

	log.Printf("Attempting to call ctx.JSON with code %d and response %+v", r.Code, r)
	ctx.JSON(r.Code, r)
	log.Println("ctx.JSON call completed")
}

func UnprocessableEntity(ctx Context, errors interface{}) {
	Fail(ctx, &FailResponse{
		Code:    http.StatusUnprocessableEntity,
		Status:  FailStatus,
		Message: "",
		Errors:  errors,
	})
}

func BadRequest(ctx Context, message string) {
	Fail(ctx, NewFail(message, http.StatusBadRequest))
}

func Forbidden(ctx Context, message string) {
	Fail(ctx, NewFail(message, http.StatusForbidden))
}

func NotFound(ctx Context, message string) {
	Fail(ctx, NewFail(message, http.StatusNotFound))
}

func InternalServerError(ctx Context, message string) {
	Fail(ctx, NewFail(message, http.StatusInternalServerError))
}

func BadGateway(ctx Context, message string) {
	Fail(ctx, NewFail(message, http.StatusBadGateway))
}

func Fail(ctx Context, err error) {
	var response *FailResponse

	if e, ok := err.(*FailResponse); err != nil && ok {
		response = e
	} else {
		var message string
		if err == nil {
			message = "Server error occurred"
		} else {
			message = err.Error()
		}

		response = &FailResponse{
			Code:    500,
			Status:  FailStatus,
			Message: message,
		}
	}

	ctx.JSON(response.Code, response)
}

func FileStreamBinaries(ctx Context, fileBytes []byte, filename string) {
	contentType := "application/octet-stream"

	ctx.SetHeader("Content-Description", "File Transfer")
	ctx.SetHeader("Content-Transfer-Encoding", "binary")
	ctx.SetHeader("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.SetHeader("Content-Type", contentType)

	ctx.Data(http.StatusOK, contentType, fileBytes)
}
