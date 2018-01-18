package api

import (
	"gopkg.in/macaron.v1"
	"net/http"
)


func TestAPI(ctx *macaron.Context) {
	ctx.WriteHeader(http.StatusOK)
	ctx.Write([]byte("OK"))
}
