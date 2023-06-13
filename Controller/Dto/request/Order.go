package request

type AddNewOrder struct {
	UserId        int    `json:"user_id" field:"required" min:"1"`
	HardwareType  string `json:"hardware_type" field:"required"`
	ServiceType   string `json:"service_type" field:"required"`
	Description   string `json:"description" field:"required"`
	Location      string `json:"location" field:"required"`
	ContactPerson string `json:"contact_person" field:"required"`
	ScheduleDate  string `json:"schedule_date" field:"required"`
}
