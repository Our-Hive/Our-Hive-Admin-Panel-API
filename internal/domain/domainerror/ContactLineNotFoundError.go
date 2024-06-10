package domainerror

type ContactLineNotFoundError struct {
	Name string
}

func (e ContactLineNotFoundError) Error() string {
	return "Contact line not found: " + e.Name
}
