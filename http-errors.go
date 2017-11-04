package httperror

import (
	"fmt"
	"net/http"
	"strings"
)

// HTTPError is a struct type representing an http error response
type HTTPError struct {
	// The HTTP status code
	Code int
	// The error message
	Message string
}

// New makes a new error given the code. If no message is set, the text from
// http.StatusText() is used. Multiple strings are joined with spaces.
func New(code int, message ...string) *HTTPError {
	if len(message) == 0 {
		return &HTTPError{
			code,
			http.StatusText(code),
		}
	}

	return &HTTPError{
		code,
		strings.Join(message, " "),
	}
}

// IsHTTPError returns the error and true if the provided error is an
// HTTPError. Otherwise, it returns an empty object and false.
func IsHTTPError(err error) (HTTPError, bool) {
	if herr, ok := err.(HTTPError); ok {
		return herr, ok
	}
	if herr, ok := err.(*HTTPError); ok {
		return *herr, ok
	}
	return HTTPError{}, false
}

// Error returns the string representation of this error
func (err HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", err.Code, err.Message)
}

// Respond uses the supplied `http.ResponseWriter` to set the
// approprate headers and a json encoded representation of the
// HTTPError for the body
func (err *HTTPError) Respond(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Header().Set("X-Content-Type-Options", "nosniff")
	res.WriteHeader(err.Code)
	res.Write([]byte(err.Message))
}
