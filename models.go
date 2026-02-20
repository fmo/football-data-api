package footballdataapi

type Season struct {
	ID              int      `json:"id"`
	StartDate       string   `json:"startDate"`
	EndDate         string   `json:"endDate"`
	CurrentMatchDay int      `json:"currentMatchDay"`
	Stages          []string `json:"stages"`
}

type Team struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type Table struct {
	Position       int `json:"position"`
	Team           Team
	PlayedGames    int    `json:"playedGames"`
	Form           string `json:"form"`
	Won            int    `json:"won"`
	Draw           int    `json:"draw"`
	Lost           int    `json:"lost"`
	Points         int    `json:"points"`
	GoalsFor       int    `json:"goalsFor"`
	GoalsAgainst   int    `json:"goalsAgainst"`
	GoalDifference int    `json:"goalDifference"`
}
