package api

type IImageGenerationServicePort interface {
	GenerateImage(prompt string) (image []byte, err error)
}
