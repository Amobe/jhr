package infra

import (
	"fmt"
	"io"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/amobe/jhr/server/dto"
	"github.com/amobe/jhr/server/infra/excel"
)

func OpenExcelStream(reader io.Reader) (excelFile []dto.ExcelSheet, err error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		err = fmt.Errorf("open excel stream: %w", err)
		return
	}
	return excel.ReadExcelFile(f), nil
}

func WriteExcelStream(writer io.Writer, excelFile []dto.ExcelSheet) error {
	dst := excel.WriteExcelFile(excelize.NewFile(), excelFile)
	if err := dst.Write(writer); err != nil {
		return fmt.Errorf("write excel stream: %w", err)
	}
	return nil
}
