package request

type UserLogin struct {
	Email    string `json:"email" field:"required"`
	Password string `json:"password" field:"required" min:"7"`
}

type UserGoogleAuthentication struct {
	GoogleJWT *string `json:"google_token"`
}

type ReqDataUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegistration struct {
	Username string `json:"username" field:"required"`
	Email    string `json:"email" field:"required"`
	Password string `json:"password" field:"required" min:"7"`
	Role     bool   `json:"role"`
	Phone    string `json:"phone" field:"required"`
}
