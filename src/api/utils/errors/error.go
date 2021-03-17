package errors

import "net/http"

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	ApiStatus  int    `json:"status"`
	ApiMessage string `json:"message"`
	ApiError   string `json:"error, omitempty"`
}

func (apiErr *apiError) Status() int {
	return apiErr.ApiStatus
}

func (apiErr *apiError) Message() string {
	return apiErr.ApiMessage
}

func (apiErr *apiError) Error() string {
	return apiErr.ApiError
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		ApiStatus:  statusCode,
		ApiMessage: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		ApiStatus:  http.StatusInternalServerError,
		ApiMessage: message,
	}
}

func NewNotFoundError(message string) ApiError {
	return &apiError{
		ApiStatus:  http.StatusNotFound,
		ApiMessage: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		ApiStatus:  http.StatusBadRequest,
		ApiMessage: message,
	}
}
