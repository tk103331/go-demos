package routes

import (
	"github.com/kataras/iris/v12"
)

func registerResponseRecorderRoute(app *iris.Application) {
	// start record.
	app.Use(func(ctx iris.Context) {
		ctx.Record()
		ctx.Next()
	})

	// collect and "log".
	app.Done(func(ctx iris.Context) {
		recorder := ctx.Recorder()
		body := string(recorder.Body())

		recorder.ResetBody()
		recorder.WriteString("result:")
		recorder.WriteString(body)
		recorder.Header().Add("reset-body", "true")
		// Should print success.
		app.Logger().Infof("sent: %s", string(body))
	})

	app.Get("/save", func(ctx iris.Context) {
		ctx.WriteString("success")
		ctx.Next() // calls the Done middleware(s).
	})
}
