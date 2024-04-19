package request

type GenerateIAImage struct {
	Prompt string `json:"prompt"`
	Name   string `json:"name"`
}
