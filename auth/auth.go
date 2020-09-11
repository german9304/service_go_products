package auth

import (
	"goapi/server"
	"log"
	"net/http"
)

// Handler handler to authenticate user
func Handler(h server.HandlerContext) server.HandlerContext {
	return func(ctx *server.ServerContext) error {
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
