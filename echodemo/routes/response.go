package routes

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func registerResponseRoute(e *echo.Echo) {
	// Context#String(code int, s string) can be used to send plain text response with status code.
	e.GET("/response/string", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// Context#HTML(code int, html string) can be used to send simple HTML response with status code.
	// If you are looking to send dynamically generate HTML see templates.
	e.GET("/response/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	})
	// Context#HTMLBlob(code int, b []byte) can be used to send HTML blob with status code.
	// You may find it handy using with a template engine which outputs []byte.
	e.GET("/response/htmlblob", func(c echo.Context) error {
		b := []byte("<strong>Hello, World!</strong>")
		return c.HTMLBlob(http.StatusOK, b)
	})
	// Context#JSON(code int, i interface{}) can be used to encode a provided Go type into JSON and send it as response with status code.
	e.GET("/response/json", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.JSON(http.StatusOK, u)
	})
	// Context#JSON() internally uses json.Marshal which may not be efficient to large JSON, in that case you can directly stream JSON.
	e.GET("/response/stream", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(u)
	})
	// Context#JSONPretty(code int, i interface{}, indent string) can be used to a send a JSON response which is pretty printed based on indent, which could be spaces or tabs.
	e.GET("/response/jsonpretty", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "joe@labstack.com",
		}
		return c.JSONPretty(http.StatusOK, u, "  ")
	})
	// Context#JSONBlob(code int, b []byte) can be used to send pre-encoded JSON blob directly from external source, for example, database.
	e.GET("/response/jsonblob", func(c echo.Context) error {
		encodedJSON := []byte(`{"name":"Jon","email":"jon@labstack.com"}`) // Encoded JSON from external source
		return c.JSONBlob(http.StatusOK, encodedJSON)
	})
	// Context#XML(code int, i interface{}) can be used to encode a provided Go type into XML and send it as response with status code.
	e.GET("/response/xml", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.XML(http.StatusOK, u)
	})
}
