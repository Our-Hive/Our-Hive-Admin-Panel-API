package domainerror

// ImageAlreadyExistsError

type ImageAlreadyExistsError struct {
	FileName string
}

func (e *ImageAlreadyExistsError) Error() string {
	return "An image with the name " + e.FileName + " already exists"
}
