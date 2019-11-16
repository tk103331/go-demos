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
	registerHTTPErrorRoute(app)
	registerSubdomainRoute(app)
	registerWrapRouterRoute(app)
	registerOverrideContextRoute(app)
	registerVersioningRoute(app)
	registerContentNegotiationRoute(app)
	registerResponseRecorderRoute(app)
	registerReferrerRoute(app)
	registerJWTAuthenticationRoute(app)
	registerQueryParameterRoute(app)
	registerFormDataRoute(app)
	registerValidationRoute(app)
	registerCacheRoute(app)
	registerFileServerRoute(app)
	registerViewEngineRoute(app)
	registerCookieRoute(app)
	registerSessionRoute(app)
	registerSessionDBRoute(app)
	registerWebsocketRoute(app)
	registerDependencyInjectionRoute(app)
}
