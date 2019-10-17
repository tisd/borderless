package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		panic("host:port not provided")
	}

	CONNECT := os.Args[1]
	conn, err := net.Dial("tcp4", CONNECT)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	for {
		fmt.Fprintf(conn, "DATA FROM CLIENT\n")
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server -> " + msg)
		time.Sleep(time.Second * 3)
	}
}
