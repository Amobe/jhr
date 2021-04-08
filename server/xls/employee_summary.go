package xls

type EmployeeRecord struct {
	Date     string
	OnDuty   TimeUnit
	OffDuty  TimeUnit
	Status   AttendanceStatus
	Duration TimeUnit
}

func NewEmployeeRecord(date string, onDuty, offDuty TimeUnit) EmployeeRecord {
	return EmployeeRecord{
		Date:    date,
		OnDuty:  onDuty,
		OffDuty: offDuty,
	}
}

type EmployeeSummary struct {
	Department string
	Name       string
	Records    []EmployeeRecord
}

func NewEmployeeSummary(department string, name string) EmployeeSummary {
	return EmployeeSummary{
		Department: department,
		Name:       name,
	}
}

type EmployeeSummaryEntity struct {
	State EmployeeSummary
}

func NewEmployeeSummaryEntity(department string, name string) *EmployeeSummaryEntity {
	return &EmployeeSummaryEntity{
		State: NewEmployeeSummary(department, name),
	}
}

func (e *EmployeeSummaryEntity) AddRecord(record EmployeeRecord) EmployeeSummary {
	e.State.Records = append(e.State.Records, record)
	return e.State
}

func (e *EmployeeSummaryEntity) CalculateAttendance() EmployeeSummary {
	for i, r := range e.State.Records {
		status, diff := getAttendanceStatus(r.OnDuty, r.OffDuty)
		e.State.Records[i].Status = status
		e.State.Records[i].Duration = diff
	}
	return e.State
}
