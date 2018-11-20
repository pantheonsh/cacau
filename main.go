package main

import (
	"fmt"
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

	println("Servindo arquivos a partir de " + servePath)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	if debug {
		fmt.Print(r.RemoteAddr, r.Method)
	}

	w.Header().Add("X-File-Server", "Cacau (Golang)")
	w.WriteHeader(200)

	fmt.Fprintf(w, "Olá, mundo!")
}
