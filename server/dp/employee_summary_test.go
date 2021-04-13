package dp_test

import (
	"testing"

	"github.com/amobe/jhr/server/dp"
	"github.com/stretchr/testify/suite"
)

type EmployeeSummarySutie struct {
	suite.Suite
}

func TestEmployeeSummarySutie(t *testing.T) {
	suite.Run(t, new(EmployeeSummarySutie))
}

func (s *EmployeeSummarySutie) TestGetValidName() {
	give := "aaa:bc?"
	expected := "aaa-bc"

	got := dp.GetValidName(give)

	s.Equal(expected, got)
}

// func (s *EmployeeSummaryEntitySutie) TestNewEmployeeSummaryEntity() {
// 	got := xls.NewEmployeeSummaryEntity("department", "name")
// 	s.Equal("department", got.State.Department)
// 	s.Equal("name", got.State.Name)
// }

// func (s *EmployeeSummaryEntitySutie) TestAddRecord() {
// 	entity := xls.NewEmployeeSummaryEntity("department", "name")
// 	record := xls.NewEmployeeRecord("date", xls.NewTimeUnit(10), xls.NewTimeUnit(50))

// 	got := entity.AddRecord(record)

// 	s.Equal(1, len(got.Records))
// 	s.Equal("date", got.Records[0].Date)
// 	s.Equal(10, got.Records[0].OnDuty.ToMinutes())
// 	s.Equal(50, got.Records[0].OffDuty.ToMinutes())
// }

// func (s *EmployeeSummaryEntitySutie) TestCalculateAttendance() {
// 	entity := xls.NewEmployeeSummaryEntity("department", "name")
// 	record := xls.NewEmployeeRecord("date", xls.NewTimeUnit(10), xls.NewTimeUnit(50))

// 	entity.AddRecord(record)
// 	got := entity.CalculateAttendance()

// 	s.Equal(1, len(got.Records))
// 	s.Equal(xls.TooLow, got.Records[0].Status)
// 	s.Equal(40, got.Records[0].Duration.ToMinutes())
// }
