package request

type UserAuth struct {
	CompanyID     string  `json:"company_id,omitempty"`
	UserID        string  `json:"user_id,omitempty"`
	CompanyName   string  `json:"company_name,omitempty"`
	CompanyType   string  `json:"company_type,omitempty"`
	RoleId        string  `json:"role_id,omitempty"`
	Token         string  `json:"token,omitempty"`
	CompanyConfig *string `json:"company_config,omitempty"`
}

type IdParam struct {
	Id int `uri:"id" binding:"required,min=1"`
}
