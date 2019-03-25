package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	defer ln.Close()
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	request(conn)
	response(conn)
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			lin1 := strings.Fields(ln)
			m := lin1[0]
			u := lin1[1]
			fmt.Println("=====the Method is ", m, "======")
			fmt.Println("=====the URL is ", u, "======")

		}
		if ln == "" {
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 ok\n")
	fmt.Fprint(conn, "Content-Length: ", len(body), "\n")
	fmt.Fprint(conn, "Content-Type: ", "text/html", "\n")
	fmt.Fprint(conn, "\n")
	fmt.Fprint(conn, body)
}
