package main

import (
	"github.com/kataras/iris/v12"
	"github.com/tk103331/go-demos/irisdemo/routes"
)

func main() {
	app := iris.New()

	// Load all templates from the "./views" folder
	// where extension is ".html" and parse them
	// using the standard `html/template` package.
	app.RegisterView(iris.HTML("./views", ".html"))

	// Register routes.
	routes.RegisterRoutes(app)

	// A new host
	// go func() {
	// 	app.NewHost(&http.Server{Addr: ":9090"}).ListenAndServe()
	// }()

	app.ConfigureHost(func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {

			println("server terminated: " + h.Server.Addr)
		})
	})

	config := iris.WithConfiguration(iris.YAML("./iris.yml"))
	app.Configure(iris.WithoutPathCorrectionRedirection)
	// Start the server using a network address.
	app.Run(iris.Addr(":9527"), config)
}
