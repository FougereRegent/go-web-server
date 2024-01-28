package main

import (
	"fmt"
	"os"
	"web-server/communication"
	"web-server/web"
)

func main() {
	fmt.Println("=== Web Server ===")

	defer os.Exit(0)

	s, err := communication.InitSocket("0.0.0.0", 8090, communication.TCP4)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	con, err := s.Listen()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	content, err := communication.RecvMessage(con)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
