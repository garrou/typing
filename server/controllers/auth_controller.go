package controllers

import (
	"net/http"
	"typing/dto"
	"typing/entities"
	"typing/services"
	"typing/utils"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	var userDto dto.UserLoginDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.NewResponse("Invalid form", nil))
		return
	}
	res := services.Login(userDto.Username, userDto.Password)

	if user, ok := res.(entities.User); ok {
		token := utils.GenerateToken(user.ID)
		ctx.JSON(http.StatusOK, utils.NewResponse("token", token))
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewResponse("Invalid username or password", nil))
	}
}

func Register(ctx *gin.Context) {

	var userDto dto.UserCreateDto

	if errDto := ctx.ShouldBind(&userDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.NewResponse("Invalid form", nil))
		return
	}
	userDto.TrimSpace()

	if !userDto.IsValid() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.NewResponse("Invalid username or password", nil))
		return
	}
	if services.IsDuplicateUsername(userDto.Username) {
		ctx.AbortWithStatusJSON(http.StatusConflict, utils.NewResponse("An account already exists with this username", nil))
	} else {
		services.Register(userDto)
		ctx.JSON(http.StatusCreated, utils.NewResponse("Account created", nil))
	}
}

func GetAuthUser(ctx *gin.Context) {

	userId := ctx.GetString("userId")
	res := services.FindUserById(userId)

	if user, ok := res.(entities.User); ok {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.AbortWithStatus(http.StatusForbidden)
	}
}
