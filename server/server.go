package server

import (
	"encoding/json"
	"context"
	mydb "goapi/db"
	"log"
	"net/http"
)

var (
	ctx context.Context = context.Background()
)

type ServerContext struct {
	W  http.ResponseWriter
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

func (s *Server) handleHTTPMethod(currentMethod string, ctx ServerContext) {
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

func (s *Server) handlerServer(db mydb.IDB) http.HandlerFunc {
	backgroundCtx := context.Background()
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := ServerContext{w, r, db, backgroundCtx}
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

func (s *Server) Run(port string) error {
	url := ":" + port
	db, err := mydb.StartDatabase()
	log.Printf("listening on port http://localhost%s \n", url)
	err = http.ListenAndServe(url, s.handlerServer(&db))
	if err != nil {
		return err
	}
	return nil
}
