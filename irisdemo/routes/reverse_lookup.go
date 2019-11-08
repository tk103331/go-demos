package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func registerReverseLookupRoute(app *iris.Application) {

	routePathReverser := router.NewRoutePathReverser(app,
		router.WithHost("example.com"),
		router.WithScheme("http"),
		// router.WithServer(srv)
	)

	h := func(ctx iris.Context) {
		ctx.HTML("<b>Hi</b1>")
	}
	app.Get("/about", h).Name = "about"
	app.Get("/page/{id}", h).Name = "page"

	app.Get("/reverse_lookup", func(ctx iris.Context) {
		// GetRoutes function to get all registered routes
		// GetRoute(routeName string) method to retrieve a route by name
		aboutRoute := app.GetRoute("about")
		println(aboutRoute)
		pagePath := routePathReverser.Path("page", 111)
		println(pagePath)
		pageURL := routePathReverser.URL("page", 222)
		println(pageURL)

		ctx.View("reverse_lookup.html")
	})
}
