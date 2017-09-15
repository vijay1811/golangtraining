package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

func handleReadC(conn net.Conn) {
	for {
		rdr := bufio.NewReader(conn)
		msg, _ := rdr.ReadString('\n')
		fmt.Print(msg)
	}
}
func handleWriteC(conn net.Conn) {
	for {
		var s string
		scr := bufio.NewScanner(os.Stdin)
		scr.Scan()
		s = scr.Text()
		fmt.Fprintf(conn, s+"\n")
	}
}
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err == io.EOF {
		fmt.Println("Connection terminated by Server")
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func(conn net.Conn) {
		handleReadC(conn)
		wg.Done()
	}(conn)

	go func(conn net.Conn) {
		handleWriteC(conn)
		wg.Done()
	}(conn)
	wg.Wait()
}
