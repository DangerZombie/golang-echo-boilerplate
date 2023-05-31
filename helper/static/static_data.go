package static

import (
	"time"
)

var (
	LayoutDefault            = "2006-01-02 15:04:05"
	LayoutDateOnly           = "2006-01-02"
	LayoutNotification       = "02 Jan 2006"
	LayoutTimeOnly           = "15:04:05"
	LayoutHourMinute         = "15:04"
	LayoutUTC                = "2006-01-02T15:04:05Z"
	LayoutUTCMiliSec         = "2006-01-02T15:04:05.000Z"
	LayoutDateGMT            = "2006-01-02 15:04:05 -0700"
	LayoutWithoutDaySimplify = "02 Jan 2006 15:04"
	LayoutInvoice            = "20060102"
	Loc, _                   = time.LoadLocation("Asia/Jakarta")
)
