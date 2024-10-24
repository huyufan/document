package common

import (
	"exec/go/exec/limit"
	"time"
)

func GetContextFromState(now time.Time, rate limit.Rate, expiration time.Time, count int64) limit.Context {
	limitr := rate.Limit
	remaining := int64(0)
	reached := true
	if count <= limitr {
		remaining = limitr - count
		reached = false
	}
	reset := expiration.Unix()
	return limit.Context{
		Limit:     limitr,
		Remaining: remaining,
		Reset:     reset,
		Reached:   reached,
	}
}
