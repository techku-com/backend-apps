package request

type TrackingListParams struct {
	UserAuth
	Page int `json:"page"`
}

type TrackingDetailDODTO struct {
	UserAuth
	DoId int `json:"do_id"`
}

type TrackingCombineListDTO struct {
	UserAuth
	PoliceNo string `json:"police_no"`
}

type DataRequestPositionOrStatusEvent struct {
	DataCoordinateDTO
	CompanyId           int    `json:"company_id"`
	Waybill             string `json:"waybill"`
	DoNo                string `json:"do_no"`
	IsUpdatePosition    bool   `json:"is_update_position"`
	NewStatus           int    `json:"new_status"`
	LastUpdatedPosition string `json:"last_updated_position"`
}

type DataCoordinateDTO struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type GetDOProgressTracking struct {
	DoId int `uri:"do_id" binding:"required"`
}

type CheckPublicLinkParam struct {
	Link string `uri:"link" binding:"required"`
}

type GetLastPositionDO struct {
	DoId int `uri:"do_id" binding:"required"`
}
