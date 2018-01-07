package service

type newMatchResponse struct {
	Id 			string 		`json:"id"`
	StartedAt	int64		`json:"started_at"`
	GridSize	int			`json:"gridsize"`
	PlayerWhite string 		`json:"playerWhite"`
	PlayerBlack string      `json:"playerBlack"`
}

type newMatchRequest struct {
	GridSize	int			`json:"gridsize"`
	PlayerWhite string 		`json:"playerWhite"`
	PlayerBlack string      `json:"playerBlack"`
}

func (request newMatchRequest) isValid() (valid bool) {
	valid = true
	if request.GridSize != 19 {
		valid = false
	}
	if request.PlayerWhite == "" {
		valid = false
	}
	if request.PlayerBlack == "" {
		valid = false
	}
	return valid
}