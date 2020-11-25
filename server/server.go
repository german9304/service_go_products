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
	DB  mydb.Database
	Ctx context.Context
}

// JSON makes an HTTP response with content-type of json
func (ctx *ServerContext) JSON(data interface{}) error {
	ctx.W.Header().Set("Content-type", "application/json")
	b, err := json.Marshal(data)
	_, err = ctx.W.Write(b)
	if err != nil {
		return err
	}
	return nil
}

type HandlerContext func(ctx *ServerContext) error

type route struct {
	method string
	path   string
	h      HandlerContext
}

// Server defines the server type
type Server struct {
	routes []route
}

// POST route for HTTP POST request
func (s *Server) POST(path string, h HandlerContext) {
	newRoute := route{http.MethodPost, path, h}
	s.routes = append(s.routes, newRoute)
}

// GET route for HTTP GET request
func (s *Server) GET(path string, h HandlerContext) {
	newRoute := route{http.MethodGet, path, h}
	s.routes = append(s.routes, newRoute)
}

// PUT route for HTTP PUT request
func (s *Server) PUT(path string, h HandlerContext) {
	newRoute := route{http.MethodPut, path, h}
	s.routes = append(s.routes, newRoute)
}

func (s *Server) handlerServer(ctx context.Context, db mydb.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := ServerContext{w, r, db, ctx}
		for i := 0; i < len(s.routes); i++ {
			currentRoute := s.routes[i]
			if isValidMethod(currentRoute, r) {
				log.Println("method allowed", currentRoute)
				err := currentRoute.h(&ctx)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
		}
		log.Println("method not allowed")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (s *Server) Run(port string) error {
	ctx := context.Background()
	url := ":" + port
	db, err := mydb.Start(ctx)
	if err != nil {
		return err
	}
	log.Printf("listening on port http://localhost%s \n", url)
	err = http.ListenAndServe(url, s.handlerServer(ctx, &db))
	if err != nil {
		return err
	}
	return nil
}
