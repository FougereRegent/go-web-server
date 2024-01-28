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

	web.AddRoute(web.GET, "/", func(req web.Request) string {
		fmt.Println("First callback")
		return "First callback"
	})
	web.AddRoute(web.GET, "/test", func(req web.Request) string {
		fmt.Println("Second Callback")
		return "Second callback"
	})

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

	req, err := web.ParseString(content)
	result := web.Router(req)
	err = communication.SendMessage(con, result)

	con.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

/*Function to display a message*/
func displayReq(request *web.Request) {
	fmt.Println(request.HttpVersion)
	fmt.Println(request.Verb)
	fmt.Println(request.Path)

	for key, val := range request.Header {
		fmt.Printf("%s : %s\n", key, val)
	}
}
