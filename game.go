package main

import "time"

type Game struct {
	ID         string    `json:"id"`
	Home       string    `json:"home"`
	Away       string    `json:"away"`
	ScoreHome  int       `json:"score_home"`
	ScoreAway  int       `json:"score_away"`
	TotalScore int       `json:"total_score"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GameSorter []*Game

func (g GameSorter) Len() int {
	return len(g)
}

func (g GameSorter) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g GameSorter) Less(i, j int) bool {
	return g[i].TotalScore > g[j].TotalScore ||
		(g[i].TotalScore == g[j].TotalScore &&
			g[i].UpdatedAt.After(g[j].UpdatedAt))
}
