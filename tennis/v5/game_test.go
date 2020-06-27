package tennis_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/Min-Feng/TDDkata/tennis/v5"
)

func TestGame_Score(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}

type GameTestSuite struct {
	suite.Suite
	Game *tennis.Game
}

func (ts *GameTestSuite) SetupTest() {
	ts.Game = tennis.NewGame("Joey", "Caesar")
}

func (ts *GameTestSuite) TestLoveAll() {
	ts.Assert().Equal("LoveAll", ts.Game.Score())
}

func (ts *GameTestSuite) TestFifteenLove() {
	ts.Player1GetScoreTimes(1)
	ts.Assert().Equal("FifteenLove", ts.Game.Score())
}

func (ts *GameTestSuite) TestThirtyLove() {
	ts.Player1GetScoreTimes(2)
	ts.Assert().Equal("ThirtyLove", ts.Game.Score())
}

func (ts *GameTestSuite) TestFortyLove() {
	ts.Player1GetScoreTimes(3)
	ts.Assert().Equal("FortyLove", ts.Game.Score())
}

func (ts *GameTestSuite) TestLoveFifteen() {
	ts.Player2GetScoreTimes(1)
	ts.Assert().Equal("LoveFifteen", ts.Game.Score())
}

func (ts *GameTestSuite) TestLoveThirty() {
	ts.Player2GetScoreTimes(2)
	ts.Assert().Equal("LoveThirty", ts.Game.Score())
}

func (ts *GameTestSuite) TestFifteenAll() {
	ts.Player1GetScoreTimes(1)
	ts.Player2GetScoreTimes(1)
	ts.Assert().Equal("FifteenAll", ts.Game.Score())
}

func (ts *GameTestSuite) TestThirtyAll() {
	ts.Player1GetScoreTimes(2)
	ts.Player2GetScoreTimes(2)
	ts.Assert().Equal("ThirtyAll", ts.Game.Score())
}

func (ts *GameTestSuite) TestDeuce() {
	ts.Player1GetScoreTimes(3)
	ts.Player2GetScoreTimes(3)
	ts.Assert().Equal("Deuce", ts.Game.Score())
}

func (ts *GameTestSuite) TestDeuce44() {
	ts.Player1GetScoreTimes(4)
	ts.Player2GetScoreTimes(4)
	ts.Assert().Equal("Deuce", ts.Game.Score())
}

func (ts *GameTestSuite) TestPlayer1Adv() {
	ts.Player1GetScoreTimes(4)
	ts.Player2GetScoreTimes(3)
	ts.Assert().Equal("JoeyAdv", ts.Game.Score())
}

func (ts *GameTestSuite) TestPlayer2Adv() {
	ts.Player1GetScoreTimes(3)
	ts.Player2GetScoreTimes(4)
	ts.Assert().Equal("CaesarAdv", ts.Game.Score())
}

func (ts *GameTestSuite) TestPlayer2Win() {
	ts.Player1GetScoreTimes(3)
	ts.Player2GetScoreTimes(5)
	ts.Assert().Equal("CaesarWin", ts.Game.Score())
}

func (ts *GameTestSuite) TestPlayer1Win() {
	ts.Player1GetScoreTimes(5)
	ts.Player2GetScoreTimes(3)
	ts.Assert().Equal("JoeyWin", ts.Game.Score())
}

func (ts *GameTestSuite) Player1GetScoreTimes(times int) {
	for i := 0; i < times; i++ {
		ts.Game.Player1GetScore()
	}
}

func (ts *GameTestSuite) Player2GetScoreTimes(times int) {
	for i := 0; i < times; i++ {
		ts.Game.Player2GetScore()
	}
}
