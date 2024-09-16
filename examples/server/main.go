package main

import (
	api "github.com/templatedop/api/apis"
	"github.com/templatedop/api/examples/server/core/port"
	h "github.com/templatedop/api/examples/server/handler"
	r "github.com/templatedop/api/examples/server/repo"
	v "github.com/templatedop/api/examples/server/validation"
)

func main() {
	app := api.New()

	app.Use(
		port.Module(),
		h.Handlermodule,
		r.Repomodule,
		v.ValidatorModule,
	)
	/*Instead of module you can use app.Provide() to register the dependencies*/
	// app.Provide(
	// 	h.NewHelloHandler,
	// )

	err := app.Start()
	if err != nil {
		panic(err)
	}
}
