package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type state struct {
	MouseX float64 `json:"mouseX"`
	MouseY float64 `json:"mouseY"`
}

func (s state) String() string {
	return fmt.Sprintf("%d, %d", s.MouseX, s.MouseY)
}

var clientState state = state{MouseX: 17, MouseY: 17}

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

	endcoder, decoder := json.NewEncoder(conn), json.NewDecoder(conn)
	for {
		fmt.Println("Client >> ", clientState)
		err = endcoder.Encode(&clientState)
		if err != nil {
			panic(err)
		}
		err = decoder.Decode(&clientState)
		if err != nil {
			panic(err)
		}
		fmt.Println("Server -> ", clientState)
		time.Sleep(time.Second * 3)
	}
}
