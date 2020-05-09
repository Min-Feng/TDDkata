package tennis

type Player struct {
	Name   string
	Points int
}

type Game struct {
	p1       Player
	p2       Player
	scoreMap map[int]string
}

func NewGame() *Game {
	return &Game{
		p1: Player{
			Name:   "",
			Points: 0,
		},
		p2: Player{
			Name:   "",
			Points: 0,
		},
		scoreMap: map[int]string{
			0: "Love",
			1: "Fifteen",
			2: "Thirty",
			3: "Forty",
		},
	}
}

func (g *Game) Start() string {
	return "LoveAll"
}

func (g *Game) Player1GetPoint() string {
	g.p1.Points++
	return g.scoreMap[g.p1.Points] + g.scoreMap[g.p2.Points]
}
