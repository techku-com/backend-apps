package request

import (
	"github.com/lib/pq"
)

type DeliveryStatusOrderDTO struct {
	UserAuth
}

type ListDoBulkDTO struct {
	DataList  []string `json:"data_list"`
	CompanyId int      `json:"company_id"`
}

type AllOrderData struct {
	CompanyIdOrderSource   int         `json:"company_id_order_source"`
	CompanyOrderSourceName string      `json:"company_order_source_name"`
	Orders                 []OrderData `json:"orders"`
}

type OrderData struct {
	CompanyIdTransporter   int               `json:"company_id_transporter"`
	CompanyTransporterName string            `json:"company_transporter_name"`
	CompanyIdCargo         int               `json:"company_id_cargo"`
	CompanyCargoName       string            `json:"company_cargo_name"`
	DoId                   int               `json:"do_id"`
	DoNo                   string            `json:"do_no"`
	RefNo                  string            `json:"ref_no"`
	LogkarNo               string            `json:"logkar_no"`
	StartDeliveryDate      string            `json:"start_delivery_date"`
	EndDeliveryDate        string            `json:"end_delivery_date"`
	TruckRecommendation    string            `json:"truck_recommendation"`
	Origin                 MasterAddressData `json:"origin"`
	Destination            MasterAddressData `json:"destination"`
	Goods                  []MasterGoodsData `json:"goods"`
	Eta                    int               `json:"eta"`
	EstimateDistance       int               `json:"estimate_distance"`
	Notes                  string            `json:"notes"`
	Parent                 string            `json:"parent"`
	Waybill                string            `json:"waybill"`
	Status                 int               `json:"status"`
	Truck                  TruckDOData       `json:"truck"`
	DriverData
}

type ScheduleListParamDTO struct {
	UserAuth
	Type     int    `json:"type"`
	Page     int    `json:"page"`
	Keywords string `json:"keywords"`
}

type UnscheduledOrderListDTO struct {
	UserAuth
	Page     int    `json:"page"`
	Keywords string `json:"keywords"`
}

type GetCombineListDTO struct {
	UserAuth
	DoId int `json:"do_id"`
}

type SubmitCombineDoDTO struct {
	UserAuth
	DoList  []int  `json:"do_list"`
	Waybill string `json:"waybill"`
}

type OrderWayBill struct {
	WayBill string `json:"waybill"`
	UserAuthMobileV2
}

type GetDOParentDTO struct {
	UserAuth
	DoNo string `json:"do_no"`
	DoId int    `json:"do_id"`
}

type TruckDOData struct {
	TruckTypeId   int    `json:"truck_type_id"`
	TruckType     string `json:"truck_type"`
	Unit          string `json:"unit"`
	PoliceNo      string `json:"police_no"`
	TruckCapacity int    `json:"truck_capacity"`
	ProvinceId    int    `json:"province_id"`
	ProvinceName  string `json:"province_name"`
	CityId        int    `json:"city_id"`
	CityName      string `json:"city_name"`
	DistrictId    int    `json:"district_id"`
	DistrictName  string `json:"district_name"`
}

type DriverData struct {
	DriverID    int     `json:"driver_id"`
	DriverName  string  `json:"driver_name"`
	DriverPhone string  `json:"driver_phone"`
	DeviceID    *string `json:"device_id"`
	DevicePUK   *string `json:"device_puk"`
}

type UpdateDOStatusDTO struct {
	DoPropertiesForUpdate
	CompanyId       string `json:"company_id"`
	NewStatus       int    `json:"new_status"`
	CurrentStatus   int    `json:"current_status"`
	DriverId        int    `json:"driver_id"`
	Puk             string `json:"puk"`
	DriverSignature string `json:"driver_sign"`
}

type UpdateStatusParentDO struct {
	UpdateDOStatusDTO
	SubDo                 bool            `json:"sub_do"`
	SumParentQty          float64         `json:"sum_parent_qty"`
	SchedulingPerformance []DoPerformance `json:"scheduling_performance"`
}

type GetSummaryDo struct {
	CompanyID string `json:"company_id"`
	DoId      int    `json:"do_id"`
}

type GetSummaryDoRefDTO struct {
	CompanyID string `json:"company_id"`
	RefNo     string `json:"ref_no"`
}

type MobileDocumentSubmit struct {
	UserAuthMobileV2
	BaseDocumentDeliveryData
	DocumentProcessData
	DocumentType string `json:"document_type"`
	UpdateDoData UpdateDOStatusDTO
}

type DocumentProcessData struct {
	DocumentData   DocumentDeliveryLogData
	DocumentImages []PushFileToServer
}

type DocumentDeliveryLogData struct {
	BaseDocumentDeliveryData
	Documents      []DocumentMedia `json:"documents"`
	IsWeightBridge bool            `json:"-"`
	NewStatus      int             `json:"-"`
	CurrentStatus  int             `json:"-"`
	CompanyId      string          `json:"-"`
	DoParentId     int             `json:"-"`
	DoParentStatus int             `json:"-"`
	DriverId       int             `json:"-"`
}
type BaseDocumentDeliveryData struct {
	Session      string            `json:"session"`
	WeightBridge weightBridgeData  `json:"weight_bridge"`
	Notes        string            `json:"notes"`
	ReceivedName string            `json:"received_name"`
	Goods        []MasterGoodsData `json:"goods"`
	DoPropertiesForUpdate
}

type weightBridgeData struct {
	Bruto float64 `json:"bruto"`
	Netto float64 `json:"netto"`
	Tara  float64 `json:"tara"`
}

type DocumentMedia struct {
	Path     string `json:"path"`
	Type     string `json:"type"`
	Filename string `json:"-"`
}

type PushFileToServer struct {
	BaseFile  string `json:"base_file"`
	NameFile  string `json:"name_file"`
	Directory string `json:"directory"`
}

type UpdateTransporterStatusDTO struct {
	DriverId  int    `json:"driver_id"`
	PoliceNo  string `json:"police_no"`
	CompanyId string `json:"company_id"`
	OldStatus int    `json:"old_status"`
	NewStatus int    `json:"new_status"`
}

type OrderCheckinData struct {
	DoPropertiesForUpdate
	UserAuthMobileV2
}

type CreateDriverSignatureForSPK struct {
	Filename       string `json:"filename"`
	OnGoingWaybill string `json:"on_going_waybill"`
}

type LogTransactionDo struct {
	BulkDoId pq.Int64Array `json:"bulk_do_id"`
	DoId     int           `json:"do_id"`
	Status   int           `json:"status"`
	DataCoordinate
}

type DataCoordinate struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type DoPropertiesForUpdate struct {
	Waybill       string `json:"waybill,omitempty"`
	DoNo          string `json:"do_no,omitempty"`
	IsNextDoExist bool   `json:"is_next_do_exist,omitempty"`
	DoId          int    `json:"do_id"`
	DataCoordinate
	Notes       string `json:"notes,omitempty"`
	IsCombineDo string `json:"is_combine_do,omitempty"`
}

type CancelOrDeleteDTO struct {
	UserAuth
	DoId  int    `json:"do_id"`
	Notes string `json:"notes"`
}

type ReassignDODTO struct {
	UserAuth
	DoId   int         `json:"do_id"`
	Truck  TruckDOData `json:"truck"`
	Driver DriverData  `json:"driver"`
}

type HistoryDODTO struct {
	UserAuthMobileV2
	Page     int    `json:"page"`
	Type     int    `json:"type"`
	Search   string `json:"search"`
	Ordering int    `json:"ordering"`
}

type ValidateIdToAssign struct {
	DoId  int  `json:"do_id"`
	SubDo bool `json:"sub_do"`
}

type InitLogkarPoint struct {
	DoId     int `json:"do_id"`
	DriverId int `json:"driver_id"`
}

type GetDoByDoNo struct {
	UserAuth
	DoNo string `json:"do_no"`
}

type DetailCancelledDO struct {
	DoNo            string        `json:"do_no"`
	Status          int           `json:"status"`
	StartDate       string        `json:"start_delivery_date"`
	EndDate         string        `json:"end_delivery_date"`
	RefNo           string        `json:"ref_no"`
	LogkarNo        string        `json:"logkar_no"`
	Waybill         string        `json:"waybill"`
	TransporterName string        `json:"transporter_name"`
	Notes           string        `json:"notes"`
	CancelledAt     string        `json:"cancelled_at"`
	SameDay         bool          `json:"same_day"`
	Goods           []detailGoods `json:"goods"`
	Origin          detailAddress `json:"origin"`
	Destination     detailAddress `json:"destination"`
}
type detailAddress struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	DetailAddress string `json:"detail_address"`
}
type detailGoods struct {
	Name         string  `json:"name"`
	Qty          float64 `json:"qty"`
	Unit         string  `json:"unit"`
	UnitWeightKg float64 `json:"unit_weight_kg"`
}
