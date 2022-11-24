package controllers

import (
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func PostVoteController(c *gin.Context) {
	p := &models.ParamVoteData{}
	err := c.ShouldBindJSON(p)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseWithMsg(c, CodeInvalidParam, errData)
		return
	}
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeneedLogin)
	}
	err = logic.PostVote(userID, p)
	if err != nil {
		zap.L().Error("logic.PostVote error")
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
