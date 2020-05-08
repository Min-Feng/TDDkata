package tennis

type Player struct {
	Name  string
	Point int
}

type Game struct {
	p1       Player
	p2       Player
	scoreMap map[int]string
}

func NewGame() *Game {
	return &Game{
		p1: Player{
			Name:  "",
			Point: 0,
		},
		p2: Player{
			Name:  "",
			Point: 0,
		},
		scoreMap: map[int]string{
			0: "Love",
			1: "Fifteen",
			2: "Forty",
			3: "LoveAll",
		},
	}
}

func (g *Game) Start() string {
	return g.Score()
}

func (g *Game) Score() string {
	var score string
	if g.p1.Point == 1 {
		score = g.scoreMap[g.p1.Point] + g.scoreMap[g.p2.Point]
	} else {
		score = g.scoreMap[g.p1.Point] + "All"
	}
	return score
}

func (g *Game) Player1GetPoint() string {
	g.p1.Point++
	return g.Score()
}
