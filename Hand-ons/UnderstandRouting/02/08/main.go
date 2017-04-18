package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
	"strings"
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
	var idx int
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if idx == 0 {
			data := strings.Fields(ln)
			fmt.Printf("Method : %v\n", data[0])
			fmt.Printf("URI : %v\n", data[1])
		}
		if ln == "" {
			fmt.Println("this is the end of Headers")
			break
		}
		idx++
	}

	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Exercises</title>
	</header>
	<body>
		<h1>HOLY COW THIS IS LOW LEVEL</h1>
	</body>
	</html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
	conn.Close()
}