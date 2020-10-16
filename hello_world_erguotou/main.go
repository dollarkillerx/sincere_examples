package main

import "github.com/dollarkillerx/erguotou"

func main() {
	app := erguotou.New()

	app.Get("/hello", func(ctx *erguotou.Context) {
		ctx.String(200, "Hello World")
	})
	
	app.Run(erguotou.SetHost("0.0.0.0:8081"))
}
