package main

import (
	"net/http"

	"github.com/gobuffalo/packr/v2"
	"github.com/webview/webview"
)

const (
	ClientBuildFolder string = "../client/build"
)

func main() {
	box := packr.New("My Box", ClientBuildFolder)
	http.Handle("/", http.FileServer(box))
	http.HandleFunc("/hello", helloHandler)
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
