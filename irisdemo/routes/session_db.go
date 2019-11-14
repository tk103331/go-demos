package routes

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/badger"
	//for redis
)

func registerSessionDBRoute(app *iris.Application) {
	// 2. Initialize the database.
	// These are the default values,
	// you can replace them based on your running redis' server settings:

	// for redis
	// db := redis.New(redis.Config{
	// 	Network:   "tcp",
	// 	Addr:      "127.0.0.1:6379",
	// 	Timeout:   time.Duration(30) * time.Second,
	// 	MaxActive: 10,
	// 	Password:  "",
	// 	Database:  "",
	// 	Prefix:    "",
	// 	Delim:     "-",
	// 	Driver:    redis.Redigo(), // redis.Radix() can be used instead.
	// })

	// for boltdb
	// db, err := boltdb.New("./sessions.db", os.FileMode(0750))

	// badger
	db, err := badger.New("./data")
	if err != nil {
		panic(err)
	}
	// Optionally configure the underline driver:
	// driver := redis.Redigo()
	// driver.MaxIdle = ...
	// driver.IdleTimeout = ...
	// driver.Wait = ...
	// redis.Config {Driver: driver}

	// Close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	sess := sessions.New(sessions.Config{
		Cookie:       "sessionscookieid",
		AllowReclaim: true,
	})

	// 3. Register it.
	sess.UseDatabase(db)

	app.Get("session_db", func(ctx iris.Context) {
		session := sess.Start(ctx)

		logintime := session.GetString("logintime")
		if logintime == "" {
			logintime = time.Now().String()
			session.Set("logintime", logintime)
		}
		ctx.WriteString(logintime)
	})
}
