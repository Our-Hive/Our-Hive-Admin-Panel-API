package externalerror

type RequestSendingError struct {
	Message string
}

func (e *RequestSendingError) Error() string {
	return e.Message
}
