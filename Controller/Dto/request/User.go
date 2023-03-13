package request

type ReqDataUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserAuthMobileV2 struct {
	NoPhone         string `json:"no_phone"`
	DeviceToken     string `json:"device_token"`
	Token           string `json:"token"`
	CompanyId       string `json:"company_id"`
	CompanyName     string `json:"company_name"`
	DriverCreatedAt string `json:"driver_created_at"`
	UserAuthAndOrder
}

type UserAuthAndOrder struct {
	DriverOnGoingDo
	ProfileAndStatus
}

type ProfileAndStatus struct {
	Puk             string  `json:"puk"`
	DriverId        string  `json:"driver_id"`
	DriverName      string  `json:"driver_name"`
	Status          string  `json:"driver_status"`
	DriverPicture   *string `json:"driver_picture"`
	DriverSignature string  `json:"driver_sign"`
}

type UserAuthIdAndPuk struct {
	Puk      string `json:"puk"`
	DriverId string `json:"driver_id"`
}

type DriverOnGoingDo struct {
	OnGoingWaybill        string `json:"on_going_waybill"`
	OnGoingDoParent       string `json:"on_going_do_parent"`
	OnGoingDoParentId     string `json:"on_going_do_parent_id"`
	OnGoingDoNo           string `json:"on_going_do_no"`
	OnGoingStatus         string `json:"on_going_status"`
	OnGoingPoliceNo       string `json:"on_going_police_no"`
	OnGoingDoParentStatus string `json:"on_going_do_parent_status"`
	OnGoingDoId           string `json:"on_going_do_id"`
	OnGoingCargoId        string `json:"on_going_cargo_id"`
	IsCombine             string `json:"is_combine"`
}
