package main

import (
	"sort"
	"testing"
	"time"
)

func TestGame(t *testing.T) {
	game := Game{ID: "1",
		Home:       "Mexico",
		Away:       "Canada",
		ScoreHome:  0,
		ScoreAway:  5,
		TotalScore: 5,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now()}

	if game.TotalScore != 5 {
		t.Errorf("TotalScore should be 5")
	}

	if game.ScoreHome != 0 {
		t.Errorf("ScoreHome should be 0")
	}

	if game.ScoreAway != 5 {
		t.Errorf("ScoreAway should be 5")
	}
}

func TestGameSorter(t *testing.T) {
	game1 := Game{ID: "1", Home: "Mexico", Away: "Canada", ScoreHome: 0, ScoreAway: 5, TotalScore: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	game2 := Game{ID: "2", Home: "Spain", Away: "Brazil", ScoreHome: 10, ScoreAway: 2, TotalScore: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	game3 := Game{ID: "3", Home: "Germany", Away: "France", ScoreHome: 2, ScoreAway: 2, TotalScore: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	game4 := Game{ID: "4", Home: "Uruguay", Away: "Italy", ScoreHome: 6, ScoreAway: 6, TotalScore: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	game5 := Game{ID: "5", Home: "Argentina", Away: "Australia", ScoreHome: 3, ScoreAway: 1, TotalScore: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	games := []*Game{&game1, &game2, &game3, &game4, &game5}

	sort.Sort(GameSorter(games))

	result := []*Game{&game4, &game2, &game1, &game5, &game3}

	if !compareGames(games, result) {
		t.Errorf("Games should be sorted by score and the same total score will be returned ordered by the most recently added")
	}

	for i, game := range games {
		if game.TotalScore != result[i].TotalScore {
			t.Errorf("TotalScore expected: %d, got: %d", result[i].TotalScore, game.TotalScore)
		}
	}
}

func compareGames(a, b []*Game) bool {
	if len(a) != len(b) {
		return false
	}

	for i, game := range a {
		if game.TotalScore != b[i].TotalScore && game.Home != b[i].Home && game.Away != b[i].Away {
			return false
		}
	}

	return true
}