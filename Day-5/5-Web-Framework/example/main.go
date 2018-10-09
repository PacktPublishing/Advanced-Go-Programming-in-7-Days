package main

import (
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-5/5-Web-Framework"
	"fmt"
)

func main() {
	app := simplex.New()
	app.Use(func(ctx *simplex.Context) {
		ctx.AddHeader("X-Info", "Hello")
	})
	app.Get("/", func(ctx *simplex.Context){
		ctx.Send("Hello World")
	})
	app.Post("/add/user", func(ctx *simplex.Context){
		name, _ := ctx.Query("name")
		if name == "" {
			ctx.Send("What's your name again?")
		} else {
			ctx.Send(fmt.Sprintf("Got Username: %s", name))
		}
	})
	app.Run()
}