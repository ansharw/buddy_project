package request

type RequestCreateContent struct {
	Title      string `json:"title"`
	Image_name string `json:"image_name"`
	Body       string `json:"body"`
	Status     int    `json:"status"`
}

type RequestUpdateContent struct {
	Title      string `json:"title"`
	Image_name string `json:"image_name"`
	Body       string `json:"body"`
	Status     int    `json:"status"`
}
