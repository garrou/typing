package dto

import (
	"strings"
	"time"
)

type UserDto struct {
	Username string    `json:"username"`
	JoinedAt time.Time `json:"joinedAt"`
}

type UserLoginDto struct {
	Username string `json:"username" binding:"required" validate:"email,max:255"`
	Password string `json:"password" binding:"required" validate:"min:8,max:50"`
}

type UserCreateDto struct {
	Username string `json:"username" binding:"required" validate:"min:3,max:50"`
	Password string `json:"password" binding:"required" validate:"min:8,max:50"`
	Confirm  string `json:"confirm" binding:"required" validate:":min:8,max:50"`
}

func (u *UserCreateDto) TrimSpace() {
	u.Username = strings.TrimSpace(u.Username)
	u.Password = strings.TrimSpace(u.Password)
	u.Confirm = strings.TrimSpace(u.Confirm)
}

func (u *UserCreateDto) IsValid() bool {
	return len(u.Password) >= 8 && u.Password == u.Confirm && len(u.Username) >= 3
}
