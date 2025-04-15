package cron

import (
	"fmt"
	"strings"
	"time"
)

type ParseOption int

const (
	Second ParseOption = 1 << iota
	SecondOptional
	Minute
	Hour
	Dom
	Month
	Dow
	DowOptional
	Descriptor
)

var places = []ParseOption{
	Second,
	Minute,
	Hour,
	Dom,
	Month,
	Dow,
}

var defaults = []string{
	"0",
	"0",
	"0",
	"*",
	"*",
}

type Parser struct {
	options ParseOption
}

func NewParser(options ParseOption) Parser {
	optionals := 0

	if options&DowOptional > 0 {
		optionals++
	}

	if options&SecondOptional > 0 {
		options++
	}
	if optionals > 1 {
		panic("multiple optionals may not be configured")
	}
	return Parser{options}

}

func (p Parser) Parse(spec string) (Schedule, error) {
	if len(spec) == 0 {
		return nil, fmt.Errorf("empty spec string")
	}
	var loc = time.Local
	if strings.HasPrefix(spec, "TZ=") || strings.HasPrefix(spec, "CRON_TZ=") {
		var err error
		i := strings.Index(spec, " ")
		eq := strings.Index(spec, "=")
		if loc, err = time.LoadLocation(spec[eq+1 : i]); err != nil {
			return nil, fmt.Errorf("provided bad location %s: %v", spec[eq+1:i], err)
		}
		spec = strings.TrimSpace(spec[i:])
	}

	if strings.HasPrefix(spec, "@") {
		if p.options&Descriptor == 0 {
			return nil, fmt.Errorf("parser does not accept descriptors: %v", spec)
		}
		return
	}
	fields := strings.Fields(spec)
	var err error

}

func normalizeFields(fields []string, options ParseOption) ([]string, error) {
	optionals := 0
	if options&SecondOptional > 0 {
		options |= Second
		optionals++
	}

	if options&DowOptional > 0 {
		options |= Dow
		optionals++
	}
	if optionals > 1 {
		return nil, fmt.Errorf("multiple optionals may not be configured")
	}

	max := 0

	for _, place := range places {
		if options&place > 0 {
			max++
		}
	}

	min := max - optionals

	if count := len(fields); count < min || count > max {
		if min == max {
			return nil, fmt.Errorf("expected exactly %d fields, found %d: %s", min, count, fields)
		}
		return nil, fmt.Errorf("expected %d to %d fields, found %d: %s", min, max, count, fields)
	}

	if min < max && len(fields) == min {
		switch {
		case options&DowOptional > 0:
			fields = append(fields, defaults[5])
		case options&SecondOptional > 0:
			fields = append([]string{defaults[0]}, fields...)
		default:
			return nil, fmt.Errorf("unknown optional field")
		}
	}

	n := 0
	expandedFields := make([]string, len(places))
	copy(expandedFields, defaults)

	for i, place := range places {
		if options&place > 0 {
			expandedFields[i] = fields[n]
			n++
		}
	}

	return expandedFields, nil

}

var standardParser = NewParser(Minute | Hour | Dom | Month | Dow | Descriptor)

func ParseStandard(standardSpec string) (Schedule, error) {
	return standardParser.Parse(standardSpec)
}

func parseDescriptor(descriptor string, loc *time.Location) (Schedule, error) {
	// switch descriptor {
	// case "@yearly", "@annually":
	// 	return &
	// }
}
