package routes

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"gopkg.in/go-playground/validator.v9"
)

// User contains user information.
type User struct {
	FirstName      string     `json:"fname"`
	LastName       string     `json:"lname"`
	Age            uint8      `json:"age" validate:"gte=18,lte=45"`
	Email          string     `json:"email" validate:"required,email"`
	FavouriteColor string     `json:"favColor" validate:"hexcolor|rgb|rgba"`
	Addresses      []*Address `json:"addresses" validate:"required,dive,required"`
}

// Address houses a users address information.
type Address struct {
	Street string `json:"street" validate:"required"`
	City   string `json:"city" validate:"required"`
	Planet string `json:"planet" validate:"required"`
	Phone  string `json:"phone" validate:"required"`
}

// Use a single instance of Validate, it caches struct info.
var validate *validator.Validate

func registerValidationRoute(app *iris.Application) {
	validate = validator.New()

	// Register validation for 'User'
	// NOTE: only have to register a non-pointer type for 'User', validator
	// internally dereferences during it's type checks.
	validate.RegisterStructValidation(userStructLevelValidation, User{})

	app.Post("/validation", func(ctx iris.Context) {
		var user User
		if err := ctx.ReadJSON(&user); err != nil {
			// [handle error...]
		}
		fmt.Println(user)
		// Returns InvalidValidationError for bad validation input,
		// nil or ValidationErrors ( []FieldError )
		err := validate.Struct(user)
		if err != nil {

			// This check is only needed when your code could produce
			// an invalid value for validation such as interface with nil
			// value most including myself do not usually have code like this.
			if _, ok := err.(*validator.InvalidValidationError); ok {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.WriteString(err.Error())
				return
			}

			ctx.StatusCode(iris.StatusBadRequest)
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println()
				fmt.Print(err.Namespace())
				fmt.Print(" ")
				fmt.Print(err.Field())
				fmt.Print(" ")
				fmt.Print(err.StructNamespace())
				fmt.Print(" ")
				fmt.Print(err.StructField())
				fmt.Print(" ")
				fmt.Print(err.Tag())
				fmt.Print(" ")
				fmt.Print(err.ActualTag())
				fmt.Print(" ")
				fmt.Print(err.Kind())
				fmt.Print(" ")
				fmt.Print(err.Type())
				fmt.Print(" ")
				fmt.Print(err.Value())
				fmt.Print(" ")
				fmt.Print(err.Param())
				fmt.Println()
			}

			return
		}

		// [save user to database...]
	})

	// Example request of JSON form:
	// {
	// 	"fname": "",
	// 	"lname": "",
	// 	"age": 45,
	// 	"email": "mail@example.com",
	// 	"favColor": "#000",
	// 	"addresses": [{
	// 		"street": "Eavesdown Docks",
	// 		"planet": "Persphone",
	// 		"phone": "none",
	// 		"city": "Unknown"
	// 	}]
	// }
}

func userStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}
}
