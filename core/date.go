package core

import "time"

type DateFormats string

var DD_MM_YYYY DateFormats = "02/01/2006"
var YYYY_MM_DD DateFormats = "2006-01-02"

func FormatDate(date time.Time, format DateFormats) string {
	newDate := date.Format(string(format))

	return newDate
}
