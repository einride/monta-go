package monta

import (
	"net/http"
)

// StatusError represents a HTTP status error from the Monta Partner API.
type StatusError struct {
	Status     string
	StatusCode int
}

func newStatusError(httpResponse *http.Response) error {
	return &StatusError{
		Status:     httpResponse.Status,
		StatusCode: httpResponse.StatusCode,
	}
}

// Error implements error.
func (s *StatusError) Error() string {
	return s.Status
}
