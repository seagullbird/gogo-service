package service

import (
	"github.com/unrolled/render"
	"net/http"
	"github.com/cloudnativego/gogo-engine"
	"encoding/json"
	"io/ioutil"
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
		req_payload, _ := ioutil.ReadAll(req.Body)
		var newMatchRequest newMatchRequest
		json.Unmarshal(req_payload, &newMatchRequest)

		match := gogo.NewMatch(newMatchRequest.GridSize, newMatchRequest.PlayerBlack, newMatchRequest.PlayerWhite)
		repo.addMatch(match)
		w.Header().Add("Location", "/matches/" + match.ID)
		formatter.JSON(w,
			http.StatusCreated,
			&newMatchResponse{
				Id: match.ID,
				GridSize: match.GridSize,
				PlayerWhite: match.PlayerWhite,
				PlayerBlack: match.PlayerBlack,
			})
	}
}