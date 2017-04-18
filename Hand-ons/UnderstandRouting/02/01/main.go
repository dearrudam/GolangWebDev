package main

import (
	"net"
	"log"
	"io"
)

func main() {
	l, err := net.Listen("tcp", ":8080");
	if err != nil {
		log.Panic(err)
	}

	defer l.Close()

	for {
		c,err:=l.Accept()
		if err!=nil {
			log.Panic(err)
		}
		io.WriteString(c,"I see you connected.")
		c.Close()
	}

}
