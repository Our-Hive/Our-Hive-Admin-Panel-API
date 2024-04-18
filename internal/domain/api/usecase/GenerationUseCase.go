package usecase

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/api"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainconstant"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/domainerror"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
	"time"
)

type GenerationUseCase struct {
	imageServicePort                api.IImageServicePort
	imageDataServicePort            api.IImageDataServicePort
	promptClassificationServicePort api.IPromptClassificationServicePort
	imageGenerationServicePort      api.IImageGenerationServicePort
}

func NewGenerationUseCase(imageServicePort api.IImageServicePort, imageDataServicePort api.IImageDataServicePort, promptClassificationServicePort api.IPromptClassificationServicePort, imageGenerationServicePort api.IImageGenerationServicePort) *GenerationUseCase {
	return &GenerationUseCase{imageServicePort: imageServicePort, imageDataServicePort: imageDataServicePort, promptClassificationServicePort: promptClassificationServicePort, imageGenerationServicePort: imageGenerationServicePort}
}

func (g GenerationUseCase) GenerateImage(prompt string, filename string) (imageUrl string, err error) {
	var imageData model.ImageData

	isEthical, err := g.promptClassificationServicePort.IsEthical(prompt)

	if err != nil {
		return "", err
	}

	if !isEthical {
		return "", &domainerror.IsNotEthicalPromptError{Message: domainconstant.PromptIsNotEthicalErrorMessage}
	}

	generatedImage, err := g.imageGenerationServicePort.GenerateImage(prompt)

	imageData.FileName = filename
	imageData.Data = generatedImage

	imageUrl, savedImageData, err := g.imageDataServicePort.SaveImageData(&imageData)

	if err != nil {
		return "", err
	}

	image := model.NewImage(
		savedImageData.ID,
		savedImageData.FileName,
		imageUrl,
		0,
		"image/jpeg",
		time.Now(),
		time.Now(),
	)

	err = g.imageServicePort.SaveImage(image)

	if err != nil {
		return "", err
	}

	return imageUrl, nil
}
