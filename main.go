package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var debug = true
var host string
var port int
var htdocs = "htdocs"

func main() {
	flag.StringVar(&host, "h", "0.0.0.0", "Endereço onde o servidor deve executar.")
	flag.IntVar(&port, "p", 1337, "Porta em que o servidor deve executar.")
	flag.Parse()

	addr := host + ":" + strconv.Itoa(port)

	cwd, _ := os.Getwd()
	servePath := filepath.Join(cwd, htdocs)
	_ = os.Mkdir(servePath, os.ModePerm)

	println("Cacau Mini - 2018 pantheonsh")
	log.Println("Servindo arquivos a partir de", servePath)
	log.Println("O servidor está executando sobre o endereço", addr)

	http.Handle("/", handler(http.FileServer(http.Dir(servePath))))
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.RequestURI)
		w.Header().Set("X-File-Server", "Cacau")

		h.ServeHTTP(w, r)
	}
}
