package service

import (
	"github.com/unrolled/render"
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/cloudnativego/gogo-engine"
)

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Test string
		}{"This is a test"})
	}
}

func createMatchHandler(formatter *render.Render, repo *inMemoryMatchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		match := gogo.NewMatch(19, "playerBlack", "playerWhite")
		repo.addMatch(match)
		guid := uuid.NewV4().String()
		w.Header().Add("Location", "/matches/" + guid)
		formatter.JSON(w,
			http.StatusCreated,
			&newMatchResponse{
				Id: guid,
			})
	}
}