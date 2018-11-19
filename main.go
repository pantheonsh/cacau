package main

import (
	"fmt"
	"net/http"
)

var debug = true
var addr = ":1337"

func main() {
	http.HandleFunc("/", handleConnection)
	http.ListenAndServe(addr, nil)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	if debug {
		fmt.Print(r.RemoteAddr, r.Method)
	}

	w.Header().Add("X-Server", "Cacau")
	w.WriteHeader(200)

	fmt.Fprintf(w, "Olá, mundo!")
}
