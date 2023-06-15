package response

type AllOrders struct {
	CreatedBy string `json:"created_by"`
	Issues    string `json:"issues"`
	Address   string `json:"address"`
	Status    int    `json:"status"`
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
