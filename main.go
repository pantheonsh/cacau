package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var debug = true
var addr = ":1337"
var htdocs = "htdocs"

func main() {
	cwd, _ := os.Getwd()
	servePath := filepath.Join(cwd, htdocs)
	_ = os.Mkdir(servePath, os.ModePerm)
	http.Handle("/", http.FileServer(http.Dir(servePath)))

	log.Fatal(http.ListenAndServe(addr, nil))
}
