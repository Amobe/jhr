package main

import (
	"fmt"
	"log"

	"github.com/amobe/jhr/server/controller"
	"github.com/amobe/jhr/server/infra"
)

func main() {
	const (
		inFilePath  = "./MonRep200725_00000_00300.xlsx"
		outFilePath = "./MonRep200725_00000_00300-summary.xlsx"
	)
	if err := analyzeExcel(inFilePath, outFilePath); err != nil {
		log.Fatalln(err)
	}
}

func analyzeExcel(inFilePath, outFilePath string) error {
	excel, err := infra.OpenExcelFile(inFilePath)
	if err != nil {
		return fmt.Errorf("open excel file: %w", err)
	}
	out, err := controller.Handle(excel)
	if err != nil {
		return fmt.Errorf("handle excel file: %w", err)
	}
	if err := infra.WriteExcelFile(outFilePath, out); err != nil {
		return fmt.Errorf("save excel file: %w", err)
	}
	return nil
}
