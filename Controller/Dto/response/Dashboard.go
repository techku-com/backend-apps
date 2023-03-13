package response

import "techku/Controller/Dto/request"

type CountDOData struct {
	Unscheduled int `json:"unscheduled"`
	InProgress  int `json:"in_progress"`
	Done        int `json:"done"`
}

type TruckBranchList struct {
	CityId       int    `json:"city_id"`
	CityName     string `json:"city_name"`
	DistrictId   int    `json:"district_id"`
	DistrictName string `json:"district_name"`
}

type TruckTypeList struct {
	TruckTypeId int    `json:"truck_type_id"`
	TruckType   string `json:"truck_type"`
}

type GetListTruckReturn struct {
	List      []ListTruckReturn `json:"list"`
	TotalData int               `json:"total_data"`
}

type ListTruckReturn struct {
	DoId        int              `json:"do_id"`
	DoNo        string           `json:"do_no"`
	TruckType   string           `json:"truck_type"`
	DriverName  string           `json:"driver_name"`
	DriverPhone string           `json:"driver_phone"`
	PoliceNo    string           `json:"police_no"`
	Status      int              `json:"status"`
	UpdatedAt   string           `json:"updated_at"`
	Latitude    string           `json:"latitude"`
	Longitude   string           `json:"longitude"`
	Destination destinationTruck `json:"destination"`
	TruckBranch string           `json:"truck_branch"`
}

type destinationTruck struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Branch    string `json:"branch"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type DashboardUnscheduledDOList struct {
	List      []DashboardUnscheduledDO `json:"list"`
	TotalData int                      `json:"total_data"`
}

type DashboardUnscheduledDO struct {
	baseManagementList
	LogkarNo            string                    `json:"logkar_no"`
	Notes               string                    `json:"notes"`
	TruckRecommendation string                    `json:"truck_recommendation"`
	Origin              request.MasterAddressData `json:"origin"`
	Destination         request.MasterAddressData `json:"destination"`
	GoodsDetail         []OrderListGoodsDetail    `json:"goods_detail"`
}

type doAddress struct {
	Name string `json:"name"`
}

type OnLoadingDOData struct {
	DoId        int       `json:"do_id"`
	DoNo        string    `json:"do_no"`
	Status      int       `json:"status"`
	Origin      doAddress `json:"origin"`
	Destination doAddress `json:"destination"`
	UpdateAt    string    `json:"update_at"`
}

type OnLoadingDOList struct {
	List      []OnLoadingDOData `json:"list"`
	TotalData int               `json:"total_data"`
}

type UsageTruckData struct {
	TruckType  string `json:"truck_type"`
	TotalUsage int    `json:"total_usage"`
	Color      string `json:"color"`
}

type UsageTruckList struct {
	List       []UsageTruckData `json:"list"`
	TotalTruck int              `json:"total_truck"`
	TotalData  int              `json:"total_data"`
}

type ListSummariesDoDone struct {
	SourceDataMap map[string]SummaryDoDone `json:"-"`
	List          []SummaryDoDone          `json:"list"`
	TotalData     int                      `json:"total_data"`
}

type SummaryDoDone struct {
	Date    string `json:"date"`
	TotalDo int    `json:"total_do"`
}
