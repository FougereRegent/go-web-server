package main

import (
	"fmt"
	"os"
	"web-server/web"
)

func main() {
	fmt.Println("=== Web Server ===")

	router := web.CreateRouter("0.0.0.0", 8080)
	router.AddRoute(web.GET, "/", func(req web.Request) web.Response {

		header := make(web.Header)
		header["server-name"] = "go-http"

		file, _ := os.Open("test.jpg")

		res := make([]byte, 2000000)
		file.Read(res)

		return web.Response{
			HttpStatus:     web.OK,
			HeaderResponse: header,
			Body:           res,
		}
	})
	router.AddRoute(web.GET, "/version", func(req web.Request) web.Response {
		header := make(web.Header)
		header["server-name"] = "go-http"
		return web.Response{
			HttpStatus:     web.OK,
			HeaderResponse: header,
			Body:           []byte("<h1>V1.0.0</h1>"),
		}
	})

	err := router.Run()

	if err != nil {
		os.Exit(-1)
	}

	os.Exit(0)
}
