package web

import "fmt"

type RouteDictionary map[string]func(req Request) string

var _dic RouteDictionary = make(RouteDictionary)

func AddRoute(verb Verb, path string, callback func(req Request) string) error {
	key := buildKey(verb, path)
	_dic[key] = callback
	return nil
}

func Router(request *Request) string {
	key := buildKey(request.Verb, request.Path)
	callbak := _dic[key]
	return callbak(*request)
}

func buildKey(verb Verb, path string) string {
	key := fmt.Sprintf("%s;%s", verb, path)
	return key
}
