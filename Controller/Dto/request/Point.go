package request

type LogkarPointCompleteDODTO struct {
	DoId        int                  `json:"do_id"`
	CompanyId   string               `json:"company_id"`
	DriverId    int                  `json:"driver_id"`
	ConfirmedBy string               `json:"confirmed_by"`
	Penalty     []LogkarPointPenalty `json:"penalty"`
}

type LogkarPointPenalty struct {
	Category int      `json:"category"`
	Reason   []string `json:"reason"`
}

type ExitsTransactionRefId struct {
	TransactionId string `uri:"transaction_id" binding:"required"`
}

type PointDataRequest struct {
	DriverId  int    `json:"driver_id"`
	CompanyId string `json:"company_id"`
	DoId      int    `json:"do_id"`
}
