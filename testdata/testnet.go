package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	i := 0
	defer func(){
		i++
		fmt.Println("test ", i)
	}()

	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Status: ", status)
	}
}