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
	registerCustomMacroRoute(app)
	registerReverseLookupRoute(app)
	registerMiddlewareRoute(app)
	registerHttpErrorRoute(app)
	registerSubdomainRoute(app)
	registerWrapRouterRoute(app)
	registerOverrideContextRoute(app)
	registerVersioningRoute(app)
}
