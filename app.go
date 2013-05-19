package main

import(
	"fmt"
	"./web"
)

func main() {
	fmt.Println("goblog")

	app := web.CreateServer()
	router := app.Router()
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