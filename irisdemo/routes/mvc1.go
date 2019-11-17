package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func registerMvc1Route(app *iris.Application) {
	mvc.Configure(app.Party("/root"), myMVC)
}
func myMVC(app *mvc.Application) {
	// app.Register(...)
	// app.Router.Use/UseGlobal/Done(...)
	app.Handle(new(MyController))
}

type MyController struct{}

func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done
	// and any standard Router API call you already know

	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the MyCustomHandler
	b.Handle("GET", "/something/{id:long}", "MyCustomHandler", anyMiddleware...)
}

// GET: http://localhost:8080/root
func (m *MyController) Get() string {
	return "Hey"
}

// GET: http://localhost:8080/root/something/{id:long}
func (m *MyController) MyCustomHandler(id int64) string {
	return "MyCustomHandler says Hey"
}
