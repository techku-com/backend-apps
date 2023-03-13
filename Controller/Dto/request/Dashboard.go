package request

type ParamTruckReturn struct {
	UserAuth
	Page           int    `json:"page"`
	Keywords       string `json:"keywords"`
	Status         []int  `json:"status"`
	TruckTypeId    []int  `json:"truck_type_id"`
	BranchLocation []int  `json:"branch_location"`
}

type ParamSummariesDoDone struct {
	UserAuth
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	DateDiff  int64  `json:"date_diff"`
}

type DoPerformance struct {
	DoNo       string `json:"do_no"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	DoPoint    int    `json:"do_point"`
	DateCreate string `json:"date_create"`
	DoId       int    `json:"do_id"`
}

type SchedulingPerformance struct {
	Date       string          `json:"date"`
	TotalPoint int             `json:"total_point"`
	DoList     []DoPerformance `json:"do_list"`
}
