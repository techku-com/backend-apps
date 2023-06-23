package response

type AllOrders struct {
	OrderId     int    `json:"order_id"`
	CreatedById int    `json:"created_by_id"`
	CreatedBy   string `json:"created_by"`
	Issues      string `json:"issues"`
	Address     string `json:"address"`
	Status      int    `json:"status"`
	PhoneNumber string `json:"phone_number"`
}

type RateOrder struct {
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	OrderId     int    `json:"order_id"`
}

type UpdatedOrder struct {
	OrderId int `json:"order_id"`
	Status  int `json:"status"`
}
