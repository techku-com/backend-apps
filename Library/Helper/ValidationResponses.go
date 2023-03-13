package Helper

import (
	"net/textproto"
	"techku/Constant"
)

func ValidationFileUpload(params textproto.MIMEHeader) bool {
	checkHeader := params.Get("Content-Type")
	if checkHeader == Constant.ImageJpeg || checkHeader == Constant.ImagePng {
		return true
	}
	return false
}

func ValidationFileSizeUpload(params int64) bool {
	var megabytes int64

	megabytes = params / (1024 * 1024)
	if megabytes <= 2 {
		return true
	}

	return false
}

func CheckConstantType(value string) bool {
	if value == Constant.ImageNpwp {
		return true
	}

	if value == Constant.ImageKtp {
		return true
	}

	if value == Constant.ImageCompanyLogo {
		return true
	}
	return false
}
