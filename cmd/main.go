package main

import application "crud-go/internal/app"

func main() {
	app, _ := application.Initialize()

	app.Run()
}
