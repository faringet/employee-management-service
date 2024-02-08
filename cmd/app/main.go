package main

import (
	"fmt"

	"github.com/engagerocketco/templates-api-svc/internal/app"
)

//	@title						Template API
//	@description				Template API service
//	@BasePath					/api/v1/template
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
func main() {
	app, err := app.NewApp()
	panicOnErr(err)
	app.Run()
}

func panicOnErr(err error) {
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}
