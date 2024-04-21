package model

type ImageData struct {
	ID       string `json:"id"`
	FileName string `json:"file_name"`
	Data     []byte `json:"data"`
}

func NewImageData(ID string, fileName string, data []byte) *ImageData {
	return &ImageData{ID: ID, FileName: fileName, Data: data}
}
