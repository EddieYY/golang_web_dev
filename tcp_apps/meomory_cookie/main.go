package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ls, err := net.Listen("tcp", ":8080")
	defer ls.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintln(conn,
		"In Meomory database\n"+
			"Please Use:\n"+
			"SET key value \n"+
			"GET key \n"+
			"DEL key \n"+
			"For example: \n"+
			" SET favfood Chocolate \n"+
			" GET favfood\n")

	s := bufio.NewScanner(conn)
	data := make(map[string]string)
	for s.Scan() {
		str := strings.Fields(s.Text())
		if len(str) < 1 {
			continue
		}
		switch str[0] {
		case "SET":
			if len(str) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE")
				continue
			}

			data[str[1]] = str[2]
		case "GET":
			fmt.Fprintf(conn, "%s\n", data[str[1]])
		case "DEL":
			delete(data, str[1])
		default:
			fmt.Fprintln(conn, "INVALID COMMAND "+str[0])
			continue
		}
	}
}
