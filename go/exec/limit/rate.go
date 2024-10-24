package limit

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Rate struct {
	Formatted string
	Period    time.Duration
	Limit     int64
}

func NewRateFromFormatted(formatted string) (Rate, error) {
	rate := Rate{}
	value := strings.Split(formatted, "-")
	if len(value) != 2 {
		str := fmt.Sprintf("incorrect format '%s'", formatted)
		return rate, errors.New(str)
	}

	periods := map[string]time.Duration{
		"S": time.Second,
		"M": time.Minute,
		"H": time.Hour,
		"D": time.Hour * 24,
	}

	limit, period := value[0], strings.ToUpper(value[1])

	p, ok := periods[period]

	if !ok {
		str := fmt.Sprintf("incorrect period '%s'", period)
		return rate, errors.New(str)
	}
	l, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		str := fmt.Sprintf("incorrect limit '%s'", limit)
		return rate, errors.New(str)
	}

	rate = Rate{
		Formatted: formatted,
		Limit:     l,
		Period:    p,
	}

	return rate, nil

}
