package domainerror

type DigitalContentAlreadyExistsError struct {
	Title string
}

func (e *DigitalContentAlreadyExistsError) Error() string {
	return "A digital content with the title " + e.Title + " already exists"
}
