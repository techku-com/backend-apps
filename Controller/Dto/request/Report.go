package request

type PdfReport struct {
	UserAuth
	DoNo    string `json:"do_no"`
	RefNo   string `json:"ref_no"`
	WayBill string `json:"waybill"`
	DoId    int    `json:"do_id"`
}
