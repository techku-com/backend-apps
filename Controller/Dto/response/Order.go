package response

type AllOrders struct {
	CreatedBy     string `json:"created_by"`
	Description   string `json:"description"`
	Location      string `json:"location"`
	ContactPerson string `json:"contact_person"`
	HardwareType  string `json:"hardware_type"`
	ServiceType   string `json:"service_type"`
	ScheduleDate  string `json:"schedule_date"`
}
