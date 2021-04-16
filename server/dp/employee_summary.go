package dp

import (
	"strings"

	"github.com/amobe/jhr/server/dto"
)

type EmployeeSummary struct {
	Name            string
	Department      string
	EmployeeName    string
	Records         []EmployeeRecord
	Status          EmployeeSummaryStatus
	AbnormalRecords []EmployeeRecord
}

func NewEmployeeSummary(data dto.ExcelSheet) (summary EmployeeSummary) {
	if len(data.DataRows) < 1 {
		summary.Status = EmployeeSummaryStatusEmpty
		return
	}
	summary.Department = getDataFromRow(data.DataRows[0], 0)
	summary.EmployeeName = getDataFromRow(data.DataRows[0], 1)
	summary.Name = GetValidName(summary.EmployeeName)
	for _, row := range data.DataRows {
		date := getDataFromRow(row, 2)
		onDutyTime := getDataFromRow(row, 9)
		offDutyTime := getDataFromRow(row, 10)
		employeeRecord := NewEmployeeRecord(date, onDutyTime, offDutyTime)
		summary.Records = append(summary.Records, employeeRecord)
	}
	return summary.CalculateAttendance()
}

func (s EmployeeSummary) CalculateAttendance() EmployeeSummary {
	for i, r := range s.Records {
		status, diff := GetAttendanceStatus(r.OnDuty, r.OffDuty)
		s.Records[i].AttendanceStatus = status
		s.Records[i].Duration = diff
	}
	return s
}

func (s EmployeeSummary) ListNonEmptyRecord() (nonEmptyRecords []EmployeeRecord) {
	for _, r := range s.Records {
		if r.Status == EmployeeRecordStatusEmpty {
			continue
		}
		nonEmptyRecords = append(nonEmptyRecords, r)
	}
	return
}

func getDataFromRow(row []string, index int) string {
	if len(row) > index {
		return row[index]
	}
	return ""
}

func GetValidName(ori string) string {
	name := strings.ReplaceAll(ori, ":", "-")
	return strings.ReplaceAll(name, "?", "")
}

func GetAttendanceStatus(start, end TimeUnit) (status AttendanceStatus, diff TimeUnit) {
	if !start.IsValid() || !end.IsValid() {
		return Abnormal, NewInvalidTimeUnit()
	}
	diff = diffTimeUnit(start, end)
	return getAttendanceStatus(diff.ToMinutes()), diff
}

func getAttendanceStatus(minutes int) AttendanceStatus {
	switch {
	case minutes < 9*60:
		return TooLow
	case minutes > 9.5*60:
		return TooHigh
	default:
		return Normal
	}
}

func diffTimeUnit(start, end TimeUnit) TimeUnit {
	diffMinutes := end.ToMinutes() - start.ToMinutes()
	return NewTimeUnit(diffMinutes)
}
