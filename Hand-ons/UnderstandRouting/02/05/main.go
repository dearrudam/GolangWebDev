package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
)

func main() {
	l, err := net.Listen("tcp", ":8080");
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)
	}

}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn);
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}

	io.WriteString(conn, "I see you connected.")
	conn.Close()
}