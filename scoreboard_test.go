package main

import "testing"

func TestAdd(t *testing.T) {
	board := InitScoreBoard()
	err := board.Add("Mexico", "Canada")

	if err != nil {
		t.Errorf("Add should not return an error")
	}
}

func TestRemove(t *testing.T) {
	board := InitScoreBoard()
	board.Add("Mexico", "Canada")
	board.Remove("Mexico", "Canada")
	key := board.generateGameId("Mexico", "Canada")

	if _, found := board.games[key]; found {
		t.Errorf("Remove should remove the game from the board")
	}
}

func TestUpdate(t *testing.T) {
	board := InitScoreBoard()
	board.Add("Mexico", "Canada")
	err := board.Update("Mexico", "Canada", 0, 5)

	if err != nil {
		t.Errorf("Update should not return an error")
	}

	key := board.generateGameId("Mexico", "Canada")

	if board.games[key].ScoreHome != 0 {
		t.Errorf("Update should update the score home got %d, expected %d", board.games[key].ScoreHome, 0)
	}

	if board.games[key].ScoreAway != 5 {
		t.Errorf("Update should update the score away got %d, expected %d", board.games[key].ScoreAway, 5)
	}

	err = board.Update("MexicoNotExist", "Canada", 10, 2)

	if err == nil {
		t.Errorf("Update should return an error")
	}

	err = board.Update("Mexico", "Canada", -10, 2)

	if err == nil {
		t.Errorf("Update should return an error")
	}

	err = board.Update("Mexico", "Canada", 10, -2)

	if err == nil {
		t.Errorf("Update should return an error")
	}

	err = board.Update("Mexico", "Canada", -10, -2)

	if err == nil {
		t.Errorf("Update should return an error")
	}
}

func TestSummary(t *testing.T) {
	board := InitScoreBoard()
	board.Add("Mexico", "Canada")
	board.Update("Mexico", "Canada", 0, 5)
	board.Add("Spain", "Brazil")
	board.Update("Spain", "Brazil", 10, 2)
	board.Add("Germany", "France")
	board.Update("Germany", "France", 2, 8)
	board.Update("Germany", "France", 2, 10)

	games := board.Summary()

	if len(games) != 3 {
		t.Errorf("Summary should return 3 games, got %d", len(games))
	}

	if games[0].TotalScore != 12 && games[0].Home != "Germany" && games[0].Away != "France" {
		t.Errorf("Summary should return the game with the highest score, got %d, expected: 12", games[0].TotalScore)
	}

	if games[1].TotalScore != 12 && games[1].Home != "Spain" && games[1].Away != "Brazil" {
		t.Errorf("Summary should return the game with the highest score, got %d, expected: 12", games[1].TotalScore)
	}

	if games[2].TotalScore != 5 && games[2].Home != "Mexico" && games[2].Away != "Canada" {
		t.Errorf("Summary should return the game with the highest score, got %d, expected: 5", games[2].TotalScore)
	}
}
