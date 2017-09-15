package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func handleReadS(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			fmt.Println("Connection terminated by client")
			return
		} else if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print(msg)
		//strings.TrimSpace(msg) - Done
		msg = strings.TrimSpace(msg)
		conn.Write([]byte(msg + " : received by server\n"))
	}
}
func handleWriteS(conn net.Conn) {
	for {
		var s string
		scr := bufio.NewScanner(os.Stdin)
		if scr.Scan() {
			s = scr.Text()
			fmt.Fprintf(conn, s+"\n")
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	errCount := 0
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			errCount++
			if errCount > 3 {
				fmt.Println("Repeatedly error generated while accepting connection")
				return
			}
			continue
		}
		go handleReadS(conn)
		go handleWriteS(conn)
	}

}
