package response

// company
type ResponseRegisteredUser struct {
	UserId   int    `json:"user_id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Point    int    `json:"point"`
}

// company
type ResponseRewardPoint struct {
	UserId   int    `json:"user_id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Point    int    `json:"point"`
}

// user
type ResponseUserPoint struct {
	Point int `json:"point"`
}
