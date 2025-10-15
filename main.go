package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "youtube.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	request := "GET / HTTP/1.1\r\n" +
		"Host: youtube.com\r\n" +
		"Connection: close\r\n" +
		"\r\n"

	fmt.Fprintf(conn, request)

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error of reading:", err)
		return
	}

	fmt.Println(string(buffer[:n]))
	fmt.Println(string(buffer))
}
