package main

import (
	"fmt"
	"os"
	"web-server/web"
)

func main() {
	fmt.Println("=== Web Server ===")

	router := web.CreateRouter("0.0.0.0", 8080)
	router.AddRoute(web.GET, "/", func(req web.Request) string {
		return "Hello World"
	})
	router.AddRoute(web.GET, "/version", func(req web.Request) string {
		return "v1.0"
	})

	err := router.Run()

	if err != nil {
		os.Exit(-1)
	}

	os.Exit(0)
}
