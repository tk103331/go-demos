package routes

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/v12"
)

var (
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey  = []byte("the-big-and-secret-fash-key-here")
	blockKey = []byte("lot-secret-of-characters-big-too")
	sc       = securecookie.New(hashKey, blockKey)
)

func registerCookieRoute(app *iris.Application) {
	log := app.Logger()
	app.Get("/getcookie", func(ctx iris.Context) {
		value1 := ctx.GetCookie("cookie1")
		log.Info(value1)

		cookie2, err := ctx.Request().Cookie("cookie2")
		if err != nil {
			log.Info(cookie2)
		}

		value6 := ctx.GetCookie("cookie6", iris.CookieDecode(sc.Decode))
		log.Info(value6)

		cookies := ctx.Request().Cookies()
		for _, cookie := range cookies {
			log.Infof("%s=%s", cookie.Name, cookie.Value)
		}

		ctx.WriteString("get cookie")
	})
	app.Get("/setcookie", func(ctx iris.Context) {

		ctx.SetCookie(&http.Cookie{Name: "cookie1", Value: "value1"})

		ctx.SetCookie(&http.Cookie{
			Name:     "cookie2",
			Value:    "value2",
			Path:     "/",
			MaxAge:   30,
			HttpOnly: true,
			Domain:   "localhost",
		})

		ctx.SetCookieKV("cookie3", "vlaue3")

		ctx.SetCookieKV("cookie4", "vlaue4", iris.CookiePath("/"), iris.CookieExpires(30*time.Second), iris.CookieHTTPOnly(true))

		ctx.SetCookieKV("cookie5", "value5", func(cookie *http.Cookie) {
			cookie.Path = "/"
			cookie.Domain = "localhost"
			cookie.MaxAge = 30
			cookie.HttpOnly = true
		})

		ctx.SetCookieKV("cookie6", "value6", iris.CookieEncode(sc.Encode))
		ctx.WriteString("set cookie")
	})
}
