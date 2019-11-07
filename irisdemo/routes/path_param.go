package routes

import (
	"github.com/kataras/iris/v12"
	"regexp"
)

func registerPathParamRoute(app *iris.Application) {

	// Method:    GET
	// Resource:  http://localhost:8080/user/42
	//
	// Need to use a custom regexp instead?
	// Easy;
	// Just mark the parameter's type to 'string'
	// which accepts anything and make use of
	// its `regexp` macro function, i.e:
	// app.Get("/user/{id:string regexp(^[0-9]+$)}")
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	patterns := []string{
		"/",                          // Matches only GET "/".
		"/assets/{asset:path}",       // Matches all GET requests prefixed with "/assets/**/*"
		"/files/{name:file}",         // Matches all GET requests prefixed with "/files/*", lowercase or uppercase letters, numbers, underscore (_), dash (-), point (.) and no spaces or other special characters that are not valid for filenames
		"/profile/{username:string}", // Matches all GET requests prefixed with "/profile/" and followed by a single path part.
		"/profile/me",                // Matches only GET "/profile/me" and it does not conflict with /profile/{username:string} or any root wildcard /{root:path}.
		"/profile/{username:string prefix(test)}",        // Matches all GET requests prefixed with /profile/ and followed by a string which prefixed with test.
		"/profile/{username:string suffix(test)}",        // Matches all GET requests prefixed with /profile/ and followed by a string which suffixed with test.
		"/users/{userid:int min(1)}",                     // Matches all GET requests prefixed with /users/ and followed by a number which should be equal or higher than 1.
		"/users/{userid:int max(9999)}",                  // Matches all GET requests prefixed with /users/ and followed by a number which should be lowwer than 9999.
		"{root:path}",                                    // Matches all GET requests except the ones that are already handled by other routes.
		"/u/{username:string}",                           // Matches all GET requests of: /u/abcd123 maps to :string
		"/u/{id:int}",                                    // Matches all GET requests of: /u/-1 maps to :int (if :int registered otherwise :string)
		"/u/{uid:uint}",                                  // Matches all GET requests of: /u/42 maps to :uint (if :uint registered otherwise :int)
		"/u/{firstname:alphabetical}",                    // Matches all GET requests of: /u/abcd maps to :alphabetical (if :alphabetical registered otherwise :string)
		"/{alias:string regexp(^[a-z0-9]{1,10}\\.xml$)}", // Matches all GET requests of /abctenchars.xml respectfully.
		"/{alias:string regexp(^[a-z0-9]{1,10}$)}",       // Matches all GET requests of /abcdtenchars respectfully.

	}
	for _, p := range patterns {
		app.Get(p, createHandler(p))
	}

	latLonExpr := "^-?[0-9]{1,3}(?:\\.[0-9]{1,10})?$"
	latLonRegex, _ := regexp.Compile(latLonExpr)

	// Register your custom argument-less macro function to the :string param type.
	// MatchString is a type of func(string) bool, so we use it as it is.
	app.Macros().Get("string").RegisterFunc("coordinate", latLonRegex.MatchString)

	app.Get("/coordinates/{lat:string coordinate()}/{lon:string coordinate()}",
		func(ctx iris.Context) {
			ctx.Writef("Lat: %s | Lon: %s", ctx.Params().Get("lat"), ctx.Params().Get("lon"))
		})

	app.Macros().Get("string").RegisterFunc("range",
		func(minLength, maxLength int) func(string) bool {
			return func(paramValue string) bool {
				return len(paramValue) >= minLength && len(paramValue) <= maxLength
			}
		})

	app.Get("/limitchar/{name:string range(1,200) else 400}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | the name should be between 1 and 200 characters length
    otherwise this handler will not be executed`, name)
	})
}

func createHandler(pattern string) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		ctx.Writef("Request from method: %s and path: %s\n", ctx.Method(), ctx.Path())
		ctx.Writef("Request matches pattern: %s\n", pattern)
		ctx.Writef("Request params:")
		params := make(map[string]string)
		ctx.Params().Visit(func(key, value string) {
			params[key] = value
		})
		ctx.JSON(params)
	}
}
