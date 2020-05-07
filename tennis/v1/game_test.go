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
	this.Game = tennis.NewGame("joey", "caesar")
}

func (this *TennisGameTestSuite) Test_Love_All() {
	want := "Love_All"
	actual := this.Game.Start()
	this.Assert().Equal(want, actual)
}

func (this *TennisGameTestSuite) Test_Fifteen_Love() {
	want := "Fifteen_Love"
	times := 1
	actual := this.Play1GetPointTimes(times)
	this.Assert().Equal(want, actual)
}

func (this *TennisGameTestSuite) Test_Thirty_Love() {
	want := "Thirty_Love"
	times := 2
	actual := this.Play1GetPointTimes(times)
	this.Assert().Equal(want, actual)
}

func (this *TennisGameTestSuite) Test_Forty_Love() {
	want := "Forty_Love"
	times := 3
	actual := this.Play1GetPointTimes(times)
	this.Assert().Equal(want, actual)
}

func (this *TennisGameTestSuite) Play1GetPointTimes(times int) string {
	var actual string
	for i := 0; i < times; i++ {
		actual = this.Game.Player1GetPoint()
	}
	return actual
}
