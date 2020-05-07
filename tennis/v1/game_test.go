package tennis_test

import (
	"testing"

	"github.com/Min-Feng/TDDkata/tennis/v1"
	"github.com/stretchr/testify/suite"
)

func TestTennisGame(t *testing.T) {
	suite.Run(t, new(TennisGameTestSuite))
}

type TennisGameTestSuite struct {
	suite.Suite
	Game *tennis.Game
}

func (this *TennisGameTestSuite) SetupTest() {
	this.Game = tennis.NewGame("Joey", "Caesar")
}

func (this *TennisGameTestSuite) Test_Love_All() {
	want := "Love_All"
	actual := this.Game.Start()
	this.Assert().Equal(want, actual)
}

func (this *TennisGameTestSuite) Test_Fifteen_Love() {
	want := "Fifteen_Love"
	actual := this.Game.Player1GetPoint()
	this.Assert().Equal(want, actual)
}

func (this *TennisGameTestSuite) Test_Thirty_Love() {
	want := "Thirty_Love"
	this.Game.Player1GetPoint()
	actual := this.Game.Player1GetPoint()
	this.Assert().Equal(want, actual)
}
