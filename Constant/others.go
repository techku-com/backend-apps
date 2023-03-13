package Constant

import "time"

const (
	TimeLocation                       = "Asia/Jakarta"
	DateTimeFormat                     = "2006-01-02 15:04:05"
	DateFormat                         = "2006-01-02"
	LogkarNo                           = "LOG-"
	Charset                            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	InitIntTypeDataValue               = 0
	InitStringTypeDataValue            = ""
	InvalidOrderStatusType             = -1
	OrderTypeContract                  = "contract"
	OrderTypeContractShuttle           = "contract-shuttle"
	OrderTypeNonContractInquiry        = "non-contract-inquiry"
	OldFormatPhoneNumber               = "62"
	NewFormatPhoneNumber               = "(+62) "
	DateFormatSlashSeparator           = "02/01/2006"
	OriginLocationType                 = 0
	DestinationLocationType            = 1
	TypeCompanyTransporter             = "transporter"
	TypeCompanyCargo                   = "cargo"
	KeyCacheDeliveryOrderListForDriver = "delivery::order::list::%d"
	WaybillPrefix                      = "IDL-"
	MobileDocumentPOD                  = "POD"
	MobileDocumentDN                   = "DN"
	SignDocumentDN                     = "DN-SIGN"
	SignDocumentPOD                    = "POD-SIGN"
	DriverSign                         = "DRIVER-SIGN"
	DeliveryOrderDetailByWaybill       = "delivery::order::detail::%s"
	DEFAULT_VALUE_1                    = 1
	FormatCabangName                   = "%s, %s"
	NotificationDetailCancelledOrder   = "notification::detail::cancel::%d"
	NewReportDirectoryName             = "New Report Recap"
	DateDDMMYYYYWithoutTime            = "02-01-2006"
	RedisSchedulingPerformanceKey      = "scheduling-performance::start_date::%s::company_id::%d"
	RedisSchedulingPerformanceKeyMap   = "%s_%d"
	TotalDaysInWeek                    = 7
	BenchmarkDoCompletedDays           = 40
	BenchmarkDoCompletedMonth          = 180
	MaximumPublicLink                  = 10
	PublicTrackingInitDataKey          = "public-tracking::waybill::%s::do_id::%s::link::%s"
	PublicTrackingInitTTL              = (time.Hour * 24) * 2
	LogkarPointValuePerTon             = 7.7
	LogkarPointValuePerKm              = 1
	LogkarPointEmailHeader             = "Pengiriman %s Selesai Dikirim, Ayo Berikan Penilaian!"
	JWTTokenFormat                     = "%d|%d|%d"
	LogkarPointOpsPage                 = "public/logkar-point?key=%s"
	RedisLogkarPointInfoKey            = "logkar-point::do_id::%d"
	DateFormatSlashSeparatorWithTime   = "02/01/2006 (15:04)"
	ComaSeperator                      = ","
)

const (
	DocumentTypeLNNo = iota
	DocumentTypePODNo
)
