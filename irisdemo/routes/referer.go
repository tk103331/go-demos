package routes

import (
	"github.com/kataras/iris/v12"
)

func registerReferrerRoute(app *iris.Application) {
	app.Get("/referrer", func(ctx iris.Context) {
		r := ctx.GetReferrer()
		switch r.Type {
		case iris.ReferrerSearch:
			ctx.Writef("Search %s: %s\n", r.Label, r.Query)
			ctx.Writef("Google: %s\n", r.GoogleType)
		case iris.ReferrerSocial:
			ctx.Writef("Social %s\n", r.Label)
		case iris.ReferrerIndirect:
			ctx.Writef("Indirect: %s\n", r.URL)
		case iris.ReferrerDirect:
			ctx.Writef("Direct: %s\n", r.URL)
		case iris.ReferrerEmail:
			ctx.Writef("Email: %s\n", r.URL)
		case iris.ReferrerInvalid:
			ctx.Writef("Invalid: %s\n", r.URL)
		}
	})
}
