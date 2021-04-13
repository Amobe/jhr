package dp

import (
	"regexp"
)

type EmployeeRecord struct {
	Date             string
	OnDuty           TimeUnit
	OffDuty          TimeUnit
	AttendanceStatus AttendanceStatus
	Duration         TimeUnit
	Status           EmployeeRecordStatus
}

func NewEmployeeRecord(date string, onDuty, offDuty string) (r EmployeeRecord) {
	if !IsValidDate(date) {
		r.Status = EmployeeRecordStatusEmpty
		return
	}
	return EmployeeRecord{
		Date:    date,
		OnDuty:  NewTimeUnitFromString(onDuty),
		OffDuty: NewTimeUnitFromString(offDuty),
	}
}

func IsValidDate(date string) bool {
	reg := regexp.MustCompile("[0-9]+/[0-9]+")
	return reg.Match([]byte(date))
}
