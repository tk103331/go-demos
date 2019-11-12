package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/view"
	"time"
)

// SimpleUser struct.
type SimpleUser struct {
	Name  string
	Age   int
	Birth time.Time
}

func registerViewEngineRoute(app *iris.Application) {
	viewDir := "views"
	engines := map[string]view.Engine{
		"html":       iris.HTML(viewDir, ".html").Reload(true),
		"django":     iris.Django(viewDir, ".django").Reload(true),
		"handlebars": iris.Handlebars(viewDir, ".handlebars").Reload(true),
		"amber":      iris.Amber(viewDir, ".amber").Reload(true),
		"pug":        iris.Pug(viewDir, ".pug").Reload(true),
		"jet":        iris.Jet(viewDir, ".jet").Reload(true),
	}

	for k, v := range engines {
		app.RegisterView(v)
		ext := v.Ext()

		func(n string) {
			app.Get("/view/"+k, func(ctx iris.Context) {
				ctx.Application().Logger().Info(ctx.Path() + " => simple" + ext)
				ctx.ViewData("engine", n)
				ctx.ViewData("title", "Information")
				ctx.ViewData("user", SimpleUser{Name: "Jack", Age: 21, Birth: time.Now()})
				ctx.View("simple" + ext)
			})
		}(k)

	}
}
