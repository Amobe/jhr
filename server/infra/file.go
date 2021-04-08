package infra

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/amobe/jhr/server/dto"
)

func OpenExcelFile(path string) (excelFile []dto.ExcelSheet, err error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		err = fmt.Errorf("open excel file: %w", err)
		return
	}
	for _, sheet := range f.GetSheetList() {
		rows, err := f.GetRows(sheet)
		if err != nil {
			continue
		}
		excelSheet := dto.NewExcelSheet(sheet, rows)
		excelFile = append(excelFile, excelSheet)
	}

	return
}

func WriteExcelFile(path string, excelFile []dto.ExcelSheet) error {
	f := excelize.NewFile()
	for _, sheet := range excelFile {
		f.NewSheet(sheet.Name)
		rows := [][]string(nil)
		rows = append(rows, sheet.HeaderRow)
		rows = append(rows, sheet.DataRows...)
		for i, row := range rows {
			writeExcelRow(f, sheet.Name, i+1, row)
		}
	}
	f.DeleteSheet("Sheet1")
	if err := f.SaveAs(path); err != nil {
		return fmt.Errorf("write excel file: %w", err)
	}
	return nil
}

func writeExcelRow(f *excelize.File, sheet string, index int, row []string) {
	if err := f.SetSheetRow(sheet, "A"+strconv.FormatInt(int64(index), 10), &row); err != nil {
		fmt.Printf("Err: write row to sheet, err: %v\n", err)
	}
}
