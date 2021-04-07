package xls

import (
	"fmt"
	"strconv"
	"strings"
)

type TimeUnit struct {
	Hours   int
	Minutes int
	Status  TimeUnitStatus
}

type TimeUnitStatus string

const (
	Valid   TimeUnitStatus = "valid"
	Invalid TimeUnitStatus = "invalid"
)

func NewTimeUnit(minutes int) TimeUnit {
	return TimeUnit{
		Hours:   minutes / 60,
		Minutes: minutes % 60,
		Status:  Valid,
	}
}

func NewTimeUnitFromString(time string) (t TimeUnit, err error) {
	units := strings.Split(time, ":")
	if len(units) != 2 {
		err = fmt.Errorf("time is not a valid string, len: %d != 2", len(units))
		return
	}
	hours, err := strconv.ParseInt(units[0], 10, 64)
	if err != nil {
		err = fmt.Errorf("could not parse hours from string: %s not a value", units[0])
		return
	}
	minutes, err := strconv.ParseInt(units[1], 10, 64)
	if err != nil {
		err = fmt.Errorf("could not parse minutes from string: %s not a value", units[1])
		return
	}
	return TimeUnit{
		Hours:   int(hours),
		Minutes: int(minutes),
		Status:  Valid,
	}, nil
}

func NewInvalidTimeUnit() TimeUnit {
	return TimeUnit{
		Status: Invalid,
	}
}

func (t TimeUnit) ToMinutes() int {
	return t.Hours*60 + t.Minutes
}

func (t TimeUnit) ToString() string {
	return fmt.Sprintf("%d:%d", t.Hours, t.Minutes)
}

func (t TimeUnit) IsValid() bool {
	return t.Status == Valid
}
