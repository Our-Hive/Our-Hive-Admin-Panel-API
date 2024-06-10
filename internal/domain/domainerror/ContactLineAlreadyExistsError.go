package domainerror

type ContactLineAlreadyExistsError struct {
	Name string
}

func (e *ContactLineAlreadyExistsError) Error() string {
	return "A contact line with the name " + e.Name + " already exists"
}
