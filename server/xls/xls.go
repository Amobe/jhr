package xls

import (
	"math"
	"regexp"
	"strconv"
)

func isValidDate(date string) bool {
	reg := regexp.MustCompile("[0-9]+/[0-9]+")
	return reg.Match([]byte(date))
}

func isValidTime(time string) bool {
	reg := regexp.MustCompile("[0-9]+:[0-9]+")
	return reg.Match([]byte(time))
}

func isNanOrTime(value string) bool {
	if v, err := strconv.ParseFloat(value, 64); err == nil {
		return math.IsNaN(v)
	}
	return isValidTime(value)
}

func getMinutes(time string) (int, error) {
	t, err := NewTimeUnitFromString(time)
	if err != nil {
		return 0, err
	}
	return t.ToMinutes(), nil
}

func formatTime(minutes int) string {
	t := NewTimeUnit(minutes)
	return t.ToString()
}

func diffTimeUnit(start, end TimeUnit) TimeUnit {
	diffMinutes := end.ToMinutes() - start.ToMinutes()
	return NewTimeUnit(diffMinutes)
}

func getAttendanceStatus(start, end TimeUnit) (status AttendanceStatus, diff TimeUnit) {
	if !start.IsValid() || !end.IsValid() {
		return Abnormal, NewInvalidTimeUnit()
	}
	diff = diffTimeUnit(start, end)
	return GetAttendanceStatus(diff.ToMinutes()), diff
}
