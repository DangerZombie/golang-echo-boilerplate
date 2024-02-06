package util

import (
	"net/http"
	"strconv"
	"time"
)

func UnixMilliToFullDate(unix int64, layout string) string {
	tm := time.UnixMilli(unix).Format(layout)

	return tm
}

func ReplaceEmptyString(str string, defaultValue string) string {
	if len(str) == 0 {
		return defaultValue
	}

	return str
}

func WrappingStatusCode(msgCode int) int {
	switch msgCode {
	case 34004:
		return http.StatusUnauthorized
	default:
		return http.StatusOK
	}
}

func StringToInt(input string) (output int) {
	if input == "" {
		output = 0
	}

	output, err := strconv.Atoi(input)
	if err != nil {
		output = 0
	}

	return
}
