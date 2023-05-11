package utils

import (
	"time"

	"github.com/Maszz/go-bitkub-sdk/types"
)

func CurrentTimestamp() types.Timestamp {
	return types.Timestamp(FormatTimestamp(time.Now()))
}

func RawCurrentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / (int64(time.Second)) // convert to Seconds
}

func IndexOf[T string](arr []T, item T) int {
	for i, v := range arr {
		if v == item {
			return i
		}
	}

	return -1
}
