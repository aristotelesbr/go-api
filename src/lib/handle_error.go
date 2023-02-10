package lib

import "net/http"

type HandleError struct {
	Code     int        `json:"code"`
	Err      string     `json:"error"`
	Messages []Messages `json:"messages,omitempty"`
}

type Messages struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *HandleError) Error() string {
	return e.Err
}

func NewHandleError(code int, err string, messages []Messages) *HandleError {
	return &HandleError{
		Code:     code,
		Err:      err,
		Messages: messages,
	}
}

func NewBadRequestError(err string) *HandleError {
	if err == "" {
		err = "Bad Request"
	}

	return &HandleError{
		Err:  err,
		Code: http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(err string, messages []Messages) *HandleError {
	if err == "" {
		err = "Bad Request"
	}

	return &HandleError{
		Err:      err,
		Code:     http.StatusBadRequest,
		Messages: messages,
	}
}

func NewInternalServerError(err string) *HandleError {
	if err == "" {
		err = "Internal Server Error"
	}

	return &HandleError{
		Err:  err,
		Code: http.StatusInternalServerError,
	}
}
