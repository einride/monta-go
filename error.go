package monta

import (
	"fmt"
	"io"
	"net/http"
)

// StatusError represents a HTTP status error from the Monta Partner API.
type StatusError struct {
	Status     string
	StatusCode int
	Body       string
}

func newStatusError(httpResponse *http.Response) error {
	body, _ := io.ReadAll(httpResponse.Body)

	return &StatusError{
		Status:     httpResponse.Status,
		StatusCode: httpResponse.StatusCode,
		Body:       string(body),
	}
}

// Error implements error.
func (s *StatusError) Error() string {
	if s.Body != "" {
		return fmt.Sprintf("HTTP %d %s: %s", s.StatusCode, s.Status, s.Body)
	}
	return s.Status
}
