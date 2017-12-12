package util

import "time"

var Started time.Time

func init() {
	Started = time.Now()
}

func ElapsedTime() time.Duration {
	return time.Since(Started)
}
