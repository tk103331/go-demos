package main

import (
	"github.com/labstack/echo"
	"github.com/tk103331/go-demos/echodemo/routes"
)

func main() {
	e := echo.New()

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
