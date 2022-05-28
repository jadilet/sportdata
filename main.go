package main

import "fmt"

func main() {

	scoreBoard := InitScoreBoard()
	scoreBoard.Add("Mexico", "Canada")
	scoreBoard.Update("Mexico", "Canada", 0, 5)

	scoreBoard.Add("Spain", "Brazil")
	scoreBoard.Update("Spain", "Brazil", 10, 2)

	scoreBoard.Add("Germany", "France")
	scoreBoard.Update("Germany", "France", 2, 2)

	scoreBoard.Add("Uruguay", "Italy")
	scoreBoard.Update("Uruguay", "Italy", 6, 6)

	scoreBoard.Add("Argentina", "Australia")
	scoreBoard.Update("Argentina", "Australia", 3, 1)
	scoreBoard.Update("Germany", "France", 2, 3)

	games := scoreBoard.Summary()

	for _, game := range games {
		fmt.Printf("%s-%s %d-%d %d\n", game.Home, game.Away, game.ScoreHome, game.ScoreAway, game.TotalScore)
	}
}
