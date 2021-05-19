package games

type Game struct {
	Id       int64    `json:"id,omitempty"`
	Dealer   string   `json:"dealer,omitempty"`
	Players  []Player `json:"players,omitempty"`
	Settings Settings `json:"settings,omitempty"`
}

//go:generate stringer -type=ScoringMode -output=scoringmode_string.g.go
type ScoringMode int

const (
	Negative ScoringMode = iota
	Standard
)

type Settings struct {
	BonusRounds bool        `json:"bonus_rounds,omitempty"`
	ScoringMode ScoringMode `json:"scoring_mode,omitempty"`
}

type Player struct {
	Name   string `json:"name,omitempty"`
	Score  int    `json:"score"`
	Bid    int    `json:"bid"`
	Tricks int    `json:"tricks"`
}
