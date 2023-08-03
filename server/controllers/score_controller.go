package controllers

import (
	"net/http"
	"typing/dto"
	"typing/services"
	"typing/utils"

	"github.com/gin-gonic/gin"
)

func PostScore(ctx *gin.Context) {

	var scoreDto dto.ScoreCreateDto

	if errDto := ctx.ShouldBind(&scoreDto); errDto != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.NewResponse("Invalid score", nil))
		return
	}
	scoreDto.UserId = ctx.GetString("userId")
	score := services.SaveScore(scoreDto)
	ctx.JSON(http.StatusCreated, utils.NewResponse("Score saved", score))
}

func GetScores(ctx *gin.Context) {
	scores := services.GetScores(ctx.GetString("userId"))
	ctx.JSON(http.StatusOK, scores)
}
