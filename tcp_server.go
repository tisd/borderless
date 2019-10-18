package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var count = 0

type state struct {
	MouseX float64 `json:"mouseX"`
	MouseY float64 `json:"mouseY"`
}

func (s state) String() string {
	return fmt.Sprintf("%d, %d", s.MouseX, s.MouseY)
}

var serverState state = state{MouseX: 15, MouseY: 15}

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
		// data, err := bufio.NewReader(conn).ReadString('\n')
		encoder, decoder := json.NewEncoder(conn), json.NewDecoder(conn)
		var data state
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}

		fmt.Println("Client"+strconv.Itoa(currentCount)+" -> ", data)

		fmt.Println("Server >> ", serverState)
		err = encoder.Encode(&serverState)
		if err != nil {
			panic(err)
		}
		// fmt.Fprintf(conn, "DATA FROM SERVER TO CLIENT"+strconv.Itoa(currentCount)+"\n")
		time.Sleep(time.Second * 3)
	}
}
