package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/versioning"
)

func registerVersioningRoute(app *iris.Application) {
	myCustomNotVersionFound := func(ctx iris.Context) {
		ctx.StatusCode(404)
		ctx.Writef("%s version not found", versioning.GetVersion(ctx))
	}

	myMiddleware := func(ctx iris.Context) {
		ctx.Application().Logger().Info("myMiddleware")
		ctx.Next()
	}

	sendHandler := func(v string) func(ctx iris.Context) {
		return func(ctx iris.Context) {
			ctx.Writef("version : %s", v)
		}
	}
	// Using the versioning.Deprecated(handler iris.Handler, options versioning.DeprecationOptions) iris.Handler function
	// you can mark a specific handler version as deprecated.
	v0xHandler := versioning.Deprecated(sendHandler("v0.x"), versioning.DeprecationOptions{
		// if empty defaults to: "WARNING! You are using a deprecated version of this API."
		// WarnMessage string
		// DeprecationDate time.Time
		// DeprecationInfo string
	})

	userAPI := app.Party("/api/user")
	// The versioning.NewMatcher(versioning.Map) iris.Handler creates a single handler
	// which decides what handler need to be executed based on the requested version.
	userAPI.Get("/", myMiddleware, versioning.NewMatcher(versioning.Map{
		"0.9":               v0xHandler,
		"1.0":               sendHandler("v1.0"),
		">= 2, < 3":         sendHandler("v2.x"),
		versioning.NotFound: myCustomNotVersionFound,
	}))

	postAPIV09 := versioning.NewGroup("0.9").Deprecated(versioning.DefaultDeprecationOptions)
	postAPIV09.Get("/", sendHandler("v0.9"))
	postAPIV10 := versioning.NewGroup("1.0")
	postAPIV10.Get("/", sendHandler("v1.0"))

	postAPIV2 := versioning.NewGroup(">= 2, < 3")
	postAPIV2.Get("/", sendHandler("v2.x"))
	postAPIV2.Post("/", sendHandler("v2.x"))
	postAPIV2.Put("/other", sendHandler("v2.x"))

	versioning.RegisterGroups(app.Party("/api/post"), versioning.NotFoundHandler, postAPIV09, postAPIV10, postAPIV2)

	// Compare version manually from inside your handlers
	// reports if the "version" is matching to the "is".
	// the "is" can be a constraint like ">= 1, < 3".
	// If(version string, is string) bool
	// same as `If` but expects a Context to read the requested version.
	// Match(ctx iris.Context, expectedVersion string) bool
	app.Get("/api/manual", func(ctx iris.Context) {
		if versioning.Match(ctx, ">= 2.2.3") {
			// [logic for >= 2.2.3 version of your handler goes here]
			ctx.WriteString(">= 2.2.3 version")
			return
		}
	})
}
