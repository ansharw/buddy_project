package response

import "time"

// all
type ResponseUserContentAll struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Image_name string    `json:"image_name"`
	Body       string    `json:"body"`
	Status     int       `json:"status"`
	Created    time.Time `json:"created"`
	Modified   time.Time `json:"modified"`
}

// create
type ResponseUserContentCreated struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Image_name string    `json:"image_name"`
	Body       string    `json:"body"`
	Status     int       `json:"status"`
	Point      int       `json:"point"`
	Created    time.Time `json:"created"`
}

// detail
type ResponseUserContentByContentId struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Image_name string    `json:"image_name"`
	Body       string    `json:"body"`
	Status     int       `json:"status"`
	Created    time.Time `json:"created"`
	Modified   time.Time `json:"modified"`
}

// update
type ResponseUserContentUpdateByContentId struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Image_name string    `json:"image_name"`
	Body       string    `json:"body"`
	Status     int       `json:"status"`
	Modified   time.Time `json:"modified"`
}
