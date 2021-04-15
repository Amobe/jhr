package excel

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/amobe/jhr/server/dto"
)

func WriteExcelFile(src *excelize.File, excelData []dto.ExcelSheet) *excelize.File {
	for _, sheet := range excelData {
		src.NewSheet(sheet.Name)
		rows := [][]string(nil)
		rows = append(rows, sheet.HeaderRow)
		rows = append(rows, sheet.DataRows...)
		for i, row := range rows {
			writeExcelRow(src, sheet.Name, i+1, row)
		}
	}
	src.DeleteSheet("Sheet1")
	return src
}

func writeExcelRow(f *excelize.File, sheet string, index int, row []string) {
	if err := f.SetSheetRow(sheet, "A"+strconv.FormatInt(int64(index), 10), &row); err != nil {
		fmt.Printf("Err: write row to sheet, err: %v\n", err)
	}
}
