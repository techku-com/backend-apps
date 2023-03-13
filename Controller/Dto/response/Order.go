package response

import (
	"techku/Controller/Dto/request"
)

type DeliveryStatusOrder struct {
	Scheduled  int `json:"scheduled"`
	InProgress int `json:"in_progress"`
	Done       int `json:"done"`
	Cancelled  int `json:"cancelled"`
	All        int `json:"all"`
}

type DoDetailsSuccess struct {
	DoNo               string                  `json:"do_no"`
	FinishDeliveryDate string                  `json:"finish_delivery_date"`
	ReffNo             string                  `json:"ref_no"`
	LogkarNo           string                  `json:"logkar_no"`
	Waybill            string                  `json:"waybill"`
	GoodsDetail        []goodsData             `json:"goods_detail"`
	LogkarPoint        LogkarPointData         `json:"logkar_point"`
	Origin             DoDetailGoodsAndAddress `json:"origin"`
	Destination        DoDetailGoodsAndAddress `json:"destination"`
}

type DoDetailGoodsAndAddress struct {
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	DetailAddress string   `json:"detail_address"`
	Documents     []string `json:"documents"`
	TotalWeightKg float64  `json:"total_weight_kg"`
	weightBridgeData
}

type goodsData struct {
	Name        string  `json:"name"`
	LoadingQty  float64 `json:"loading_qty"`
	ReceivedQty float64 `json:"received_qty"`
	Unit        string  `json:"unit"`
}

type LogkarPointData struct {
	PointStatus   int      `json:"point_status"`
	TotalPoint    int      `json:"total_point"`
	ReceivedPoint int      `json:"received_point"`
	Reason        []string `json:"reason"`
}

type LogkarPointDetail struct {
	AllowReceivePoint bool   `json:"allow_receieve_point"`
	CalculatedPoint   int    `json:"calculated_point"`
	DoOwnership       int    `json:"do_ownership"`
	DriverName        string `json:"driver_name"`
	TransporterName   string `json:"transporter_name"`
	DoNo              string `json:"do_no"`
	CargoName         string `json:"cargo_name"`
	FinishTime        string `json:"finish_time"`
}

type baseManagementList struct {
	DoIdParent        int    `json:"do_id_parent,omitempty"`
	DoId              int    `json:"do_id"`
	DoNo              string `json:"do_no"`
	RefNo             string `json:"ref_no"`
	StartDeliveryDate string `json:"start_delivery_date"`
	EndDeliveryDate   string `json:"end_delivery_date"`
	Status            int    `json:"status"`
	CargoName         string `json:"cargo_name"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	SameDay           bool   `json:"same_day"`
	DoParent          string `json:"do_parent"`
}

type ListSchedule struct {
	List      []ScheduleList `json:"list"`
	TotalData int            `json:"total_data"`
}

type ScheduleList struct {
	baseManagementList
	DriverName  string                    `json:"driver_name"`
	PoliceNo    string                    `json:"police_no"`
	TruckType   string                    `json:"truck_type"`
	Origin      DeliveryAddress           `json:"origin"`
	Destination DeliveryAddress           `json:"destination"`
	GoodsDetail []ScheduleListGoodsDetail `json:"goods_detail"`
}

type ScheduleListGoodsDetail struct {
	listDoGoodsDetail
	GoodsId       int     `json:"goods_id"`
	TotalWeightKg float64 `json:"total_weight_kg"`
	CommodityId   int     `json:"commodity_id"`
}

type UnscheduledOrderList struct {
	List      []UnscheduledOrder `json:"list"`
	TotalData int                `json:"total_data"`
}

type UnscheduledOrder struct {
	baseManagementList
	LogkarNo            string                    `json:"logkar_no"`
	TotalSubDo          int                       `json:"total_sub_do"`
	Notes               *string                   `json:"notes"`
	TruckRecommendation string                    `json:"truck_recommendation"`
	Origin              request.MasterAddressData `json:"origin"`
	Destination         request.MasterAddressData `json:"destination"`
	GoodsDetail         []OrderListGoodsDetail    `json:"goods_detail"`
}

type OrderListGoodsDetail struct {
	listDoGoodsDetail
	DeliveryQty float64 `json:"delivery_qty"`
}

type listDoGoodsDetail struct {
	Code         string  `json:"code"`
	Unit         string  `json:"unit"`
	GoodsName    string  `json:"goods_name"`
	OrderQty     float64 `json:"order_qty"`
	UnitWeightKg float64 `json:"unit_weight_kg"`
}

type ScheduleChildDoList struct {
	DoNo          string  `json:"do_no"`
	DoId          int     `json:"do_id"`
	TotalWeightKg float64 `json:"total_weight_kg"`
	Status        int     `json:"status"`
}

type CombineFeatureDo struct {
	DriverId    int          `json:"-"`
	ListDoNo    map[int]bool `json:"-"`
	DataCombine []CombineListDo
}

type CombineListDo struct {
	DoId            int                       `json:"do_id"`
	DoNo            string                    `json:"do_no"`
	Goods           []ScheduleListGoodsDetail `json:"goods"`
	DestinationName string                    `json:"destination_name"`
}

type OrderDatum struct {
	DoId            int    `json:"do_id"`
	Status          int    `json:"status"`
	WayBill         string `json:"waybill"`
	DoNo            string `json:"do_no"`
	OriginName      string `json:"origin_name"`
	DestinationName string `json:"destination_name"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	SameDay         bool   `json:"same_day"`
	UpdatedAt       string `json:"updated_at"`
}
type DeliveryOrderDetails struct {
	WayBill  string                      `json:"waybill"`
	Delivery []DeliveryOrderDetailsDatum `json:"delivery"`
}

type DeliveryOrderDetailsDatum struct {
	DoId        int               `json:"do_id"`
	DoNo        string            `json:"do_no"`
	RefNo       string            `json:"ref_no"`
	LogkarNo    string            `json:"logkar_no"`
	StartDate   string            `json:"start_date"`
	EndDate     string            `json:"end_date"`
	SameDay     bool              `json:"same_day"`
	Notes       *string           `json:"notes"`
	Origin      DeliveryAddress   `json:"origin"`
	Destination DeliveryAddress   `json:"destination"`
	Goods       []GoodsDatumAttrs `json:"goods"`
}

type GoodsDatumAttrs struct {
	GoodsId      int     `json:"goods_id"`
	Name         string  `json:"name"`
	Qty          float64 `json:"qty"`
	Unit         string  `json:"unit"`
	UnitWeightKg float64 `json:"unit_weight_kg"`
}

type DeliveryAddress struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	DetailAddress string `json:"detail_address"`
	PICPhone      string `json:"pic_phone"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
}

type DODataParent struct {
	CompanyIdOrderSource   int               `json:"company_id_order_source"`
	CompanyOrderSourceName string            `json:"company_order_source_name"`
	Order                  request.OrderData `json:"order"`
}
type SummaryOrderByRefNo struct {
	DoId          int                 `json:"do_id"`
	DoNo          string              `json:"do_no"`
	RefNo         string              `json:"ref_no"`
	DriverName    *string             `json:"driver_name"`
	TruckType     *string             `json:"truck_type"`
	PoliceNo      *string             `json:"police_no"`
	Status        int                 `json:"status"`
	SubDo         int                 `json:"sub_do"`
	StartDelivery string              `json:"start_delivery"`
	EndDelivery   string              `json:"end_delivery"`
	Goods         []summaryGoodsRefNo `json:"goods"`
}

type summaryGoodsRefNo struct {
	summaryGoods
	DeliveryQty float64 `json:"delivery_qty"`
}

type summaryGoods struct {
	Code         string  `json:"code"`
	Unit         string  `json:"unit"`
	GoodsName    string  `json:"goods_name"`
	OrderQty     float64 `json:"order_qty"`
	UnitWeightKg float64 `json:"unit_weight_kg"`
	Received     float64 `json:"received_qty"`
}

type DOParentSummary struct {
	summaryBaseStruct
	GoodsDetail []OrderListGoodsDetail `json:"goods_detail"`
	Child       []ChildDOData          `json:"child"`
}

type ChildDOData struct {
	DoId                  int                        `json:"do_id"`
	DoNo                  string                     `json:"do_no"`
	DriverName            string                     `json:"driver_name"`
	TruckType             string                     `json:"truck_type"`
	PoliceNo              string                     `json:"police_no"`
	Status                int                        `json:"status"`
	TotalOrderGoods       float64                    `json:"total_order_goods"`
	TotalReceivedGoods    float64                    `json:"total_received_goods"`
	TotalWeightDifference float64                    `json:"total_weight_difference"`
	TotalLoadingGoods     float64                    `json:"total_loading_goods"`
	Goods                 []summariesGoodsDataParent `json:"goods"`
}

type summariesGoodsDataParent struct {
	summaryGoods
	Unit       string  `json:"unit"`
	Difference float64 `json:"difference"`
	LoadingQty float64 `json:"loading_qty"`
}

type GetSummaryDoNo struct {
	summaryBaseStruct
	Waybill          string                     `json:"waybill"`
	DriverName       string                     `json:"driver_name"`
	PoliceNo         string                     `json:"police_no"`
	TruckType        string                     `json:"truck_type"`
	TruckCapacity    int                        `json:"truck_capacity"`
	TruckUnit        string                     `json:"truck_unit"`
	PodReceivedNotes string                     `json:"pod_received_notes"`
	DnReceivedNotes  string                     `json:"dn_received_notes"`
	PodReceivedBy    string                     `json:"pod_received_by"`
	DnReceivedBy     string                     `json:"dn_received_by"`
	DnReceivedDate   string                     `json:"dn_received_date"`
	PodReceivedDate  string                     `json:"pod_received_date"`
	Documents        []DocumentMedia            `json:"documents"`
	WeightBridge     summariesWeightBridge      `json:"weight_bridge"`
	TransactionLog   []BaseLogData              `json:"transaction_log"`
	Goods            []summariesGoodsDataParent `json:"goods"`
	DoNo             string                     `json:"do_no"`
}

type summaryBaseStruct struct {
	DoIdParent      int             `json:"do_id_parent,omitempty"`
	DoId            int             `json:"do_id,omitempty"`
	DoParent        string          `json:"do_parent"`
	RefNo           string          `json:"ref_no"`
	StartDate       string          `json:"start_date"`
	EndDate         string          `json:"end_date"`
	CargoName       string          `json:"cargo_name"`
	TransporterName string          `json:"transporter_name"`
	LogkarNo        string          `json:"logkar_no"`
	Status          int             `json:"status"`
	Origin          DeliveryAddress `json:"origin"`
	Destination     DeliveryAddress `json:"destination"`
}

type summariesWeightBridge struct {
	Dn  weightBridgeData `json:"dn"`
	Pod weightBridgeData `json:"pod"`
}

type weightBridgeData struct {
	Bruto float64 `json:"bruto"`
	Netto float64 `json:"netto"`
	Tara  float64 `json:"tara"`
}

type DocumentMedia struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

type BaseLogData struct {
	Status      int    `json:"status"`
	StatusValue string `json:"status_value,omitempty"`
	Latitude    string `json:"latitude,omitempty"`
	Longitude   string `json:"longitude,omitempty"`
	CreatedAt   string `json:"created_at"`
}

type DODataForUpdateStatus struct {
	DoParentId   int              `json:"do_parent_id"`
	Waybill      string           `json:"waybill"`
	CompanyID    int              `json:"company_id"`
	Goods        []GoodsFromSubDO `json:"goods"`
	ParentStatus int              `json:"parent_status"`
	Status       int              `json:"status"`
	DriverId     int              `json:"driver_id"`
	DoNo         string           `json:"do_no"`
	PoliceNo     string           `json:"police_no"`
	DoId         int              `json:"do_id"`
	AssignedAt   string           `json:"created_at"`
}

type GoodsFromSubDO struct {
	Code         string  `json:"code"`
	OrderQty     float64 `json:"order_qty"`
	UnitWeightKg float64 `json:"unit_weight_kg"`
}

type AllDriverDOHistoryList struct {
	List      []OrderDatum `json:"list"`
	TotalData int          `json:"total_data"`
}

type AllDriverDOHistoryData struct {
	Month int `json:"month"`
	Year  int `json:"year"`
	Total int `json:"total"`
}

type DoByDoNo struct {
	summaryBaseStruct
	DoNo        string        `json:"do_no"`
	CreatedAt   string        `json:"created_at"`
	GoodsDetail []interface{} `json:"goods_detail"`
}
