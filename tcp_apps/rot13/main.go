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
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	s := bufio.NewScanner(conn)
	for s.Scan() {
		sByte := []byte(strings.ToLower(s.Text()))
		outByte := make([]byte, len(sByte))
		for i, v := range sByte {
			if v == 32 {
				outByte[i] = v
				continue
			}
			if v < 110 {
				outByte[i] = v + 13
			}
			if v >= 110 {
				outByte[i] = v - 13
			}
		}
		fmt.Fprintf(conn, "%s => %s\n", sByte, outByte)
	}
}
