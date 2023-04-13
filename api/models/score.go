package models

type Score struct {
	ScoreWorth float64 `json:"score_worth"`
	UserID     string  `json:"user_id"`
	Timestamp  int64   `json:"timestamp"`
}
