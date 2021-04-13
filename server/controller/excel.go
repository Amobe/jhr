package controller

import (
	"github.com/amobe/jhr/server/dp"
	"github.com/amobe/jhr/server/dto"
)

func Handle(excelFile []dto.ExcelSheet) ([]dto.ExcelSheet, error) {
	var out []dto.ExcelSheet
	for _, employeeData := range excelFile {
		summary := dp.NewEmployeeSummary(employeeData)
		out = append(out, applySummary(employeeData, summary))
	}
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

func insertData(row []string, index int, data string) []string {
	if len(row) < index+1 {
		span := make([]string, index+1-len(row))
		row = append(row, span...)
	}
	row[index] = data
	return row
}
