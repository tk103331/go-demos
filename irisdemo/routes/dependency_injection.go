package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

type Service interface {
	SayHello(to string) string
}

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type myService struct {
	prefix string
}

func (s *myService) SayHello(to string) string {
	return hello(to)
}

func hello(to string) string {
	return "Hello " + to
}

func helloService(to string, service Service) string {
	return service.SayHello(to)
}

func dologin(form LoginForm) string {
	return "Login " + form.Username
}

func registerDependencyInjectionRoute(app *iris.Application) {
	// Path Parameters - Built-in Dependencies
	app.Get("/di/hello1/{to:string}", hero.Handler(hello))
	// Services - Static Dependencies
	hero.Register(&myService{prefix: "Service: Hello"})
	app.Get("/di/hello2/{to:string}", hero.Handler(helloService))
	// Per-Request - Dynamic Dependencies
	hero.Register(func(ctx iris.Context) (form LoginForm) {
		// it binds the "form" with the
		// x-www-form-urlencoded from data,
		// sent by the client, and returns it.
		ctx.ReadForm(form)
		return
	})
	app.Post("/di/login", hero.Handler(dologin))
}
