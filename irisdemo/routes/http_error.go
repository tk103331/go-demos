package routes

import (
	"github.com/kataras/iris/v12"
)

func registerHttpErrorRoute(app *iris.Application) {
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	// to register a handler for all "error"
	// status codes(kataras/iris/context.StatusCodeNotSuccessful)
	// defaults to < 200 || >= 400:
	// app.OnAnyErrorCode(handler)
	app.Get("/error", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Get("/problem", fireProblem)
}

func notFound(ctx iris.Context) {
	// when 404 then render the template
	// $views_dir/errors/404.html
	ctx.View("errors/404.html")
}

func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}

func newProductProblem(productName, detail string) iris.Problem {
	return iris.NewProblem().
		// The type URI, if relative it automatically convert to absolute.
		Type("/product-error").
		// The title, if empty then it gets it from the status code.
		Title("Product validation problem").
		// Any optional details.
		Detail(detail).
		// The status error code, required.
		Status(iris.StatusBadRequest).
		// Any custom key-value pair.
		Key("productName", productName)
	// Optional cause of the problem, chain of Problems.
	// .Cause(other iris.Problem)
}

func fireProblem(ctx iris.Context) {
	// Response like JSON but with indent of "  " and
	// content type of "application/problem+json"
	ctx.Problem(newProductProblem("product name", "problem details"),
		iris.ProblemOptions{
			// Optional JSON renderer settings.
			JSON: iris.JSON{
				Indent: "  ",
			},
			// OR
			// Render as XML:
			// RenderXML: true,
			// XML:       iris.XML{Indent: "  "},
			// Sets the "Retry-After" response header.
			//
			// Can accept:
			// time.Time for HTTP-Date,
			// time.Duration, int64, float64, int for seconds
			// or string for date or duration.
			// Examples:
			// time.Now().Add(5 * time.Minute),
			// 300 * time.Second,
			// "5m",
			//
			RetryAfter: 300,
			// A function that, if specified, can dynamically set
			// retry-after based on the request.
			// Useful for ProblemOptions reusability.
			// Overrides the RetryAfter field.
			//
			// RetryAfterFunc: func(iris.Context) interface{} { [...] }
		})
}
