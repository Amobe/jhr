package dp

import (
	"strings"

	"github.com/amobe/jhr/server/dto"
)

type EmployeeSummary struct {
	Name         string
	Department   string
	EmployeeName string
	Records      []EmployeeRecord
	Status       EmployeeSummaryStatus
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
