package web

type Request struct {
	HttpVersion string
	Path        string
	Verb        Verb
	Header      Header
	Body        string
}
