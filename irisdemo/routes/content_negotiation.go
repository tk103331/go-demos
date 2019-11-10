package routes

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

type testdata struct {
	Name string `json:"name" xml:"Name"`
	Age  int    `json:"age" xml:"Age"`
	City string `json:"city" xml:"city"`
}

func registerContentNegotiationRoute(app *iris.Application) {

	// Render a resource with "gzip" encoding algorithm as application/json or text/xml or application/xml
	// when client's accept header contains one of them
	// or JSON (the first declared) if accept is empty,
	// and when client's accept-encoding header contains "gzip" or it's empty.
	app.Get("/resource", func(ctx iris.Context) {
		data := testdata{
			Name: "test name",
			Age:  26,
			City: "北京",
		}

		ctx.Negotiation().JSON().XML().EncodingGzip().Charset("gbk").Encoding("gbk")

		_, err := ctx.Negotiate(data)
		if err != nil {
			ctx.Writef("error: %v", err)
		}
	})
	// OR define them in a middleware and call Negotiate with nil in the final handler.
	app.Get("/resource2", func(ctx iris.Context) {
		data := testdata{
			Name: "test name",
			Age:  26,
			City: "北京",
		}

		ctx.Negotiation().
			JSON(data).
			XML(data).
			HTML(fmt.Sprintf("<h1>%s</h1><h2>Age %d</h2><h2>%s</h2>", data.Name, data.Age, data.City))

		ctx.Negotiate(nil)
	})
}
