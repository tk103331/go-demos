package routes

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func registerErrorHandingRoute(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the credentials from HTTP request header and perform a security
			// check
			path := c.Path()
			c.Logger().Info(path)
			// Request path start with '/admin' require credentials, read token from cookie
			if path == "/admin" || path[0:7] == "/admin/" {
				cookie, err := c.Cookie("token")
				if err != nil || !checkAdminToken(cookie.Value) {
					// For invalid credentials
					return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
				}

			}
			// For valid credentials call next
			return next(c)
		}
	})
	// HTTP Error Handler
	e.HTTPErrorHandler = customHTTPErrorHandler
}

// Custom Error Pages
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("views/errors/%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

func checkAdminToken(token string) bool {
	return token == "admin-token"
}
