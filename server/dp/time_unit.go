package dp

import (
	"fmt"
	"regexp"
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
	if minutes < 0 {
		minutes *= -1
	}
	return TimeUnit{
		Hours:   minutes / 60,
		Minutes: minutes % 60,
		Status:  Valid,
	}
}

func NewTimeUnitFromString(time string) (t TimeUnit) {
	if !isValidTime(time) {
		t.Status = Invalid
		return
	}
	units := strings.Split(time, ":")
	hours, err := strconv.ParseInt(units[0], 10, 64)
	if err != nil {
		t.Status = Invalid
		return
	}
	minutes, err := strconv.ParseInt(units[1], 10, 64)
	if err != nil {
		t.Status = Invalid
		return
	}
	return TimeUnit{
		Hours:   int(hours),
		Minutes: int(minutes),
		Status:  Valid,
	}
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

func isValidTime(time string) bool {
	reg := regexp.MustCompile("[0-9]+:[0-9]+")
	return reg.Match([]byte(time))
}
