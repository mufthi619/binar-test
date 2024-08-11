package response

import (
	"net/http"
)

type ResponseFormatter struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponseFormatter() *ResponseFormatter {
	return &ResponseFormatter{}
}

func (data *ResponseFormatter) SetCode(code int) {
	data.Code = code
}

func (data *ResponseFormatter) SetMessage(msg string) {
	data.Message = msg
}

func (data *ResponseFormatter) SetData(resp any) {
	data.Data = resp
}

func (data *ResponseFormatter) ReturnSuccessfullyWithData(resp any, msg ...string) {
	data.Code = http.StatusOK
	if len(msg) > 0 {
		data.Message = msg[0]
	} else {
		data.Message = "Successfully"
	}
	data.Data = resp
}

func (data *ResponseFormatter) ReturnSuccessfullyWithoutData(msg ...string) {
	data.Code = http.StatusOK
	if len(msg) > 0 {
		data.Message = msg[0]
	} else {
		data.Message = "Successfully"
	}
	data.Data = nil
}

func (data *ResponseFormatter) ReturnCreatedWithoutData(msg ...string) {
	data.Code = http.StatusCreated
	if len(msg) > 0 {
		data.Message = msg[0]
	} else {
		data.Message = "Successfully"
	}
	data.Data = nil
}

func (data *ResponseFormatter) ReturnCreatedWithData(resp any, msg ...string) {
	data.Code = http.StatusCreated
	if len(msg) > 0 {
		data.Message = msg[0]
	} else {
		data.Message = "Successfully"
	}
	data.Data = resp
}

func (data *ResponseFormatter) ReturnInternalServerError(msg ...string) {
	data.Code = http.StatusInternalServerError
	if len(msg) > 0 {
		data.Message = msg[0]
	} else {
		data.Message = "Failed ! There's some trouble in our system, please try again"
	}
	data.Data = nil
}

func (data *ResponseFormatter) ReturnInternalUnavailable(msg ...string) {
	data.Code = http.StatusServiceUnavailable
	if len(msg) > 0 {
		data.Message = msg[0]
	} else {
		data.Message = "Failed ! There's some trouble in our system, please try again"
	}
	data.Data = nil
}

func (data *ResponseFormatter) ReturnUnauthorized(msg ...string) {
	data.Code = http.StatusUnauthorized
	if len(msg) > 0 {
		data.Message = msg[0]
	} else {
		data.Message = "Failed ! Unauthorized"
	}
	data.Data = nil
}

func (data *ResponseFormatter) ReturnBadRequest(msg ...string) {
	data.Code = http.StatusBadRequest
	if len(msg) > 0 {
		data.Message = msg[0]
	} else {
		data.Message = "Failed ! Invalid body"
	}
	data.Data = nil
}
