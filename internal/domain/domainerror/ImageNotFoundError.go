package domainerror

type ImageNotFoundError struct {
	FileName string
}

func (e ImageNotFoundError) Error() string {
	return "Image not found: " + e.FileName
}
