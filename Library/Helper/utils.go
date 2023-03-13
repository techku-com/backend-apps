package Helper

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"techku/Constant"
	"time"
)

func ConvertToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}

func ReplaceSpecialChar(str string) (replacedStr string) {
	replacer := strings.NewReplacer("_", "\\_", "%", "\\%")
	replacedStr = replacer.Replace(str)
	return
}

func GenerateUUID() (id string) {
	id = strings.Replace(uuid.NewV4().String(), "-", "", -1)
	return
}

func ConvertToFloat(value string) float64 {
	FloatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return FloatValue
}

func RandomString(length int) (result string) {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = Constant.Charset[seededRand.Intn(len(Constant.Charset))]
	}

	result = string(bytes)
	return
}

func ConvertSHA256(param string) (result string) {
	convert := sha256.Sum256([]byte(param))
	result = fmt.Sprintf("%x", convert)

	return
}

func RequestCodeFormat(clientId, apiKey string, additional ...interface{}) string {
	requestCode := fmt.Sprintf("#%s#%s#", clientId, apiKey)
	if len(additional) > 0 {
		for _, value := range additional {
			requestCode += fmt.Sprintf("%v#", value)
		}
	}

	return ConvertSHA256(requestCode)
}

func IsKeyExist(decoded map[string]interface{}, key string) bool {
	val, ok := decoded[key]
	return ok && val != nil
}

func RandomNumber(from, to int) int {
	return rand.Intn(to-from+1) + from
}

func MatchSpecialChar(pattern string, values ...string) (bool, error) {
	for _, value := range values {
		matched, err := regexp.MatchString(pattern, value)
		if matched || err != nil {
			return matched, err
		}
	}
	return false, nil
}

func EncodeHtml(data interface{}) (res []byte, err error) {
	buffer := &bytes.Buffer{}
	encode := json.NewEncoder(buffer)
	encode.SetEscapeHTML(false)
	err = encode.Encode(data)
	return buffer.Bytes(), err
}

func FormatPhoneNumberINA(phoneNumber string) string {
	return strings.Replace(phoneNumber, Constant.OldFormatPhoneNumber, Constant.NewFormatPhoneNumber, 1)
}

func FormatPlateNumberWithSpace(plateNumber string) (formattedPlate string) {
	reg := regexp.MustCompile("(\\d+)")
	number := reg.FindAllString(plateNumber, 1)
	formatNumber := fmt.Sprintf(" %s ", number[0])
	formattedPlate = reg.ReplaceAllString(plateNumber, formatNumber)
	return
}

func GenerateLogkarWaybill() string {
	randNum := time.Now().UTC().Unix() + int64(RandomNumber(10000, 99999))
	waybillNo := fmt.Sprintf("%s%d%d", Constant.WaybillPrefix, randNum, RandomNumber(10000, 99999))
	return waybillNo
}

func GenerateUniqueCodeForImages() int {
	return RandomNumber(1, 999)
}

func CustomErrMsg(dataRequest interface{}, errMsg string) error {
	dataBytes, _ := json.Marshal(dataRequest)
	errInfo := map[string]interface{}{
		"payload": string(dataBytes),
		"error":   errMsg,
	}
	errBytes, _ := json.Marshal(errInfo)
	return errors.New(string(errBytes))
}
