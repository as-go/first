package main

import "github.com/as-go/first/internal/app"

func main() {
	// config := config.New()
	app := app.New(app.Config{
		Address:          "127.0.0.1:8080",
		ConnectionString: "",
	})

	app.Start()
}
