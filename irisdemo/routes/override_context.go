package routes

import (
	"reflect"

	"github.com/kataras/iris/v12"
	// 1.
	"github.com/kataras/iris/v12/context"
)

// MyContext is custom Context
// 2.
// Create your own custom Context, put any fields you'll need.
type MyContext struct {
	// Embed the `iris.Context` -
	// It's totally optional but you will need this if you
	// don't want to override all the context's methods!
	iris.Context
}

// Optionally: validate MyContext implements iris.Context on compile-time.
var _ iris.Context = &MyContext{}

// Do implements
// 3.
func (ctx *MyContext) Do(handlers context.Handlers) {
	context.Do(ctx, handlers)
}

// Next implements
// 3.
func (ctx *MyContext) Next() {
	context.Next(ctx)
}

// HTML override
// [Override any context's method you want here...]
// Like the HTML below:
func (ctx *MyContext) HTML(htmlContents string, args ...interface{}) (int, error) {
	ctx.Application().Logger().Infof("Executing .HTML function from MyContext")
	ctx.ContentType("text/html")
	ctx.Header("override-context-header", "override-context-header-value")
	return ctx.Writef(htmlContents, args)
}

func registerOverrideContextRoute(app *iris.Application) {
	// 4.
	app.ContextPool.Attach(func() iris.Context {
		return &MyContext{
			// If you use the embedded Context,
			// call the `context.NewContext` to create one:
			Context: context.NewContext(app),
		}
	})

	// Register your route, as you normally do
	app.Handle("GET", "/override_context", recordWhichContextForExample,
		func(ctx iris.Context) {
			println(reflect.TypeOf(ctx).String())
			// use the context's overridden HTML method.
			ctx.HTML("<h1> Hello from my custom context's HTML! </h1>", "ccc")
		})

	// This will be executed by the
	// MyContext.Context embedded default context
	// when MyContext is not directly define the View function by itself.
	app.Handle("GET", "/hi/{firstname:alphabetical}", recordWhichContextForExample,
		func(ctx iris.Context) {
			firstname := ctx.Params().GetString("firstname")
			println(reflect.TypeOf(ctx).Elem().String())
			ctx.ViewData("firstname", firstname)
			ctx.Gzip(true)

			ctx.View("hi.html")
		})
}

// Should always print "($PATH) Handler is executing from 'MyContext'"
func recordWhichContextForExample(ctx iris.Context) {
	ctx.Application().Logger().Infof("(%s) Handler is executing from: '%s'",
		ctx.Path(), reflect.TypeOf(ctx).Elem().String())

	ctx.Next()
}
