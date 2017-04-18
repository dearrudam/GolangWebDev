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
	var method string
	var uri string
	for scanner.Scan() {
		scanner.Scan()
		ln := scanner.Text()
		fmt.Println(ln)
		if idx == 0 {
			data := strings.Fields(ln)
			fmt.Printf("Method : %v\n", data[0])
			fmt.Printf("URI : %v\n", data[1])
			method = data[0]
			uri = data[1]
		}
		if ln == "" {
			fmt.Println("this is the end of Headers")
			break
		}
		idx++
	}

	switch  {
	case strings.ToUpper(method) == "GET" && uri == "/":
		processIndex(conn)
	case strings.ToUpper(method) == "GET" && uri == "/apply":
		processApplyForm(conn)
	case strings.ToUpper(method) == "POST" && uri == "/apply":
		processAppliedForm(conn)
	default:
		processAny(conn)
	}
	conn.Close()
}

func processIndex(conn net.Conn) {

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Exercises</title>
		</header>
		<body>
			<h1>Index</h1>
			<ul>
			   <li><a href="/">Index</a></li>
			   <li><a href="/apply">Applying data</a></li>
			</ul>
		</body>
		</html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
func processApplyForm(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Exercises</title>
		</header>
		<body>
			<h1>Applying data form</h1>
			<ul>
			   <li><a href="/">Index</a></li>
			   <li><a href="/apply">Applying data</a></li>
			</ul>
			<form method="post" action="/apply">
				<input type="text" placeholder="Insert some text here!" name="data" />
				<input type="submit" />
			</form>
		</body>
		</html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func processAppliedForm(conn net.Conn) {

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Exercises</title>
		</header>
		<body>
			<h1>Applied data</h1>
			<ul>
			   <li><a href="/">Index</a></li>
			   <li><a href="/apply">Applying data</a></li>
			</ul>
		</body>
		</html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func processAny(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Exercises</title>
	</header>
	<body>
		<h1>Unknowed place</h1>
		<ul>
		   <li><a href="/">Index</a></li>
		   <li><a href="/apply">Applying data</a></li>
		</ul>
	</body>
	</html>`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}