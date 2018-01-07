package service

import (
	"testing"
	"github.com/cloudnativego/gogo-engine"
)

func TestAddMatchShowsUpInRepository(t *testing.T) {
	match := gogo.NewMatch(19, "playerBlack", "playerWhite")

	repo := newInMemoryRepository()
	err := repo.addMatch(match)
	if err != nil {
		t.Error("Got an error adding a match to repository, should not have.")
	}

	matches := repo.getMatches()
	if len(matches) != 1 {
		t.Errorf("Expected to have 1 match in the repository, got %d", len(matches))
	}

	if match.PlayerBlack != "playerBlack" {
		t.Errorf("Player 1's name should have been playerBlack, got %s", match.PlayerBlack)
	}

	if match.PlayerWhite != "playerWhite" {
		t.Errorf("Player 2's name should have been playerWhite, got %s", match.PlayerWhite)
	}
}

func TestNewRepositoryIsEmpty(t *testing.T) {
	repo := newInMemoryRepository()

	matches := repo.getMatches()
	if len(matches) != 0 {
		t.Errorf("Expected to have 0 matches in newly created repository, got %d", len(matches))
	}
}