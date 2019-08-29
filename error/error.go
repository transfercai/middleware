package error

type SError struct {
	error
	Message   string
	ErrorCode int
}

func NewSError(errorCode int, message string, err error) *SError {
	return &SError{
		error:     err,
		Message:   message,
		ErrorCode: errorCode,
	}
}
