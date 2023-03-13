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
	}
	return errorCategory[p]
}
