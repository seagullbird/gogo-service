package service

import (
	"testing"
	"github.com/unrolled/render"
	"net/http"
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"encoding/json"
)

const (
	fakeMatchLocationResult = "/matches/5a003b78-409e-4452-b456-a6f0dcee05bd"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

func TestCreateMatch(t *testing.T) {
	client := &http.Client{}
	repo := newInMemoryRepository()
	server := httptest.NewServer(http.HandlerFunc(createMatchHandler(formatter, repo)))
	defer server.Close()

	body := []byte("{\n  \"gridsize\": 19,\n  \"players\": [\n    {\n      \"color\": \"white\",\n      \"name\": \"bob\"\n    },\n    {\n      \"color\": \"black\",\n      \"name\": \"alfred\"\n    }\n  ]\n}")

	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to createMatchHandler: %v", err)
		return
	}
	defer resp.Body.Close()
	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Errored reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected response status 201, received %s", resp.Status)
	}

	loc, headerOk := resp.Header["Location"]
	if !headerOk {
		t.Error("Location header is not set")
	} else {
		if !strings.Contains(loc[0], "/matches/") {
			t.Errorf("Location header should contain '/matches/', current: %s", loc[0])
		}

		if len(loc[0]) != len(fakeMatchLocationResult) {
			t.Errorf("Location value does not contain guid of new match: %s", loc[0])
		}
	}

	var matchResponse newMatchResponse
	err = json.Unmarshal(payload, &matchResponse)
	if err != nil {
		t.Errorf("Could not unmarshal payload into newMatchResponse object")
	}
	if matchResponse.Id == "" || !strings.Contains(loc[0], matchResponse.Id) {
		t.Error("matchResponse.Id does not match Location header")
	}

	matches := repo.getMatches()
	if len(matches) != 1 {
		t.Errorf("Expected a match repo of 1 match, got size %d", len(matches))
	}
}
