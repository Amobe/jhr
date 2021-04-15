package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/amobe/jhr/server/dto"
)

func ReadExcelFile(f *excelize.File) (excelFile []dto.ExcelSheet) {
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
