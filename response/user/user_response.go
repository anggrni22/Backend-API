package response

type User struct {
	IDUser       int    `json:"id_user"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	FullName     string `json:"first_name"`
	Password     string `json:"password"`
	IDIdentifier int    `json:"id_identifier"`
}

type InputToken struct {
	IDUser int    `json:"id_user"`
	Token  string `json:"token"`
}

type UserDataRsp struct {
	IdUser   int    `json:"id_user"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
