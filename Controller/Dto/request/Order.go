package request

type AddNewOrder struct {
	UserId  int    `json:"user_id" field:"required" min:"1"`
	Issues  string `json:"issues" field:"required"`
	Address string `json:"address" field:"required"`
}
