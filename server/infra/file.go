package infra

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/amobe/jhr/server/dto"
	"github.com/amobe/jhr/server/infra/excel"
)

func OpenExcelFile(path string) (excelFile []dto.ExcelSheet, err error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		err = fmt.Errorf("open excel file: %w", err)
		return
	}
	return excel.ReadExcelFile(f), nil
}

func WriteExcelFile(path string, excelFile []dto.ExcelSheet) error {
	dst := excel.WriteExcelFile(excelize.NewFile(), excelFile)
	if err := dst.SaveAs(path); err != nil {
		return fmt.Errorf("write excel file: %w", err)
	}
	return nil
}
