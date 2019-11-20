package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func registerCookieRoute(e *echo.Echo) {

	e.GET("/setcookie", writeCookie)
	e.GET("/getcookie", readCookie)
	e.GET("/getallcookies", readAllCookies)
}

// Create a Cookie
func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

// Read a Cookie
func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	c.Logger().Info(cookie.Name)
	c.Logger().Info(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}

// Read all the Cookies
func readAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		c.Logger().Infof("%s = %s", cookie.Name, cookie.Value)
	}
	return c.String(http.StatusOK, "read all the cookies")
}
