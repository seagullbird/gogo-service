package service

import (
	"github.com/unrolled/render"
	"net/http"
	"github.com/satori/go.uuid"
)

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Test string
		}{"This is a test"})
	}
}

func createMatchHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		guid := uuid.NewV4().String()
		w.Header().Add("Location", "/matches/" + guid)
		formatter.JSON(w,
			http.StatusCreated,
			&newMatchResponse{
				Id: guid,
			})
	}
}