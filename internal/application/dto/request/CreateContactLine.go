package request

type CreateContactLine struct {
	Name        string `json:"name" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1"`
}
