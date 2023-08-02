package repositories

import (
	"typing/database"
	"typing/entities"

	"gorm.io/gorm"
)

func SaveUser(user entities.User) entities.User {
	database.Db.Save(&user)
	return user
}

func FindUserByUsername(username string) interface{} {
	var user entities.User
	database.Db.Find(&user, "username = ?", username)
	return user
}

func FindUserById(id string) interface{} {
	var user entities.User
	database.Db.Find(&user, "id = ?", id)
	return user
}

func UsernameExists(username string) *gorm.DB {
	return database.Db.Take(&entities.User{}, "username = ?", username)
}
