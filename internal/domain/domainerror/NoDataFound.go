package domainerror

type NoDataFound struct {
	Message string
}

func (e *NoDataFound) Error() string {
	return e.Message
}
