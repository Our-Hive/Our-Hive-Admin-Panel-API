package model

type ContactLine struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewContactLine(id string, name string, description string) *ContactLine {
	return &ContactLine{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
