package response

type UsertResultUser struct {
	IdUser   int    `json:"id_user"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type ResponeRegister struct {
	User       interface{} `json:"user"`
	Token      string      `json:"token"`
	JwtToken   string      `json:"jwt_token"`
	ServerTime string      `json:"server_time"`
}
