package routes

import (
	"github.com/kataras/iris/v12"
)

// RegisterRoutes routes.
func RegisterRoutes(app *iris.Application) {

	registerSimpleRoute(app)
	registerOfflineRoute(app)
	registerGroupRoute(app)
	registerPathParamRoute(app)

	// Method:    GET
	// Resource:  http://localhost:8080/user/42
	//
	// Need to use a custom regexp instead?
	// Easy;
	// Just mark the parameter's type to 'string'
	// which accepts anything and make use of
	// its `regexp` macro function, i.e:
	// app.Get("/user/{id:string regexp(^[0-9]+$)}")
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

}
