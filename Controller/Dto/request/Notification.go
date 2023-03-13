package request

type MobileNotificationRequest struct {
	Status      int    `json:"status"`
	DriverId    int    `json:"driver_id"`
	DoId        int    `json:"do_id"`
	Description string `json:"description,omitempty"`
	Date        string `json:"date,omitempty"`
	DoNo        string `json:"do_no,omitempty"`
	WayBill     string `json:"way_bill,omitempty"`
	IsCombineDo string `json:"is_combine_do,omitempty"`
	LogkarPoint int    `json:"logkar_point"`
}
