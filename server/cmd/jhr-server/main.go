package main

import (
	"net/http"

	"github.com/amobe/jhr/server/controller"
	"github.com/gobuffalo/packr/v2"
	"github.com/webview/webview"
)

const ClientBuildFolder = "../../../client/build"
const ListeningPort = ":8000"
const ServingUrl = "http://localhost:8000"

func main() {
	startWebServer()
	startWebView()
}

func startWebServer() {
	box := packr.New("My Box", ClientBuildFolder)
	http.Handle("/", http.FileServer(box))
	http.Handle("/excel", controller.ExcelHandler())
	http.Handle("/last", controller.LastHandler())
	go http.ListenAndServe(ListeningPort, nil)
}

func startWebView() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("JHR")
	w.SetSize(600, 400, webview.HintNone)
	w.Navigate(ServingUrl)
	w.Run()
}
