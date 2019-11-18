package routes

import (
	"github.com/labstack/echo"
)

// RegisterRoutes register all routes
func RegisterRoutes(e *echo.Echo) {
	registerOverviewRoute(e)
}
