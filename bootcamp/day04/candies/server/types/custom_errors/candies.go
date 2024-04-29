package customerrors

type CustomError struct {
	HttpStatus int
	Message    string
}

type Error interface {
	Error() CustomError
}

func (err *CustomError) Error() CustomError {
	return *err
}

func NewCustomError(status int, message string) Error {
	return &CustomError{
		HttpStatus: status,
		Message:    message,
	}
}
