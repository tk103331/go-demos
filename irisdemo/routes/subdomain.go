package routes

import (
	"github.com/kataras/iris/v12"
)

func registerSubdomainRoute(app *iris.Application) {
	admin := app.Subdomain("admin")
	// add hosts:
	// 127.0.0.1 mydomain.com
	// 127.0.0.1 admin.mydomain.com
	// admin.mydomain.com
	admin.Get("/", func(ctx iris.Context) {
		ctx.Writef("INDEX FROM admin.mydomain.com")
	})

	// admin.mydomain.com/hey
	admin.Get("/hey", func(ctx iris.Context) {
		ctx.Writef("HEY FROM admin.mydomain.com/hey")
	})

	wildcard := app.WildcardSubdomain()
	wildcard.Get("/hi", func(ctx iris.Context) {
		method := ctx.Method()
		subdomain := ctx.Subdomain()
		path := ctx.Path()

		ctx.Writef("\nInfo\n\n")
		ctx.Writef("Method: %s\nSubdomain: %s\nPath: %s", method, subdomain, path)
	})

	// register a global redirection rule for subdomains as well.
	www := app.Subdomain("www")
	app.SubdomainRedirect(www, app)
}
