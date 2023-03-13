package response

type DataFromSubmitLogkarPoint struct {
	TransactionId string `json:"transaction_id"`
	DoNo          string `json:"do_no"`
	ReceivedPoint int    `json:"received_point"`
	FinalPoint    int    `json:"final_point"`
	PointStatus   string `json:"point_status"`
	DriverDataBalance
}

type DriverDataBalance struct {
	DriverId        int     `json:"driver_id"`
	DriverName      string  `json:"driver_name"`
	DriverCompanyId int     `json:"driver_company_id"`
	CompanyName     string  `json:"company_name"`
	BalancePoint    float64 `json:"balance_point"`
	RefSource       string  `json:"ref_source"`
	Description     string  `json:"description"`
	RefId           string  `json:"ref_id"`
}

type GetInfoPointDo struct {
	DoDetail          GetSummaryDoNo  `json:"do_detail"`
	LogkarPointStatus InfoStatusPoint `json:"logkar_point_status"`
	LogkarPointInfo   InfoLogkarPoint `json:"logkar_point_info"`
}

type GetLogkarPointData struct {
	LogkarPointStatus InfoStatusPoint `json:"logkar_point_status"`
	LogkarPointInfo   InfoLogkarPoint `json:"logkar_point_info"`
}

type InfoLogkarPoint struct {
	TruckCapacity            int               `json:"truck_capacity"`
	Eta                      int               `json:"eta"`
	TruckCapacityCalculation float64           `json:"truck_capacity_calculation"`
	EtaCalculation           float64           `json:"eta_calculation"`
	TotalPoint               int               `json:"total_point"`
	PenaltyCategory          []PenaltyCategory `json:"penalty_category"`
}

type PenaltyCategory struct {
	CategoryCode int    `json:"category_code"`
	Category     string `json:"category"`
	Percentage   int    `json:"percentage"`
}

type InfoStatusPoint struct {
	DoNo       string   `json:"do_no"`
	DriverName string   `json:"driver_name"`
	StatusCode int      `json:"status_code"`
	StatusName string   `json:"status_name"`
	Point      int      `json:"point"`
	SubmitBy   string   `json:"submit_by"`
	SubmitDate string   `json:"submit_date"`
	Reason     string   `json:"reason"`
	ReasonData []string `json:"-"`
}
