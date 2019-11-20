package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tk103331/go-demos/echodemo/routes"
	"net/http"
	"os"
	"time"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	routes.RegisterRoutes(e)

	if l, ok := e.Logger.(*log.Logger); ok {
		// Echo#Logger.SetHeader(io.Writer) can be used to set the header for the logger.
		// Default value: {"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}
		l.SetHeader("${time_rfc3339} ${level} ${short_file} ${line} ${prefix}")
		// Echo#Logger.SetOutput(io.Writer) can be used to set the output destination for the logger.
		// Default value is os.Stdout
		l.SetOutput(os.Stdout)
		// Echo#Logger.SetLevel(log.Lvl) can be used to set the log level for the logger.
		// Default value is ERROR
		l.SetLevel(log.ERROR)
	}

	// Echo#HideBanner can be used to hide the startup banner.
	e.HideBanner = true
	// Echo#DisableHTTP2 can be used disable HTTP/2 protocol.
	e.DisableHTTP2 = true
	// Echo#Validator can be used to register a validator for performing data validation on request payload.
	// Echo#Binder can be used to register a custom binder for binding request payload.
	// Echo#Renderer can be used to register a renderer for template rendering.
	// Echo#HTTPErrorHandler can be used to register a custom http error handler.

	// Echo#StartServer() can be used to run a custom server.
	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Minute, // Echo#*Server#ReadTimeout can be used to set the maximum duration before timing out read of the request.
		WriteTimeout: 20 * time.Minute, // Echo#*Server#WriteTimeout can be used to set the maximum duration before timing out write of the response.
	}
	e.Logger.Fatal(e.StartServer(s))
}
