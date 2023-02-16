package request

type RequestRegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type RequestLoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
