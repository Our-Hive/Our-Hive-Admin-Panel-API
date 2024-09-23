package api

import "github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/model"

type IDigitalContentServicePort interface {
	CreateDigitalContent(content *model.DigitalContent) (err error)
	GetAllDigitalContent() (content []*model.DigitalContent, err error)
}
