package services

import (
	"time"
	"typing/dto"
	"typing/entities"
	"typing/repositories"
	"typing/utils"

	"github.com/google/uuid"
)

func IsDuplicateUsername(username string) bool {
	res := repositories.UsernameExists(username)
	return res.Error == nil
}

func FindUserById(id string) interface{} {
	return repositories.FindUserById(id)
}

func Register(user dto.UserCreateDto) dto.UserDto {
	toCreate := entities.User{
		ID:       uuid.New().String(),
		Username: user.Username,
		Password: utils.HashPassword(user.Password),
		JoinedAt: time.Now(),
	}
	created := repositories.SaveUser(toCreate)

	return dto.UserDto{
		Username: created.Username,
		JoinedAt: created.JoinedAt,
	}
}

func Login(username, password string) interface{} {
	res := repositories.FindUserByUsername(username)

	if user, ok := res.(entities.User); ok {
		same := utils.ComparePassword(user.Password, password)

		if user.Username == username && same {
			return res
		}
		return false
	}
	return false
}
