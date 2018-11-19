package main

import (
	"fmt"
	"net"
)

var debug = false
var addr = ":1337"

func main() {
	ln, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println("Erro ao iniciar o servidor.")
		fmt.Println(err)
		return
	}

	if debug {
		fmt.Println("Servidor HTTP está rodando sobre o endereço " + addr)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar uma conexão.")
			fmt.Println(err)
		} else {
			handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) {
	conn.
}
