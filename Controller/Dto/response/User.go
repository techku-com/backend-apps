package response

type UserLogin struct {
	UserId   int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserRegistration struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
