package mapper

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/response"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
)

func FromImageToImageResponse(model *model.Image) (responseImage *response.Image) {
	return &response.Image{
		ID:          model.ID,
		Name:        model.Name,
		URL:         model.URL,
		CreatedTime: model.CreatedTime,
		IsApproved:  model.IsApproved,
		ContentType: model.ContentType,
	}
}

func FromImagesToImageResponses(models []*model.Image) (responseImages []*response.Image) {
	for _, m := range models {
		responseImages = append(responseImages, FromImageToImageResponse(m))
	}
	return responseImages
}
