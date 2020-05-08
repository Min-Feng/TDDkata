package tennis_test

import (
	"testing"

	"github.com/Min-Feng/TDDkata/tennis/v2"
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
	ts.Game.Start()
	actual := ts.Game.Score()
	ts.Assert().Equal(want, actual)
}

func (ts *TennisGameTestSuite) Test_FifteenLove() {
	want := "FifteenLove"
	ts.Game.Player1GetPoint()
	actual := ts.Game.Score()
	ts.Assert().Equal(want, actual)
}

func (ts *TennisGameTestSuite) Test_FortyAll() {
	want := "FortyAll"
	ts.Game.Player1GetPoint()
	ts.Game.Player1GetPoint()
	actual := ts.Game.Score()
	ts.Assert().Equal(want, actual)
}


func (ts *TennisGameTestSuite) Test_l() {
	want := "FortyAll"
	ts.Game.Player1GetPoint()
	ts.Game.Player1GetPoint()
	actual := ts.Game.Score()
	ts.Assert().Equal(want, actual)
}