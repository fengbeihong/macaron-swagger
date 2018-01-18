package main

import (
	"github.com/fengbeihong/macaron-swagger"
	"github.com/fengbeihong/macaron-swagger/example/api"
	"github.com/fengbeihong/macaron-swagger/swaggerFiles"
	_ "github.com/fengbeihong/macaron-swagger/example/docs"

	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Get("/testapi", api.TestAPI)
	m.Get("/swagger/*", macaronSwagger.WrapHandler(swaggerFiles.Handler))
	m.Run()
}