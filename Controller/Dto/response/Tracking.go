package response

import "techku/Controller/Dto/request"

type TrackingListData struct {
	TotalData int                  `json:"total_data"`
	Page      int                  `json:"page"`
	Limit     int                  `json:"limit"`
	TruckList []BaseMonitoringData `json:"truck_list"`
}

type TrackingDetailDOData struct {
	LogkarNo          string                    `json:"logkar_no"`
	DoNo              string                    `json:"do_no"`
	RefNo             string                    `json:"ref_no"`
	Waybill           string                    `json:"waybill"`
	DoId              int                       `json:"do_id"`
	LastStatus        int                       `json:"last_status"`
	StartDeliveryDate string                    `json:"start_delivery_date"`
	EndDeliveryDate   string                    `json:"end_delivery_date"`
	DriverName        string                    `json:"driver_name"`
	DriverPhone       string                    `json:"driver_phone"`
	PoliceNo          string                    `json:"police_no"`
	TruckType         string                    `json:"truck_type"`
	TruckCapacity     int                       `json:"truck_capacity"`
	TruckUnit         string                    `json:"truck_unit"`
	TransporterName   string                    `json:"transporter_name"`
	Origin            request.MasterAddressData `json:"origin"`
	Destination       request.MasterAddressData `json:"destination"`
	IsDelivering      bool                      `json:"is_delivering"`
	CargoName         string                    `json:"cargo_name"`
	Goods             []GoodsTrackingData       `json:"goods"`
	DoProgress        []DoProgressTracking      `json:"do_progress"`
}

type GoodsTrackingData struct {
	GoodsName string  `json:"goods_name"`
	Qty       float64 `json:"qty"`
	Unit      string  `json:"unit"`
}

type DoProgressTracking struct {
	Status     int    `json:"status"`
	StatusName string `json:"status_name"`
	DateTime   string `json:"date_time"`
}

type TrackingCombineDOList struct {
	TruckType       string              `json:"truck_type"`
	PoliceNo        string              `json:"police_no"`
	DriverName      string              `json:"driver_name"`
	DriverPhone     string              `json:"driver_phone"`
	TruckCapacity   int                 `json:"truck_capacity"`
	TruckUnit       string              `json:"truck_unit"`
	TransporterName string              `json:"transporter_name"`
	TotalDO         int                 `json:"total_do"`
	DoList          []TrackingCombineDO `json:"do_list"`
}

type TrackingCombineDO struct {
	DoId              int    `json:"do_id"`
	DoNo              string `json:"do_no"`
	LastStatus        int    `json:"last_status"`
	StartDeliveryDate string `json:"start_delivery_date"`
	EndDeliveryDate   string `json:"end_delivery_date"`
	DestinationName   string `json:"destination_name"`
}

type MonitoringTracking struct {
	NewStatus int `json:"new_status"`
	BaseMonitoringData
}

type BaseMonitoringData struct {
	DataCoordinateDTO
	TransporterId string               `json:"transporter_id"`
	Waybill       string               `json:"waybill"`
	PoliceNo      string               `json:"police_no"`
	TruckType     string               `json:"truck_type"`
	DriverName    string               `json:"driver_name"`
	DriverPhone   string               `json:"driver_phone"`
	Origin        regionMonitoring     `json:"origin"`
	DoDetail      []DetailDoMonitoring `json:"do_detail"`
}

type DetailDoMonitoring struct {
	DoId        int              `json:"do_id"`
	DoNo        string           `json:"do_no"`
	Status      int              `json:"status"`
	CargoId     string           `json:"cargo_id"`
	Destination regionMonitoring `json:"destination"`
}

type DataCoordinateDTO struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type regionMonitoring struct {
	Name string `json:"name"`
	DataCoordinateDTO
}

type CheckExistPublicLink struct {
	UniqueLink string `json:"unique_link"`
	DoNo       string `json:"do_no"`
}
