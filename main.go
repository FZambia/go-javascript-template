package main

import (
	"flag"
	"log"

	_ "github.com/FZambia/go-javascript-template/statik"

	"github.com/FZambia/go-javascript-template/server"
)

var addr = flag.String("addr", ":8000", "http service address")
var web = flag.Bool("web", false, "serve web app")
var webPath = flag.String("web_path", "", "path to custom web app directory to serve [optional]")

func main() {
	flag.Parse()
	conf := &server.Config{
		Address: *addr,
		Web:     *web,
		WebPath: *webPath,
	}
	if err := server.Run(conf); err != nil {
		log.Fatal("Run server error: ", err)
	}
}
