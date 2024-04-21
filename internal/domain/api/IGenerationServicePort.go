package api

type IGenerationServicePort interface {
	GenerateImage(prompt string, filename string) (imageUrl string, err error)
}
