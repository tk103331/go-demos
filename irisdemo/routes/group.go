package routes

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func registerGroupRoute(app *iris.Application) {

	users1 := app.Party("/users1", myAuthMiddlewareHandler1)

	// /users1/profile
	users1.Get("/profile", userProfileHandler)
	// /users1/messages
	users1.Get("/messages", userMessageHandler)

	app.PartyFunc("/users2", func(users2 iris.Party) {
		users2.Use(myAuthMiddlewareHandler2)

		// /users2/profile
		users2.Get("/profile", userProfileHandler)
		// /users2/messages
		users2.Get("/messages", userMessageHandler)
	})
}

func myAuthMiddlewareHandler1(ctx iris.Context) {
	println(fmt.Sprintf("Request party 'users1' from method: %s and path: %s\n", ctx.Method(), ctx.Path()))
	// go next handler
	ctx.Next()
}
func myAuthMiddlewareHandler2(ctx iris.Context) {
	println(fmt.Sprintf("Request party 'users2' from method: %s and path: %s\n", ctx.Method(), ctx.Path()))
	// go next handler
	ctx.Next()
}

func userProfileHandler(ctx iris.Context) {
	ctx.Writef("Hello from method: %s and path: %s\n", ctx.Method(), ctx.Path())
}

func userMessageHandler(ctx iris.Context) {
	ctx.Writef("Hello from method: %s and path: %s\n", ctx.Method(), ctx.Path())
}
