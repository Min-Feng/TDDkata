package tennis_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Min-Feng/TDDkata/tennis/v4"
)

func TestGame_Score(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}

type GameTestSuite struct {
	suite.Suite
	Game        *tennis.Game
	player1Name string
	player2Name string
}

func (ts *GameTestSuite) SetupTest() {
	ts.player1Name = "Joey"
	ts.player2Name = "caesar"
	ts.Game = tennis.NewGame(ts.player1Name, ts.player2Name)
}

func (ts *GameTestSuite) TestLoveAll() {
	ts.Assert().Equal("LoveAll", ts.Game.Score())
}

func (ts *GameTestSuite) TestFifteenLove() {
	ts.player1GetScoreTimes(1)
	ts.Assert().Equal("FifteenLove", ts.Game.Score())
}

func (ts *GameTestSuite) TestThirtyLove() {
	ts.player1GetScoreTimes(2)
	ts.Assert().Equal("ThirtyLove", ts.Game.Score())
}

func (ts *GameTestSuite) TestFortyLove() {
	ts.player1GetScoreTimes(3)
	ts.Assert().Equal("FortyLove", ts.Game.Score())
}

func (ts *GameTestSuite) TestLoveFifteen() {
	ts.player2GetScoreTimes(1)
	ts.Assert().Equal("LoveFifteen", ts.Game.Score())
}

func (ts *GameTestSuite) TestLoveThirty() {
	ts.player2GetScoreTimes(2)
	ts.Assert().Equal("LoveThirty", ts.Game.Score())
}

func (ts *GameTestSuite) TestFifteenAll() {
	ts.player1GetScoreTimes(1)
	ts.player2GetScoreTimes(1)
	ts.Assert().Equal("FifteenAll", ts.Game.Score())
}

func (ts *GameTestSuite) TestThirtyAll() {
	ts.player1GetScoreTimes(2)
	ts.player2GetScoreTimes(2)
	ts.Assert().Equal("ThirtyAll", ts.Game.Score())
}

func (ts *GameTestSuite) TestDeuce() {
	ts.player1GetScoreTimes(3)
	ts.player2GetScoreTimes(3)
	ts.Assert().Equal("Deuce", ts.Game.Score())
}

func (ts *GameTestSuite) TestPlayer1Adv() {
	ts.player1GetScoreTimes(4)
	ts.player2GetScoreTimes(3)
	ts.Assert().Equal(ts.player1Name+"Adv", ts.Game.Score())
}

func (ts *GameTestSuite) player1GetScoreTimes(times int) {
	for i := 0; i < times; i++ {
		ts.Game.Player1GetScore()
	}
}

func (ts *GameTestSuite) player2GetScoreTimes(times int) {
	for i := 0; i < times; i++ {
		ts.Game.Player2GetScore()
	}
}
