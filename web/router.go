package web

import (
	"fmt"
	"web-server/communication"
)

type RouteDictionary map[string]func(req Request) string

type DefaultRouter struct {
	Port int
	Addr string
	_dic RouteDictionary
}

func (router *DefaultRouter) AddRoute(verb Verb, path string, callback func(req Request) string) error {
	key := buildKey(verb, path)
	router._dic[key] = callback
	return nil
}

func (router *DefaultRouter) router(request *Request) string {
	key := buildKey(request.Verb, request.Path)
	if callbak, ok := router._dic[key]; ok {
		return callbak(*request)
	}
	return "Not Found"
}

func (router *DefaultRouter) Run() error {
	socket, err := communication.InitSocket(router.Addr, uint(router.Port), communication.TCP4)

	if err != nil {
		fmt.Println(err)
		return err
	}

	for {
		con, err := socket.Listen()
		if err != nil {
			fmt.Println(err)
			return err
		}

		content, err := communication.RecvMessage(con)
		if err != nil {
			fmt.Println(err)
			con.Close()
			continue
		}

		req, err := ParseString(content)
		if err != nil {
			fmt.Println(err)
			con.Close()
			continue
		}

		result := router.router(req)
		communication.SendMessage(con, result)
		communication.CloseConnection(con)
	}
	return nil
}

func CreateRouter(addr string, port int) DefaultRouter {
	return DefaultRouter{
		Port: port,
		Addr: addr,
		_dic: make(RouteDictionary),
	}
}

func buildKey(verb Verb, path string) string {
	key := fmt.Sprintf("%s;%s", verb, path)
	return key
}
