package c_err

import "net/http"

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type CErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

func (c *CErr) Error() string {
	return c.Err
}

func NewCErr(message, err string, code int, causes []Causes) *CErr {
	return &CErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestErr(message string) *CErr {
	return &CErr{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationErr(message string, causes []Causes) *CErr {
	return &CErr{
		Message: message,
		Err:     "Bad Request Validation",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerErr(message string) *CErr {
	return &CErr{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundErr(message string) *CErr {
	return &CErr{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenErr(message string) *CErr {
	return &CErr{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}
