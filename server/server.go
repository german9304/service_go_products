package server

import (
	"context"
	"encoding/json"
	mydb "goapi/db"
	"log"
	"net/http"
)

func isValidMethod(currentRoute route, req *http.Request) bool {
	isCorrectMethod := currentRoute.method == req.Method
	isCorrectPath := currentRoute.path == req.URL.Path
	return isCorrectMethod && isCorrectPath
}

type ServerContext struct {
	W   http.ResponseWriter
	R   *http.Request
	DB  mydb.IDB
	Ctx context.Context
}

// JSON makes an HTTP response with content-type of json
func (sctx *ServerContext) JSON(data interface{}) error {
	sctx.W.Header().Set("Content-type", "application/json")
	b, err := json.Marshal(data)
	_, err = sctx.W.Write(b)
	if err != nil {
		return err
	}
	return nil
}

type handlerContext func(ctx *ServerContext) error

type route struct {
	method string
	path   string
	h      handlerContext
}

type Server struct {
	routes []route
}

func (s *Server) POST(path string, h handlerContext) {
	newRoute := route{http.MethodPost, path, h}
	s.routes = append(s.routes, newRoute)
}

func (s *Server) GET(path string, h handlerContext) {
	newRoute := route{http.MethodGet, path, h}
	s.routes = append(s.routes, newRoute)
}

func (s *Server) PUT(path string, h handlerContext) {
	newRoute := route{http.MethodPut, path, h}
	s.routes = append(s.routes, newRoute)
}

func (s *Server) handlerServer(db mydb.IDB, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := ServerContext{w, r, db, ctx}
		for i := 0; i < len(s.routes); i++ {
			currentRoute := s.routes[i]
			if isValidMethod(currentRoute, r) {
				err := currentRoute.h(&ctx)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func (s *Server) Run(port string) error {
	ctx := context.Background()
	url := ":" + port
	db, err := mydb.StartDatabase(ctx)
	if err != nil {
		return err
	}
	log.Printf("listening on port http://localhost%s \n", url)
	err = http.ListenAndServe(url, s.handlerServer(&db, ctx))
	if err != nil {
		return err
	}
	return nil
}
