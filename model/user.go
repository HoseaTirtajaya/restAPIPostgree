package model

//Init user
type User struct {
	ID       string  `json:"id_user"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Token    *string `json:"token"`
	RoleID   string  `json:"role_id"`
}
