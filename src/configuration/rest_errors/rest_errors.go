package rest_errors

import "net/http"

type RestErrors struct {
	Message string `json:"message"`
	Err string `json:"error"`
	Code int `json:"code"`
	Causes []Causes `json:"causes"`
}

type Causes struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

// Some libs needs to param Error
// So, I created this func for use in this libs how param
func (r *RestErrors) Error() string {
	return r.Message
}

func NewRestErrors(message, err string, code int, causes []Causes) *RestErrors {
	return &RestErrors{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErrors {
	return &RestErrors{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
		Causes: causes,
	}
}

func NewINternalServerError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Err: "internal_server_error",
		Code: http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Err: "not_found",
		Code: http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewUnauthorizedRequestError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}