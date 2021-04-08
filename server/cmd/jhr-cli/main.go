package main

import (
	"log"

	"github.com/amobe/jhr/server/controller"
	"github.com/amobe/jhr/server/infra"
)

func main() {
	const (
		inputFilePath  = "./MonRep200725_00000_00300.xlsx"
		outputFilePath = "./MonRep200725_00000_00300-summary.xlsx"
	)
	excel, err := infra.OpenExcelFile(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	out, err := controller.Handle(excel)
	if err != nil {
		log.Fatal(err)
	}
	if err := infra.WriteExcelFile(outputFilePath, out); err != nil {
		log.Fatal(err)
	}
}
