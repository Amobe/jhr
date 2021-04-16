package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/amobe/jhr/server/infra"
	"github.com/amobe/jhr/server/service"
	"github.com/gobuffalo/packr/v2"
	"github.com/webview/webview"
)

const ClientBuildFolder = "../../../client/build"
const ListeningPort = ":8000"
const ServingUrl = "http://localhost:8000"
const MaximunDataSize = 10 * 1000 * 1000 // 10MB

func main() {
	startWebServer()
	startWebView()
}

func startWebServer() {
	box := packr.New("My Box", ClientBuildFolder)
	http.Handle("/", http.FileServer(box))
	http.Handle("/excel", handlerWrapper(excelHandler))
	go http.ListenAndServe(ListeningPort, nil)
}

func startWebView() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("JHR")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate(ServingUrl)
	w.Run()
}

type handlerWrapper func(w http.ResponseWriter, r *http.Request) error

func (fn handlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err == nil {
		return
	}
	w.Write([]byte(err.Error()))
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

		if err := analyzeExcel(file, w); err != nil {
			return fmt.Errorf("analyze excel: %w", err)
		}
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	}

	return nil
}

func analyzeExcel(inStream io.Reader, outStream io.Writer) error {
	excel, err := infra.OpenExcelStream(inStream)
	if err != nil {
		return fmt.Errorf("open excel stream: %w", err)
	}
	out, err := service.SummaryExcel(excel)
	if err != nil {
		return fmt.Errorf("summary excel file: %w", err)
	}
	if err := infra.WriteExcelStream(outStream, out); err != nil {
		return fmt.Errorf("save excel stream: %w", err)
	}
	return nil
}
