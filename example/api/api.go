package api

import (
	"gopkg.in/macaron.v1"
	"net/http"
)


// swagger:operation GET /testapi api TestAPI
//
// test api for swagger 2.0
//
// ---
// produces:
// - application/json
// - text/html
// parameters:
// - name: param
//   in: path
//   description: param
//   required: false
//   type: string
// responses:
//   '200':
//     description: login response
//     schema:
//       type: string
func TestAPI(ctx *macaron.Context) {
	ctx.WriteHeader(http.StatusOK)
	ctx.Write([]byte("OK"))
}
