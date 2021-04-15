package main

import (
	"fmt"
	"net/http"

	"github.com/amobe/jhr/server/controller"
	"github.com/amobe/jhr/server/infra"
	"github.com/gobuffalo/packr/v2"
	"github.com/webview/webview"
)

const (
	ClientBuildFolder string = "../../../client/build"
)

func main() {
	box := packr.New("My Box", ClientBuildFolder)
	http.Handle("/", http.FileServer(box))
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/excel", excelHandler)
	go http.ListenAndServe(":8000", nil)

	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("JHR")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:8000")
	w.Run()

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := []byte("{\"text\":world}")
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func excelHandler(w http.ResponseWriter, r *http.Request) {
	const _10MB = 10 * 1000 * 1000
	r.ParseMultipartForm(_10MB)
	form := r.MultipartForm
	if form == nil {
		w.Write([]byte("no data"))
		return
	}

	for k, _ := range form.File {
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		defer file.Close()
		fmt.Printf("the uploaded file: name[%s], size[%d], header[%#v]\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)
		excel, err := infra.OpenExcelStream(file)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		out, err := controller.Handle(excel)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		if err := infra.WriteExcelStream(w, out); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}
