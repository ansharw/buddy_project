package request

type RequestCreateUserContent struct {
	UserId    int `json:"user_id"`
	ContentId int `json:"content_id"`
	Point     int `json:"point"`
}
