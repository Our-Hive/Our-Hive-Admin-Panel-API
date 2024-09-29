package response

import "time"

type Image struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	CreatedTime time.Time `json:"created_time"`
	IsApproved  bool      `json:"is_approved"`
	ContentType string    `json:"content_type"`
}
