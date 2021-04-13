package dp_test

import "github.com/amobe/jhr/server/dp"

func (s *EmployeeSummarySutie) TestIsValidDate() {
	give := "08/06 (四)"
	expected := true

	got := dp.IsValidDate(give)

	s.Equal(expected, got)
}
