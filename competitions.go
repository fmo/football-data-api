package footballdataapi

type CompetitionCode string

const (
	PL CompetitionCode = "PL"
)

var competitions = map[CompetitionCode]struct{}{
	PL: {},
}

type Competition struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Emblem string `json:"emblem"`
}
