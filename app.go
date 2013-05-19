package main

import(
	"fmt"
	"./web"
)

func main() {
	fmt.Println("goblog")

	app := web.CreateServer()
	router := app.Router()

	router.Get("/", func(res *web.Response, req *web.Request) {
		view := new(web.View)
		data := struct {
			Name string
			Message string
		} {
			"vrenc",
			"sup?",
		}
		res.Send(view.Render("./templates/index.html", data))
	})
	router.Get("/abc", func(res *web.Response, req *web.Request) {
		res.Send("abcdefghijklmnopqrstuvwxyz")
	})
	router.Get("/user/:id", func(res *web.Response, req *web.Request) {
		res.Send(req.Param("id"))
	})
	router.Get("/post/:user/:id", func(res *web.Response, req *web.Request) {

	})

	app.Run(4000)
}