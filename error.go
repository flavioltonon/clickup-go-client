package clickup

import "errors"

type Error struct {
	StatusCode int
	Message    string
}

func (e Error) Error() string {
	return e.Message
}

func NewError(statusCode int, message string) error {
	return Error{
		StatusCode: statusCode,
		Message:    message,
	}
}

var (
	// ErrInvalidRequestBody is an error that should be returned if the body of a request is not valid
	ErrInvalidRequestBody = errors.New("invalid request body")

	// ErrInvalidRequestSignature is an error that should be returned if a Webhook event request signature is not valid
	ErrInvalidRequestSignature = errors.New("invalid request signature")
)
