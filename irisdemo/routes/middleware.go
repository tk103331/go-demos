package routes

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

func registerMiddlewareRoute(app *iris.Application) {
	app.Get("/middleware", before, mainHandler, after)

	// Order of those calls does not matter,
	// `UseGlobal` and `DoneGlobal` are applied to existing routes
	// and future routes also.
	//
	// Remember: the `Use` and `Done` are applied to the current party's and its children,
	// so if we used the `app.Use/Done before the routes registration
	// it would work like UseGlobal/DoneGlobal in this case,
	// because the `app` is the root "Party".
	app.UseGlobal(func(ctx iris.Context) {
		println("Global Before ")
		ctx.Next()
	})
	// You could also use the ExecutionRules to force Done handlers to be executed without the need of ctx.Next() in your route handlers, do it like this:
	app.SetExecutionRules(iris.ExecutionRules{
		// Begin: ...
		// Main:  ...
		Done: iris.ExecutionOptions{Force: true},
	})
	app.DoneGlobal(func(ctx iris.Context) {
		println("Global After ")
		ctx.Next()
	})

	// builtin basicauth middleware

	authConfig := basicauth.Config{
		Users:   map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"},
		Realm:   "Authorization Required", // defaults to "Authorization Required"
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	// to global app.Use(authentication) (or app.UseGlobal before the .Run)
	// to routes
	/*
		app.Get("/mysecret", authentication, h)
	*/
	needAuth := app.Party("/admin", authentication)
	{
		h := func(ctx iris.Context) {
			ctx.WriteString(ctx.Path())
		}
		// /admin
		needAuth.Get("/", h)
		// /admin/profile
		needAuth.Get("/profile", h)

		// /admin/settings
		needAuth.Get("/settings", h)
	}
}

func before(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next() // execute the next handler, in this case the main one.
}

func after(ctx iris.Context) {
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")

	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after".

}
