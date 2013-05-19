package web

import(
	"fmt"
	"net/http"
)

type Response struct {
	writer http.ResponseWriter
}

func (res *Response) Send(text string) {
	fmt.Fprintf(res.writer, text)
}

func (res *Response) NotFound() {
	res.writer.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(res.writer, "404 Not found")
}