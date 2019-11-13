package routes

import (
	"github.com/kataras/iris/v12"
)

func registerFileServerRoute(app *iris.Application) {
	app.HandleDir("/static", "./assets", iris.DirOptions{
		ShowList: true,
	})
}
