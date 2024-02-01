package web

import (
	"fmt"
	"web-server/communication"
)

type RouteDictionary map[string]func(req Request) Response

type DefaultRouter struct {
	Port int
	Addr string
	_dic RouteDictionary
}

func (router *DefaultRouter) AddRoute(verb Verb, path string, callback func(req Request) Response) error {
	key := buildKey(verb, path)
	router._dic[key] = callback
	return nil
}

func (router *DefaultRouter) router(request *Request) Response {
	key := buildKey(request.Verb, request.Path)
	if callbak, ok := router._dic[key]; ok {
		return callbak(*request)
	}
	return Response{
		HttpStatus:     NOT_FOUND,
		Body:           "<h1>Not Found</h1>",
		HeaderResponse: nil,
	}
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
		str_result := CreateResponse(result)
		communication.SendMessage(con, str_result)
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
