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
	println("Cacau Mini - 2018 pantheonsh")

	cwd, _ := os.Getwd()
	servePath := filepath.Join(cwd, htdocs)
	_ = os.Mkdir(servePath, os.ModePerm)

	log.Println("Servindo arquivos a partir de", servePath)

	http.Handle("/", handler(http.FileServer(http.Dir(servePath))))

	log.Fatal(http.ListenAndServe(addr, nil))
}

func handler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-File-Server", "Cacau")

		h.ServeHTTP(w, r)
	}
}
