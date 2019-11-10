package routes

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/kataras/iris/v12"
)

const maxSize = 5 << 20 // 5MB

func registerFormDataRoute(app *iris.Application) {
	// Urlencoded Form
	app.Post("/form_post", func(ctx iris.Context) {
		message := ctx.FormValue("message")
		nick := ctx.FormValueDefault("nick", "anonymous")

		ctx.JSON(iris.Map{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	// Another example: query + post form
	app.Post("/post", func(ctx iris.Context) {
		id := ctx.URLParam("id")
		page := ctx.URLParamDefault("page", "0")
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		// or `ctx.PostValue` for POST, PUT & PATCH-only HTTP Methods.

		msg := fmt.Sprintf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		ctx.WriteString(msg)
		app.Logger().Infof(msg)
	})
	// Upload files
	app.Post("/upload", iris.LimitRequestBodySize(maxSize), func(ctx iris.Context) {
		//
		// UploadFormFiles
		// uploads any number of incoming files ("multiple" property on the form input).
		//

		// The second, optional, argument
		// can be used to change a file's name based on the request,
		// at this example we will showcase how to use it
		// by prefixing the uploaded file with the current user's ip.
		_, err := ctx.UploadFormFiles("./uploads", beforeSave)
		if err != nil {
			ctx.Application().Logger().Error(err)
		}

	})

	app.Post("/upload_manual", func(ctx iris.Context) {
		// Get the max post value size passed via iris.WithPostMaxMemory.
		maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

		err := ctx.Request().ParseMultipartForm(maxSize)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
			return
		}

		form := ctx.Request().MultipartForm

		files := form.File["files[]"]
		failures := 0
		for _, file := range files {
			_, err = saveUploadedFile(file, "./uploads")
			if err != nil {
				failures++
				ctx.Writef("failed to upload: %s\n", file.Filename)
				ctx.Writef("error: %s\n", err.Error())
			}
		}
		ctx.Writef("%d files uploaded", len(files)-failures)
	})
}

func saveUploadedFile(fh *multipart.FileHeader, destDirectory string) (int64, error) {
	e := os.MkdirAll(destDirectory, os.ModeDir|os.ModePerm)
	if e != nil {
		return 0, e
	}
	fi, err := os.Lstat(destDirectory)
	if err == os.ErrNotExist {
		e := os.MkdirAll(destDirectory, os.ModeDir)
		if e != nil {
			return 0, e
		}
	} else if err != nil {
		return 0, err
	}
	if !fi.IsDir() {
		return 0, errors.New("Destination Directory is not dir: " + destDirectory)
	}

	src, err := fh.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()

	out, err := os.OpenFile(filepath.Join(destDirectory, fh.Filename),
		os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
	if err != nil {
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, src)
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	ctx.Application().Logger().Info(ip)
	// make sure you format the ip in a way
	// that can be used for a file name (simple case):
	ip = strings.Replace(ip, ".", "_", -1)
	ip = strings.Replace(ip, ":", "_", -1)

	// you can use the time.Now, to prefix or suffix the files
	// based on the current time as well, as an exercise.
	// i.e unixTime :=    time.Now().Unix()
	// prefix the Filename with the $IP-
	// no need for more actions, internal uploader will use this
	// name to save the file into the "./uploads" folder.
	file.Filename = ip + "-" + file.Filename
}
