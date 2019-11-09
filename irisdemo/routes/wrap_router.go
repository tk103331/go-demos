package routes

import (
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
)

func registerWrapRouterRoute(app *iris.Application) {

	myOtherHandler := func(ctx iris.Context) {
		ctx.Writef("inside a handler which is fired manually by our custom router wrapper")
	}

	// wrap the router with a native net/http handler.
	// if url does not contain any "." (i.e: .css, .js...)
	// (depends on the app , you may need to add more file-server exceptions),
	// then the handler will execute the router that is responsible for the
	// registered routes (look "/" and "/profile/{username}")
	// if not then it will serve the files based on the root "/" path.
	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		path := r.URL.Path

		if strings.HasPrefix(path, "/other") {
			// acquire and release a context in order to use it to execute
			// our custom handler
			// remember: we use net/http.Handler because here
			// we are in the "low-level", before the router itself.
			ctx := app.ContextPool.Acquire(w, r)
			myOtherHandler(ctx)
			app.ContextPool.Release(ctx)
			return
		}

		// else continue serving routes as usual.
		router.ServeHTTP(w, r)
	})

	// Note: In this example we just saw one use case,
	// you may want to .WrapRouter or .Downgrade in order to
	// bypass the Iris' default router, i.e:
	// you can use that method to setup custom proxies too.
}
