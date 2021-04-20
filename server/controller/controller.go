package controller

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/amobe/jhr/server/dto"
	"github.com/amobe/jhr/server/infra"
	"github.com/amobe/jhr/server/service"
)

const MaximunDataSize = 10 * 1000 * 1000 // 10MB

var storage *ExcelData

type ExcelData struct {
	ContentType string
	Data        []dto.ExcelSheet
}

type HandlerWrapper struct {
	name string
	fn   func(w http.ResponseWriter, r *http.Request) error
}

func (wrapper HandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: %s\n", wrapper.name)
	err := wrapper.fn(w, r)
	if err == nil {
		return
	}
	fmt.Printf("Err: %+v\n", err)
	w.Write([]byte(err.Error()))
}

func ExcelHandler() HandlerWrapper {
	return HandlerWrapper{"excel", excelHandler}
}

func excelHandler(w http.ResponseWriter, r *http.Request) error {
	r.ParseMultipartForm(MaximunDataSize)
	form := r.MultipartForm
	if form == nil {
		return fmt.Errorf("multipart form is empty")
	}

	for k, _ := range form.File {
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			return fmt.Errorf("open file from form: %w", err)
		}
		defer file.Close()
		fmt.Printf("the uploaded file: name[%s], size[%d], header[%#v]\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		out, err := analyzeExcel(file, w)
		if err != nil {
			return fmt.Errorf("analyze excel: %w", err)
		}
		storage = &ExcelData{
			ContentType: fileHeader.Header.Get("Content-Type"),
			Data:        out,
		}
		downloadLastResultFromBrowser(r.Host)
	}
	return nil
}

func LastHandler() HandlerWrapper {
	return HandlerWrapper{"last", lastHandler}
}

func lastHandler(w http.ResponseWriter, r *http.Request) error {
	if storage == nil {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	w.Header().Set("Content-Type", storage.ContentType)
	if err := infra.WriteExcelStream(w, storage.Data); err != nil {
		return fmt.Errorf("save excel stream: %w", err)
	}
	return nil
}

func analyzeExcel(inStream io.Reader, outStream io.Writer) ([]dto.ExcelSheet, error) {
	excel, err := infra.OpenExcelStream(inStream)
	if err != nil {
		return nil, fmt.Errorf("open excel stream: %w", err)
	}
	out, err := service.SummaryExcel(excel)
	if err != nil {
		return nil, fmt.Errorf("summary excel file: %w", err)
	}
	return out, nil
}

func downloadLastResultFromBrowser(host string) {
	openBrowser(fmt.Sprintf("http://%s/last", host))
}

func openBrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
