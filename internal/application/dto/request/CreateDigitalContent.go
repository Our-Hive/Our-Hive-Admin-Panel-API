package request

type CreateDigitalContent struct {
	Title       string `json:"title" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1"`
	URL         string `json:"url" validate:"required,min=1"`
}
