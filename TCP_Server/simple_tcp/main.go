// Package main create a server that returns the URL of the GET request.
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	defer ln.Close()

	if err != nil {
		fmt.Print(err)
	}
	for {
		conn, err := ln.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Print(err)
			continue
		}

		// writer go io.WriteString(conn, "Hello World")
		// reader //go readhandle(conn)
		go rwhandle(conn)

	}

}

// Readhanle is a read tcp handle
func readhandle(conn net.Conn) {
	Scanner := bufio.NewScanner(conn)
	for Scanner.Scan() {
		fmt.Println(Scanner.Text())
	}
	defer conn.Close()
}

// Read & Write handle

func rwhandle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(conn)
	for s.Scan() {
		fmt.Println(s.Text())
		fmt.Fprintln(conn, "i known you say '"+s.Text()+"'")
		defer conn.Close()

		fmt.Println("sorry, it's timeOut")
	}

}
