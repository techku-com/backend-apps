package response

type ReportingHtmlPdf struct {
	SPK              SPKPdf              `json:"spk"`
	SummaryDoParent  SummaryDoParentPdf  `json:"summary_do_parent"`
	SummaryCombineDo SummaryCombineDoPdf `json:"summary_combine_do"`
	SummaryDoRef     SummaryDoRefPdf     `json:"summary_do_ref"`
	SummaryDo        SummaryDoPdf        `json:"summary_do"`
}

type SPKPdf struct {
	Host    string `json:"host"`
	Token   string `json:"token"`
	SpkDate string `json:"spk_date"`
	baseDataPdf
	Origin       region   `json:"origin"`
	Destination  region   `json:"destination"`
	DNRecipient  string   `json:"dn_recipient"`
	PODRecipient string   `json:"pod_recipient"`
	DN           document `json:"dn"`
	POD          document `json:"pod"`
}

type SummaryDoParentPdf struct {
	Host   string     `json:"host"`
	Header HeaderHtml `json:"header"`
	baseDataPdf
	TotalDo            int                `json:"total_do"`
	Goods              []goodsDetailPdf   `json:"goods"`
	Origin             region             `json:"origin"`
	Destination        region             `json:"destination"`
	SummaryGoodsWeight summaryGoodsWeight `json:"summary_goods_weight"`
	DoDetails          []DoDetails        `json:"do_details"`
}

type DoDetails struct {
	baseDataPdf
	summaryGoodsWeight
	Goods []goodsDetailPdf `json:"goods"`
}

type SummaryCombineDoPdf struct {
	Host    string     `json:"host"`
	Header  HeaderHtml `json:"header"`
	TotalDo int        `json:"total_do"`
	baseDataPdf
	SummaryGoodsWeight summaryGoodsWeight `json:"summary_goods_weight"`
	Origin             region             `json:"origin"`
	DoDetails          []DoDetailsCombine `json:"do_details"`
}

type DoDetailsCombine struct {
	baseDataPdf
	summaryGoodsWeight
	Destination region           `json:"destination"`
	Goods       []goodsDetailPdf `json:"goods"`
}

type SummaryDoRefPdf struct {
	Host               string             `json:"host"`
	Header             HeaderHtml         `json:"header"`
	TotalDo            int                `json:"total_do"`
	SummaryGoodsWeight summaryGoodsWeight `json:"summary_goods_weight"`
	DoDetails          []DoDetailByRefNo  `json:"do_details"`
}

type DoDetailByRefNo struct {
	baseDataPdf
	summaryGoodsWeight
	Origin      region           `json:"origin"`
	Destination region           `json:"destination"`
	Goods       []goodsDetailPdf `json:"goods"`
}

type SummaryDoPdf struct {
	baseDataPdf
	summaryGoodsWeight
	Header   HeaderHtml   `json:"header"`
	Host     string       `json:"host"`
	Token    string       `json:"token"`
	Logs     LogsData     `json:"logs"`
	Loading  deliveryData `json:"loading"`
	Demolish deliveryData `json:"demolish"`
}

type deliveryData struct {
	region
	document
	Goods           []goodsDetailPdf `json:"goods"`
	ReceivedBy      string           `json:"received_by"`
	Weighbridge     bool             `json:"weighbridge"`
	Bruto           float64          `json:"bruto"`
	Tara            float64          `json:"tara"`
	Netto           float64          `json:"netto"`
	TotalOrderGoods float64          `json:"total_order_goods"`
}

type LogsData struct {
	ScheduledTime                  string `json:"scheduled"`
	TowardsPickupPointTime         string `json:"towards_pickup_point"`
	BeingAtThePickupPointTime      string `json:"being_at_the_pickup_point"`
	FinishLoadingTime              string `json:"finish_loading"`
	TowardsTheDestinationPointTime string `json:"towards_the_destination_point"`
	ReduceTheCargoTime             string `json:"reduce_the_cargo"`
	DoCompletedTime                string `json:"do_completed"`
}

type HeaderHtml struct {
	Title        string `json:"title"`
	Number       string `json:"number"`
	DownloadTime string `json:"download_time"`
}

type summaryGoodsWeight struct {
	TotalOrderGoods      float64 `json:"total_order_goods"`
	TotalScheduleGoods   float64 `json:"total_schedule_goods"`
	TotalReceivedGoods   float64 `json:"total_received_goods"`
	TotalDifferenceGoods float64 `json:"total_difference_goods"`
	TotalLoadingGoods    float64 `json:"total_loading_goods"`
}

type region struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Notes   string `json:"notes"`
}

type document struct {
	Type     string `json:"type"`
	Filename string `json:"filename"`
	Path     string `json:"path"`
	SignDate string `json:"sign_date"`
}

type goodsDetailPdf struct {
	Unit           string  `json:"unit"`
	GoodsName      string  `json:"goods_name"`
	OrderQty       float64 `json:"order_qty"`
	LoadingQty     float64 `json:"loading_qty"`
	ReceivedQty    float64 `json:"received_qty"`
	UnitWeightKg   float64 `json:"unit_weight_kg"`
	Difference     float64 `json:"difference"`
	TotalWeightKg  float64 `json:"total_weight_kg"`
	TotalWeightTon float64 `json:"total_weight_ton"`
}

type baseDataPdf struct {
	DoParent            string `json:"do_parent,omitempty"`
	DoNo                string `json:"do_no,omitempty"`
	Waybill             string `json:"waybill,omitempty"`
	LogkarNo            string `json:"logkar_no,omitempty"`
	RefNo               string `json:"ref_no,omitempty"`
	Status              int    `json:"status,omitempty"`
	StatusDoValue       string `json:"status_do_value,omitempty"`
	StartDate           string `json:"start_date,omitempty"`
	EndDate             string `json:"end_date,omitempty"`
	DriverName          string `json:"driver_name,omitempty"`
	TruckType           string `json:"truck_type,omitempty"`
	PoliceNo            string `json:"police_no,omitempty"`
	TruckCapacity       int    `json:"truck_capacity,omitempty"`
	CargoName           string `json:"cargo_name,omitempty"`
	TransporterName     string `json:"transporter_name,omitempty"`
	TruckRecommendation string `json:"truck_recommendation,omitempty"`
	OwnershipId         int    `json:"-"`
}

type DailySummaryReport struct {
	LogkarNo                 string  `json:"logkar_no"`
	RefNo                    *string `json:"ref_no"`
	DoNo                     string  `json:"do_no"`
	CargoName                string  `json:"cargo_name"`
	TransporterName          string  `json:"transporter_name"`
	DriverName               *string `json:"driver_name"`
	DriverPhone              *string `json:"driver_phone"`
	PoliceNo                 *string `json:"police_no"`
	DoStatus                 string  `json:"do_status"`
	UpdatedAt                string  `json:"updated_at"`
	CreatedAt                string  `json:"created_at"`
	StartDelivery            string  `json:"start_delivery"`
	EndDelivery              string  `json:"end_delivery"`
	OriginProvince           *string `json:"origin_province"`
	OriginCity               *string `json:"origin_city"`
	OriginDistrict           *string `json:"origin_district"`
	OriginVillage            *string `json:"origin_village"`
	OriginPostalCode         *string `json:"origin_postal_code"`
	OriginName               *string `json:"origin_name"`
	OriginAddress            *string `json:"origin_address"`
	OriginPhone              *string `json:"origin_phone"`
	OriginDetailAddress      *string `json:"origin_detail_address"`
	DestinationProvince      *string `json:"destination_province"`
	DestinationCity          *string `json:"destination_city"`
	DestinationDistrict      *string `json:"destination_district"`
	DestinationVillage       *string `json:"destination_village"`
	DestinationPostalCode    *string `json:"destination_postal_code"`
	DestinationName          *string `json:"destination_name"`
	DestinationPhone         *string `json:"destination_phone"`
	DestinationAddress       *string `json:"destination_address"`
	DestinationDetailAddress *string `json:"destination_detail_address"`
	TruckCapacity            string  `json:"truck_capacity"`
	TruckType                string  `json:"truck_type"`
	Qty                      string  `json:"qty"`
	TotalLoadingQty          *string `json:"total_loading_qty"`
	TotalDeliveringQty       *string `json:"total_delivering_qty"`
	OriginNetto              *string `json:"origin_netto"`
	OriginBruto              *string `json:"origin_bruto"`
	OriginTara               *string `json:"origin_tara"`
	DestinationNetto         *string `json:"destination_netto"`
	DestinationBruto         *string `json:"destination_bruto"`
	DestinationTara          *string `json:"destination_tara"`
	ActualFinisihLoading     *string `json:"actual_finisih_loading"`
	ActualStartDelivering    *string `json:"actual_start_delivering"`
	ActualEndDelivering      *string `json:"actual_end_delivering"`
}
