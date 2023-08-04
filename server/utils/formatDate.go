package utils

import "time"

func FormatDate(timestamp int64) string {
	seconds := timestamp / 1000
	nanoseconds := (timestamp % 1000) * 1000000

	t := time.Unix(seconds, nanoseconds)
	return t.Format("2006-01-02 15:04:05")
}
