package repositories

import (
	"typing/database"
	"typing/dto"
	"typing/entities"
)

func SaveScore(score entities.Score) entities.Score {
	database.Db.Save(&score)
	return score
}

func FindScoreByUserId(userId string) []entities.Score {
	var scores []entities.Score

	database.Db.
		Order("created_at DESC").
		Find(&scores, "user_id = ?", userId)

	return scores
}

func FindBestScores() []dto.BestScoreDto {
	var scores []dto.BestScoreDto
	database.Db.
		Model(entities.Score{}).
		Select("DISTINCT users.username AS username, MAX(scores.score) AS score").
		Joins("JOIN users ON users.id = scores.user_id").
		Group("username").
		Order("score DESC").
		Scan(&scores)
	return scores
}
