package web

import(
	"fmt"
	"net/http"
	"encoding/json"
)

type Response struct {
	writer http.ResponseWriter
}

func (res *Response) Send(text string) {
	fmt.Fprintf(res.writer, text)
}

func (res *Response) Json(data interface{}) {
	encoder := json.NewEncoder(res.writer)

	encoder.Encode(data)
}

func (res *Response) NotFound() {
	res.writer.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(res.writer, "404 Not found")
}