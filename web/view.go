package web

import(
	"html/template"
	"bytes"
)

type View struct {
	temp *template.Template
}

func (v *View) Render(file string, data interface{}) string {
	v.temp, _ = template.ParseFiles(file)

	html := new(bytes.Buffer) 
	v.temp.Execute(html, data)

	return string(html.Bytes())
}

