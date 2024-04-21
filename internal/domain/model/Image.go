package model

import "time"

type Image struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	Size        int64     `json:"size"`
	ContentType string    `json:"content_type"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

func NewImage(ID string, name string, URL string, size int64, contentType string, createdTime time.Time, updatedTime time.Time) *Image {
	return &Image{ID: ID, Name: name, URL: URL, Size: size, ContentType: contentType, CreatedTime: createdTime, UpdatedTime: updatedTime}
}
