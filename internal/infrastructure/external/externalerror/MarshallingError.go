package externalerror

type MarshallingError struct {
	Message string
}

func (e *MarshallingError) Error() string {
	return e.Message
}
