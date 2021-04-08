package infra

import (
	"fmt"
	"io"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func OpenExcelFromStream(reader io.Reader) (*excelize.File, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, fmt.Errorf("open excel stream: %w", err)
	}
	return f, nil
}
