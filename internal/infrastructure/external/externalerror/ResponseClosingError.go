package externalerror

type ResponseClosingError struct {
	Message string
}

func (e *ResponseClosingError) Error() string {
	return e.Message
}
