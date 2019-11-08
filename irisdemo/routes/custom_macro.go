package routes

import (
	"github.com/kataras/iris/v12"
	"regexp"
)

func registerCustomMacroRoute(app *iris.Application) {
	latLonExpr := "^-?[0-9]{1,3}(?:\\.[0-9]{1,10})?$"
	latLonRegex, _ := regexp.Compile(latLonExpr)

	// Register your custom argument-less macro function to the :string param type.
	// MatchString is a type of func(string) bool, so we use it as it is.
	app.Macros().Get("string").RegisterFunc("coordinate", latLonRegex.MatchString)

	app.Get("/coordinates/{lat:string coordinate()}/{lon:string coordinate()}",
		func(ctx iris.Context) {
			ctx.Writef("Lat: %s | Lon: %s", ctx.Params().Get("lat"), ctx.Params().Get("lon"))
		})
	// Register your custom macro function which accepts two int arguments.
	app.Macros().Get("string").RegisterFunc("range",
		func(minLength, maxLength int) func(string) bool {
			return func(paramValue string) bool {
				return len(paramValue) >= minLength && len(paramValue) <= maxLength
			}
		})
	// the name should be between 1 and 50 characters length otherwise this handler will not be executed, and response 400.
	app.Get("/limitchar/{name:string range(1,50) else 400}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | the name should be between 1 and 50 characters length otherwise this handler will not be executed`, name)
	})
	// Register your custom macro function which accepts a slice of strings [...,...].
	app.Macros().Get("string").RegisterFunc("has",
		func(validNames []string) func(string) bool {
			return func(paramValue string) bool {
				for _, validName := range validNames {
					if validName == paramValue {
						return true
					}
				}

				return false
			}
		})

	app.Get("/static_validation/{name:string has([kataras,maropoulos])}",
		func(ctx iris.Context) {
			name := ctx.Params().Get("name")
			ctx.Writef(`Hello %s | the name should be "kataras" or "maropoulos" otherwise this handler will not be executed`, name)
		})
}
