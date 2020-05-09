package tennis_test

import (
	"testing"

	"github.com/Min-Feng/TDDkata/tennis/v3"
	"github.com/stretchr/testify/suite"
)

func TestTennisGame(t *testing.T) {
	suite.Run(t, new(TennisGameTestSuite))
}

type TennisGameTestSuite struct {
	suite.Suite
	Game *tennis.Game
}

func (ts *TennisGameTestSuite) SetupTest() {
	ts.Game = tennis.NewGame()
}

func (ts *TennisGameTestSuite) Test_LoveAll() {
	want := "LoveAll"
	actual := ts.Game.Start()
	ts.Assert().Equal(want, actual)
}

func (ts *TennisGameTestSuite) Test_FifteenLove() {
	want := "FifteenLove"
	var actual string
	for i := 0; i < 1; i++ {
		actual = ts.Game.Player1GetPoint()
	}
	ts.Assert().Equal(want, actual)
}

func (ts *TennisGameTestSuite) Test_ThirtyLove() {
	want := "ThirtyLove"
	ts.Game.Player1GetPoint()
	actual := ts.Game.Player1GetPoint()
	ts.Assert().Equal(want, actual)
}

func (ts *TennisGameTestSuite) Player1GetPointTimes()  {
	
}