package response

type AllOrders struct {
	CreatedBy string `json:"created_by"`
	Issues    string `json:"issues"`
	Address   string `json:"address"`
	Status    int    `json:"status"`
}
