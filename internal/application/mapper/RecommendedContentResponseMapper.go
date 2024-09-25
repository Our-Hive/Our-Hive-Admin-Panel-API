package mapper

import (
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/dto/response"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"
)

func FromRecommendedContentModelToResponse(model *model.DigitalContent) *response.RecommendedContent {
	return &response.RecommendedContent{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		URL:         model.URL,
	}
}

func FromRecommendedContentModelsToResponse(models []*model.DigitalContent) (responseRecommendedContents []*response.RecommendedContent) {
	for _, m := range models {
		responseRecommendedContents = append(responseRecommendedContents, FromRecommendedContentModelToResponse(m))
	}
	return responseRecommendedContents
}
