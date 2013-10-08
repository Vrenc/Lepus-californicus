package main

import (
	"fmt"
	"vrenc/web"
)

func main() {
	fmt.Println("goblog")

	app := web.CreateServer()
	router := app.Router()

	router.Get("/", func(res *web.Response, req *web.Request) {
		view := new(web.View)
		data := struct {
			Name    string
			Message string
		}{
			"vrenc",
			"sup?",
		}
		res.Send(view.Render("templates/index.html", data))
	})
	router.Get("/abc", func(res *web.Response, req *web.Request) {
		res.Send("abcdefghijklmnopqrstuvwxyz")
	})
	router.Get("/user/:id", func(res *web.Response, req *web.Request) {
		res.Send(req.Param("id"))
	})
	router.Get("/post/:user/:id", func(res *web.Response, req *web.Request) {
		res.Send(req.Param("user") + ", " + req.Param("id"))
	})
	router.Get("/json", func(res *web.Response, req *web.Request) {
		data := struct {
			Posts []interface{}
			Count int
			Date  string
		}{
			[]interface{}{
				struct {
					Message string
					User    string
				}{
					"Wat is deze?",
					"vrenc",
				},
				struct {
					Message string
					User    string
				}{
					"jeweetzelluf",
					"freek",
				},
				struct {
					Message string
					User    string
				}{
					"#YOLO #SWAG",
					"jesus",
				},
			},
			3,
			"19-05-2013",
		}

		res.Json(data)
	})

	app.Run(3000)
}
