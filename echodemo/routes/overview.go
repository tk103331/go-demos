package routes

import (
	"github.com/labstack/echo/middleware"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func registerOverviewRoute(e *echo.Echo) {

	// Routing
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", saveAvatar)
	e.DELETE("/users/:id", func(c echo.Context) error {
		c.Logger().Info(c.Path())
		return nil
	})

	// Handling Request
	e.GET("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		u.Name = "Jack"
		u.Email = "jack@echo.com"
		return c.JSON(http.StatusCreated, []User{*u})
		// or
		// return c.XML(http.StatusCreated, u)
	})

	// Static Content
	e.Static("/static", "static")

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

// Path Parameters
// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// Query Parameters
//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// Form application/x-www-form-urlencoded
// e.POST("/save", save)
func saveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

// Form multipart/form-data
func saveAvatar(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
