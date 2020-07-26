package server

import (
	"context"
	mydb "goapi/db"
	"log"
	"net/http"
)

var (
	ctx context.Context = context.Background()
)

type serverContext struct {
	w   http.ResponseWriter
	r   *http.Request
	db  mydb.IDB
	ctx context.Context
}

type handlerContext func(ctx *serverContext) error

type route struct {
	method string
	path   string
	h      handlerContext
}

type server struct {
	routes []route
}

func (s *server) POST(path string, h handlerContext) {
	newRoute := route{http.MethodPost, path, h}
	s.routes = append(s.routes, newRoute)
}

func (s *server) GET(path string, h handlerContext) {
	newRoute := route{http.MethodGet, path, h}
	s.routes = append(s.routes, newRoute)
}

func (s *server) PUT(path string, h handlerContext) {
	newRoute := route{http.MethodPut, path, h}
	s.routes = append(s.routes, newRoute)
}

func (s *server) handleHTTPMethod(currentMethod string, ctx serverContext) {
	switch currentMethod {
	case http.MethodGet:
		log.Println("method GET!!!")
		return
	case http.MethodPost:
		log.Println("method POST!!!")
		return
	case http.MethodPut:
		log.Println("mehod PUT")
	default:
		log.Println("Default case")
	}
}

func (s *server) handlerServer(db mydb.IDB) http.HandlerFunc {
	backgroundCtx := context.Background()
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := serverContext{w, r, db, backgroundCtx}
		for i := 0; i < len(s.routes); i++ {
			currentRoute := s.routes[i]
			isCorrectMethod := currentRoute.method == r.Method
			isCorrectPath := currentRoute.path == r.URL.Path
			if isCorrectMethod && isCorrectPath {
				err := currentRoute.h(&ctx)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func (s *server) Run(port string) error {
	url := ":" + port
	db, err := mydb.StartDatabase()
	log.Printf("listening on port http://localhost%s \n", url)
	err = http.ListenAndServe(url, s.handlerServer(&db))
	if err != nil {
		return err
	}
	return nil
}
