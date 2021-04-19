package server

import (
	"log"
	"net/http"
)

// Handler handler to authenticate user
func AuthHandler(h HandlerContext) HandlerContext {
	return func(ctx *ServerContext) error {
		username, password, ok := ctx.R.BasicAuth()
		if ok {
			log.Printf("user: %s is authenticated with %s \n", username, password)
			h(ctx)
		}
		ctx.W.Header().Add("WWW-Authenticate", `Basic realm="prected api"`)
		http.Error(ctx.W, "credentials needed", http.StatusUnauthorized)
		return nil
	}
}
