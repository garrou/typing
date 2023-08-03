package repositories

import (
	"typing/database"
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
