package dto

import "time"

type ScoreCreateDto struct {
	Score  int `json:"score" binding:"required" validate:"min:0,max:1000"`
	UserId string
}

type ScoreDto struct {
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"createdAt"`
}

type BestScoreDto struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
}
