package tennis

import "math"

type Game struct {
	player1Score int
	scoreLookup  map[int]string
	player2Score int
	player1Name  string
	player2Name  string
}

func NewGame(player1Name string, player2Name string) *Game {
	return &Game{
		scoreLookup: map[int]string{
			0: "Love",
			1: "Fifteen",
			2: "Thirty",
			3: "Forty",
		},
		player1Name: player1Name,
		player2Name: player2Name,
	}
}

func (g *Game) Score() string {
	if g.isSameScore() {
		if g.isDeuce() {
			return g.Deuce()
		}
		return g.SameScore()
	}

	if g.isReadyToEnd() {
		if g.isAdv() {
			return g.Adv()
		}
		if g.isEnd() {
			return g.End()
		}
	}

	return g.Normal()
}

func (g *Game) Normal() string {
	return g.scoreLookup[g.player1Score] + g.scoreLookup[g.player2Score]
}

func (g *Game) End() string {
	return g.whoScoreHigher() + "Win"
}

func (g *Game) isEnd() bool {
	return math.Abs(float64(g.player1Score-g.player2Score)) == 2
}

func (g *Game) Adv() string {
	return g.whoScoreHigher() + "Adv"
}

func (g *Game) isAdv() bool {
	return math.Abs(float64(g.player1Score-g.player2Score)) == 1
}

func (g *Game) isReadyToEnd() bool {
	return g.player1Score >= 3 || g.player2Score >= 3
}

func (g *Game) SameScore() string {
	return g.scoreLookup[g.player1Score] + "All"
}

func (g *Game) Deuce() string {
	return "Deuce"
}

func (g *Game) isDeuce() bool {
	return g.player1Score >= 3
}

func (g *Game) isSameScore() bool {
	return g.player1Score == g.player2Score
}

func (g *Game) whoScoreHigher() string {
	if g.player1Score > g.player2Score {
		return g.player1Name
	}
	return g.player2Name
}

func (g *Game) Player1GetScore() {
	g.player1Score++
}

func (g *Game) Player2GetScore() {
	g.player2Score++
}
