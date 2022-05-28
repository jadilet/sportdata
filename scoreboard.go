package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type ScoreBoard struct {
	games map[string]*Game
}

func InitScoreBoard() *ScoreBoard {
	return &ScoreBoard{games: make(map[string]*Game)}
}

func (board *ScoreBoard) Add(home, away string) error {
	board.games[board.generateGameId(home, away)] = &Game{
		ID: board.generateGameId(home, away), Home: home, Away: away}

	return nil
}

func (board *ScoreBoard) Remove(home, away string) {
	delete(board.games, board.generateGameId(home, away))
}

// Update the score of a game
// Arguments: home, away, scoreHome, scoreAway
func (board *ScoreBoard) Update(home, away string, scoreHome, scoreAway int) error {
	if scoreAway < 0 || scoreHome < 0 {
		return fmt.Errorf("Score cannot be negative")
	}

	gameId := board.generateGameId(home, away)

	if _, found := board.games[gameId]; !found {
		return fmt.Errorf("The game didn't start yet")
	}

	board.games[gameId].ScoreHome = scoreHome
	board.games[gameId].ScoreAway = scoreAway
	board.games[gameId].TotalScore = scoreHome + scoreAway
	board.games[gameId].UpdatedAt = time.Now()

	return nil
}

func (board *ScoreBoard) generateGameId(home, away string) string {
	home = strings.ToLower(home)
	away = strings.ToLower(away)
	return fmt.Sprintf("%s%s", strings.Trim(home, " "), strings.Trim(away, " "))
}

func (board *ScoreBoard) Summary() []*Game {
	games := []*Game{}

	for _, game := range board.games {
		games = append(games, game)
	}

	sort.Sort(GameSorter(games))

	return games
}
