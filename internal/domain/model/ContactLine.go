package model

type ContactLine struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewContactLine(name string, description string) *ContactLine {
	return &ContactLine{Name: name, Description: description}
}
