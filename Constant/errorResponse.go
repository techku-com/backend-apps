package Constant

type ErrorType int

type ErrorStruct struct {
	Code  int
	Error string
}

const (
	InvalidJsonRequest ErrorType = iota
	InvalidGoogleAuth
	InvalidGoogleEmail
	InvalidRegister
	InvalidLogin
	InvalidOrder
)

func (p ErrorType) GetErrorStatus() ErrorStruct {
	errorCategory := map[ErrorType]ErrorStruct{
		InvalidJsonRequest: {
			Code:  400,
			Error: "Invalid JSON Request",
		},
		InvalidGoogleAuth: {
			Code:  403,
			Error: "Invalid Google Auth",
		},
		InvalidGoogleEmail: {
			Code:  403,
			Error: "Invalid Google Email",
		},
		InvalidRegister: {
			Code:  403,
			Error: "Invalid Registration",
		},
		InvalidLogin: {
			Code:  403,
			Error: "Invalid Login",
		},
		InvalidOrder: {
			Code:  403,
			Error: "Invalid Order",
		},
	}
	return errorCategory[p]
}
