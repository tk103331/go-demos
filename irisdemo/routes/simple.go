package routes

import (
	"github.com/kataras/iris/v12"
)

func registerSimpleRoute(app *iris.Application) {
	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hello.html
		ctx.View("hello.html")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	// Method: "GET"
	app.Get("/method", handler)

	// Method: "POST"
	app.Post("/method", handler)

	// Method: "PUT"
	app.Put("/method", handler)

	// Method: "DELETE"
	app.Delete("/method", handler)

	// Method: "OPTIONS"
	app.Options("/method", handler)

	// Method: "TRACE"
	app.Trace("/method", handler)

	// Method: "CONNECT"
	app.Connect("/method", handler)

	// Method: "HEAD"
	app.Head("/method", handler)

	// Method: "PATCH"
	app.Patch("/method", handler)

	// register the route for all HTTP Methods
	app.Any("/method", handler)

}

func handler(ctx iris.Context) {
	ctx.Writef("Hello from method: %s and path: %s\n", ctx.Method(), ctx.Path())
}
