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

func NewGame(name1, name2 string) *Game {
	return &Game{
		p1: Player{
			Name:   name1,
			Points: 0,
		},
		p2: Player{
			Name:   name2,
			Points: 0,
		},
		scoreMap: map[int]string{
			0: "Love",
			1: "Fifteen",
			2: "Thirty",
			3: "",
			4: "",
		},
	}
}

func (g *Game) Start() string {
	return "Love_All"
}

func (g *Game) Player1GetPoint() string {
	g.p1.Points++
	return g.scoreMap[g.p1.Points] + "_" + g.scoreMap[g.p2.Points]
}
