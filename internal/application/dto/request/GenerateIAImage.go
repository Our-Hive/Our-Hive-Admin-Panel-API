package request

type GenerateIAImage struct {
	Prompt   string `json:"prompt" validate:"required,min=1"`
	FileName string `json:"file_name" validate:"required,min=1"`
}
