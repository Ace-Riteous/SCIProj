package utils

import "time"

func Int64ToTime(i int64) time.Time {
	return time.Unix(i, 0)
}
