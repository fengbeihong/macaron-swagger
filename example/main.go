// Package main Test API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// http://swagger.io/terms/
//
//     Schemes: http
//     Host: 127.0.0.1:4000
//     Version: 1.0
//     License: Apache 2.0 http://www.apache.org/licenses/LICENSE-2.0.html
//     Contact: API Support<support@swagger.io> http://www.swagger.io/support
//
// swagger:meta
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