package response

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegistration struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
