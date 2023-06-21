package response

type ArticleList struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type WebStatus struct {
	RegisteredUser  int `json:"registered_user"`
	TotalOrder      int `json:"total_order"`
	TotalRating     int `json:"total_rating"`
	TotalCustomer   int `json:"total_customer"`
	TotalTechnician int `json:"total_technician"`
}
