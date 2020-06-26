package net

import (
	"bufio"
	"fmt"
	//"machine"
)

func main() {
	conn, err := Dial("tcp", "golang.org:80")
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
