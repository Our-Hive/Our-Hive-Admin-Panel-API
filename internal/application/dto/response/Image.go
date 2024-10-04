package response

import "time"

type Image struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	CreatedTime time.Time `json:"createdTime"`
	IsApproved  bool      `json:"isApproved"`
	ContentType string    `json:"contentType"`
}
