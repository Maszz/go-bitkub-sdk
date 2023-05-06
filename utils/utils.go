package utils

import "time"

func CurrentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
