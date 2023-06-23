package response

type UserLogin struct {
	UserId   int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
	Username string `json:"username"`
}

type MyOrderList struct {
	OrderId         int         `json:"order_id"`
	CreatedBy       int         `json:"created_by"`
	TakenBy         int         `json:"taken_by"`
	CreatedByName   string      `json:"created_by_name"`
	TakenByName     string      `json:"taken_by_name"`
	Status          int         `json:"status"`
	Issues          string      `json:"issues"`
	CreatedAt       string      `json:"created_at"`
	Rating          OrderRating `json:"order_rating"`
	CustomerPhone   string      `json:"customer_phone"`
	TechnicianPhone string      `json:"technician-phone"`
}

type OrderRating struct {
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}

type UserRegistration struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
