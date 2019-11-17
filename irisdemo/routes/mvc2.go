package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func registerMvc2Route(app *iris.Application) {
	// Serve a controller based on the root Router, "/".
	mvc.New(app.Party("/mvc2")).Handle(new(ExampleController))
}

// ExampleController serves the "/", "/ping" and "/hello".
type ExampleController struct{}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

// GetPing serves
// Method:   GET
// Resource: http://localhost:8080/ping
func (c *ExampleController) GetPing() string {
	return "pong"
}

// GetHello serves
// Method:   GET
// Resource: http://localhost:8080/hello
func (c *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

// BeforeActivation called once, before the controller adapted to the main application
// and of course before the server ran.
// After version 9 you can also add custom routes for a specific controller's methods.
// Here you can register custom method's handlers
// use the standard router with `ca.Router` to
// do something that you can do without mvc as well,
// and add dependencies that will be binded to
// a controller's fields or method function's input arguments.
func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}

	b.Handle(
		"GET",
		"/custom_path",
		"CustomHandlerWithoutFollowingTheNamingGuide",
		anyMiddlewareHere,
	)

	// or even add a global middleware based on this controller's router,
	// which in this example is the root "/":
	// b.Router().Use(myMiddleware)
}

// CustomHandlerWithoutFollowingTheNamingGuide serves
// Method:   GET
// Resource: http://localhost:8080/custom_path
func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}

// GetUserBy serves
// Method:   GET
// Resource: http://localhost:8080/user/{username:string}
// By is a reserved "keyword" to tell the framework that you're going to
// bind path parameters in the function's input arguments, and it also
// helps to have "Get" and "GetBy" in the same controller.
//
func (c *ExampleController) GetUserBy(username string) mvc.Result {
	return mvc.View{
		Name: "username.html",
		Data: username,
	}
}

/* Can use more than one, the factory will make sure
that the correct http methods are being registered for each route
for this controller, uncomment these if you want:

func (c *ExampleController) Post() {}
func (c *ExampleController) Put() {}
func (c *ExampleController) Delete() {}
func (c *ExampleController) Connect() {}
func (c *ExampleController) Head() {}
func (c *ExampleController) Patch() {}
func (c *ExampleController) Options() {}
func (c *ExampleController) Trace() {}
*/

/*
func (c *ExampleController) All() {}
//        OR
func (c *ExampleController) Any() {}



func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
    // 1 -> the HTTP Method
    // 2 -> the route's path
    // 3 -> this controller's method name that should be handler for that route.
    b.Handle("GET", "/mypath/{param}", "DoIt", optionalMiddlewareHere...)
}

// After activation, all dependencies are set-ed - so read only access on them
// but still possible to add custom controller or simple standard handlers.
func (c *ExampleController) AfterActivation(a mvc.AfterActivation) {}

*/
