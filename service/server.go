package service

import (
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()

	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/test", testHandler(formatter)).Methods("GET")
	repo := &inMemoryMatchRepository{}
	mx.HandleFunc("/matches", createMatchHandler(formatter, repo)).Methods("POST")
}
