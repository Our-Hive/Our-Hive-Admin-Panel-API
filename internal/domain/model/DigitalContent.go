package model

type DigitalContent struct {
	ID          int
	Title       string
	Description string
	URL         string
}

func NewDigitalContent(title string, description string, URL string) *DigitalContent {
	return &DigitalContent{Title: title, Description: description, URL: URL}
}
