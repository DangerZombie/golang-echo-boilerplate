package util

import (
	"net/http"
	"time"
)

func UnixToFullDate(unix int64, layout string) string {
	tm := time.Unix(unix, 0).Format(layout)

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
