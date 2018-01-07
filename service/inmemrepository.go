package service

import "github.com/cloudnativego/gogo-engine"

type inMemoryMatchRepository struct {
	matches []gogo.Match
}

func newInMemoryRepository() *inMemoryMatchRepository {
	repo := &inMemoryMatchRepository{}
	repo.matches = []gogo.Match{}
	return repo
}

func (repo *inMemoryMatchRepository) addMatch(match gogo.Match) (err error) {
	repo.matches = append(repo.matches, match)
	return
}

func (repo *inMemoryMatchRepository) getMatches() []gogo.Match {
	return repo.matches
}