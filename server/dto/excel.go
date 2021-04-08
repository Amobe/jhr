package dto

type ExcelSheet struct {
	Name      string
	HeaderRow []string
	DataRows  [][]string
}

func NewExcelSheet(name string, data [][]string) (sheet ExcelSheet) {
	sheet.Name = name
	if len(data) < 1 {
		return
	}
	sheet.HeaderRow = cloneRow(data[0])
	for _, row := range data[1:] {
		sheet.DataRows = append(sheet.DataRows, cloneRow(row))
	}
	return
}

func cloneRow(in []string) (out []string) {
	out = make([]string, len(in))
	copy(out, in)
	return
}
