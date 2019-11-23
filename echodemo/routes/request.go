package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func registerRequestRoute(e *echo.Echo) {

	// Handler
	e.POST("/user/create", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})
}
