package domainerror

type IsNotEthicalPromptError struct {
	Message string
}

func (e *IsNotEthicalPromptError) Error() string {
	return e.Message
}
