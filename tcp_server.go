package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var count = 0

func main() {
	if len(os.Args) == 1 {
		panic("Port number not provided")
	}

	PORT := os.Args[1]
	listener, err := net.Listen("tcp4", ":"+PORT)

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		count++
		fmt.Print("Client Connected")
		go handleConnection(conn, count)
	}
}

func handleConnection(conn net.Conn, currentCount int) {
	defer conn.Close()
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Print("Client"+strconv.Itoa(currentCount)+" -> ", string(data))
		fmt.Fprintf(conn, "DATA FROM SERVER TO CLIENT"+strconv.Itoa(currentCount)+"\n")
		time.Sleep(time.Second * 3)
	}
}
