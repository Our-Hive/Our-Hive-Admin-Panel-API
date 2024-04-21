package externalerror

type RequestCreationError struct {
	Message string
}

func (e *RequestCreationError) Error() string {
	return e.Message
}
