package web

import (
	"errors"
	"fmt"
	"strings"
)

type Verb string
type Header map[string]string

const (
	GET       Verb = "GET"
	PUT       Verb = "PUT"
	DELETE    Verb = "DELETE"
	POST      Verb = "POST"
	HEAD      Verb = "HEAD"
	UNDEFINED Verb = "UNDEFINED"
)

const (
	HTTP_1_1 string = "1.1"
	HTTP_1_0 string = "1.0"
	HTTP_2_0 string = "2.0"
	HTTP_3_0 string = "3.0"
)

const (
	LEN_GET    int = len(GET)
	LEN_PUT    int = len(PUT)
	LEN_DELETE int = len(DELETE)
	LEN_POST   int = len(POST)
	LEN_HEAD   int = len(HEAD)
)

func ParseString(content []byte) (*Request, error) {
	v, nb := getVerb(content)
	if nb < 0 {
		return nil, errors.New("Verb not recognize")
	}

	nb++

	path, size_path := getPath(content[nb:])
	if size_path < 0 {
		return nil, errors.New("Path not recognize")
	}

	nb += size_path + 1

	version, size_version := getVersion(content[nb:])
	if size_version < 0 {
		return nil, errors.New("Not recognized http version")
	}

	nb += size_version + 2

	header, size_header := getHeader(content[nb:])

	if size_header < 0 {
		return nil, errors.New("Header not recognize")
	}

	return &Request{
		Verb:        v,
		Path:        path,
		Header:      header,
		Body:        "test",
		HttpVersion: version}, nil
}

func getVerb(content []byte) (Verb, int) {
	if string(content[0:LEN_GET]) == "GET" {
		return GET, LEN_GET
	} else if string(content[0:LEN_PUT]) == "PUT" {
		return PUT, LEN_PUT
	} else if string(content[0:LEN_DELETE]) == "DELETE" {
		return DELETE, LEN_DELETE
	} else if string(content[0:LEN_HEAD]) == "HEAD" {
		return HEAD, LEN_HEAD
	} else if string(content[0:LEN_POST]) == "POST" {
		return POST, LEN_POST
	} else {
		return UNDEFINED, -1
	}
}

func getPath(content []byte) (string, int) {
	space_index := findIndex(content, ' ')
	if space_index == -1 {
		return "", -1
	}
	return string(content[0:space_index]), space_index
}

func getVersion(content []byte) (string, int) {
	var major, minor int
	V_NAME := "HTTP/"
	LEN_NAME := len(V_NAME)

	if string(content[0:LEN_NAME]) != V_NAME {
		return "", -1
	}
	index_new_line := findIndex(content[LEN_NAME:], '\r')

	if index_new_line < 0 {
		return "", -1
	}

	fmt.Sscanf(string(content[LEN_NAME:LEN_NAME+index_new_line]), "%d.%d", &major, &minor)
	res := fmt.Sprintf("%d.%d", major, minor)

	switch res {
	case HTTP_1_0:
		break
	case HTTP_1_1:
		break
	case HTTP_2_0:
		break
	case HTTP_3_0:
		break
	default:
		return "", -1
	}

	return res, LEN_NAME + len(res)
}

func getHeader(content []byte) (Header, int) {
	data := strings.TrimLeft(string(content), "\r\n\r\n")
	data = strings.ReplaceAll(data, "\r\n\r\n", "")
	tab := strings.Split(data, "\r\n")
	result := make(Header)

	if len(tab) < 1 {
		return nil, -1
	}

	for _, value := range tab {
		head := strings.Split(value, ":")
		result[head[0]] = strings.TrimSpace(strings.Join(head[1:], ":"))
	}

	return result, len(data)
}

func findIndex(content []byte, element byte) int {
	for index, value := range content {
		if value == element {
			return index
		}
	}
	return -1
}

func getBody(content []byte) string {
	return string(content)
}
