package request

type AddNewOrder struct {
	UserId  int    `json:"user_id" field:"required" min:"1"`
	Issues  string `json:"issues" field:"required"`
	Address string `json:"address" field:"required"`
}

type RateOrder struct {
	UserId      int    `json:"user_id" field:"required" min:"1"`
	TechId      int    `json:"tech_id" field:"required" min:"1"`
	Description string `json:"description"`
	Rating      int    `json:"rating" field:"required" min:"1"`
	OrderId     int    `json:"order_id" field:"required" min:"1"`
}

type UpdateOrder struct {
	OrderId   int `json:"order_id" field:"required" min:"1"`
	UserId    int `json:"user_id" field:"required" min:"1"`
	NewStatus int `json:"new_status" field:"required" min:"0"`
	TakenBy   int `json:"taken_by" field:"required" min:"1"`
}
