package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/FZambia/go-javascript-template/statik" // TODO: Replace with your import path

	"github.com/rakyll/statik/fs"
)

var addr = flag.String("addr", ":8000", "http service address")
var web = flag.Bool("web", false, "serve web app")
var webPath = flag.String("web_path", "", "path to custom web app directory to serve [optional]")

func main() {
	flag.Parse()
	if *web {
		var appFS http.FileSystem
		if *webPath != "" {
			appFS = http.Dir("app")
		} else {
			appFS, _ = fs.New()
		}
		http.Handle("/", http.FileServer(appFS))
	}
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
