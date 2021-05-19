package games

type Game struct {
	Id           int64    `json:"id,omitempty"`
	Dealer       string   `json:"dealer,omitempty"`
	Players      []Player `json:"players,omitempty"`
	Settings     Settings `json:"settings,omitempty"`
	MaxRounds    int      `json:"maxRounds,omitempty"`
	MaxCards     int      `json:"maxCards,omitempty"`
	Round        int      `json:"round,omitempty"`
	Cards        int      `json:"cards,omitempty"`
	IsBonusRound int      `json:"isBonusRound,omitempty"`
}

//go:generate stringer -type=ScoringMode -output=scoringmode_string.g.go
type ScoringMode int

const (
	Negative ScoringMode = iota
	Standard
)

type Settings struct {
	BonusRounds bool        `json:"bonusRounds,omitempty"`
	ScoringMode ScoringMode `json:"scoringMode,omitempty"`
}

type Player struct {
	Name   string `json:"name,omitempty"`
	Score  int    `json:"score"`
	Bid    int    `json:"bid"`
	Tricks int    `json:"tricks"`
}
