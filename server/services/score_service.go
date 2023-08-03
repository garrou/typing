package services

import (
	"time"
	"typing/dto"
	"typing/entities"
	"typing/repositories"
)

func SaveScore(score dto.ScoreCreateDto) dto.ScoreDto {
	toCreate := entities.Score{
		Score:     score.Score,
		CreatedAt: time.Now(),
		UserID:    score.UserId,
	}
	created := repositories.SaveScore(toCreate)

	return dto.ScoreDto{
		Score:     created.Score,
		CreatedAt: created.CreatedAt,
	}
}

func GetScores(userId string) []dto.ScoreDto {
	res := repositories.FindScoreByUserId(userId)
	scores := []dto.ScoreDto{}

	for _, score := range res {
		scores = append(scores, dto.ScoreDto{
			Score:     score.Score,
			CreatedAt: score.CreatedAt,
		})
	}
	return scores
}
