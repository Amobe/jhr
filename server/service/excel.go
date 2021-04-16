package service

import (
	"strconv"

	"github.com/amobe/jhr/server/dp"
	"github.com/amobe/jhr/server/dto"
)

func SummaryExcel(excelFile []dto.ExcelSheet) ([]dto.ExcelSheet, error) {
	var out []dto.ExcelSheet
	var summaryList []dp.EmployeeSummary
	for _, employeeData := range excelFile {
		summary := dp.NewEmployeeSummary(employeeData)
		summaryList = append(summaryList, summary)
		out = append(out, applySummary(employeeData, summary))
	}
	summarySheet := abnormalSummary(summaryList)
	out = append([]dto.ExcelSheet{summarySheet}, out...)
	return out, nil
}

func applySummary(sheet dto.ExcelSheet, summary dp.EmployeeSummary) dto.ExcelSheet {
	sheet.Name = summary.Name
	for i := range sheet.DataRows {
		record := summary.Records[i]
		if record.AttendanceStatus == dp.Abnormal {
			continue
		}
		sheet.DataRows[i] = insertData(sheet.DataRows[i], 12, string(record.AttendanceStatus))
		sheet.DataRows[i] = insertData(sheet.DataRows[i], 13, record.Duration.ToString())
	}
	return sheet
}

func abnormalSummary(summaryList []dp.EmployeeSummary) dto.ExcelSheet {
	sheet := dto.NewExcelSheet("Summary", nil)
	sheet.HeaderRow = append(sheet.HeaderRow, "", "部門1", "名稱", "日期", "上班", "下班", "出勤狀態", "出勤時數")
	for _, s := range summaryList {
		abnormalRecords := s.ListNonEmptyRecord()
		for i, r := range abnormalRecords {
			sheet.DataRows = append(sheet.DataRows, recordToRow(i, s.Department, s.EmployeeName, r))
		}
	}
	return sheet
}

func insertData(row []string, index int, data string) []string {
	if len(row) < index+1 {
		span := make([]string, index+1-len(row))
		row = append(row, span...)
	}
	row[index] = data
	return row
}

func recordToRow(index int, departmentName string, employeeName string, r dp.EmployeeRecord) []string {
	res := append([]string(nil), strconv.Itoa(index), departmentName,
		employeeName, r.Date, getDisplayTime(r.OnDuty), getDisplayTime(r.OffDuty),
		getAttendanceStatus(r.AttendanceStatus), getDisplayTime(r.Duration))
	return res
}

func getAttendanceStatus(s dp.AttendanceStatus) string {
	switch s {
	case dp.Normal:
		return "正常"
	case dp.Abnormal:
		return "異常"
	case dp.TooLow:
		return "小於九小時"
	case dp.TooHigh:
		return "大於九小時"
	default:
		return ""
	}
}

func getDisplayTime(t dp.TimeUnit) string {
	if t.ToMinutes() == 0 {
		return ""
	}
	return t.ToString()
}
