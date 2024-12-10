package externalerror

type ClassificationError struct {
	Message string
}

func (e *ClassificationError) Error() string {
	return e.Message
}
