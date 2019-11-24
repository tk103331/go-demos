package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type CustomBinder struct {}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// You may use default binder
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}

	// Define your custom implementation

	return
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}


func registerRequestRoute(e *echo.Echo) {
	// Custom Binder
	e.Binder = &CustomBinder{}

	// Validate Data
	e.Validator = &CustomValidator{validator: validator.New()}

	// Bind Data
	// To bind request body into a Go type use Context#Bind(i interface{}). The default binder supports decoding application/json, application/xml and application/x-www-form-urlencoded data based on the Content-Type header.
	e.POST("/user/create", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})

	// Form Data
	e.POST("/user/create2", func(c echo.Context) error {
		name := c.FormValue("name")
		return c.String(http.StatusOK, name)
	})

	// Query Parameters
	e.GET("/user/byname", func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, name)
	})

	// Path Parameters
	e.GET("/user/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, name)
	})
}
