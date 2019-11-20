package routes

import "github.com/labstack/echo"

// Define a custom context
type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	c.Context.Logger().Info("foo")
}

func (c *CustomContext) Bar() {
	c.Context.Logger().Info("bar")
}
func registerContextRoute(e *echo.Echo) {

	// Create a middleware to extend default context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			c.Logger().Info("Create a middleware to extend default context")
			return next(cc)
		}
	})

	//Use in handler
	e.GET("/context", func(c echo.Context) error {
		cc := c.(*CustomContext)
		cc.Foo()
		cc.Bar()
		return cc.String(200, "OK")
	})
}
