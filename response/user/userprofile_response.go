package response

type UserResultProfil struct {
	ID       int    `json:"id"`
	IDUser   int    `json:"id_user"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
}
type UserJsonResultProfil struct {
	User UserResultProfil `json:"user"`
}
