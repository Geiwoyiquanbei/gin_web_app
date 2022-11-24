package controllers

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"
	"web_app/pkg/JWT"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//
func SignUpHandler(con *gin.Context) {
	//1.获取参数和参数校验
	var p models.ParamSignUp
	err := con.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("signup with invalid param", zap.Error(err))
		errors := err.(validator.ValidationErrors)
		ResponseWithMsg(con, CodeInvalidParam, removeTopStruct(errors.Translate(trans)))
		return
	}
	//手动对参数进行校验
	//if len(p.UserName) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("signup with invalid param")
	//	con.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	//2.业务处理
	err1 := logic.SignUp(&p)
	if err1 != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err1))
		if errors.Is(err, mysql.ErrorUSerNotExist) {
			ResponseError(con, CodeUserExist)
			return
		}
		ResponseError(con, CodeServerBusy)
		return
	}
	//3.返回相应

	ResponseSuccess(con, nil)
}

func LoginHandler(c *gin.Context) {
	var p models.ParamLogin
	err := c.ShouldBindJSON(&p)
	var user = &models.User{}
	if err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		errors := err.(validator.ValidationErrors)
		ResponseWithMsg(c, CodeInvalidParam, removeTopStruct(errors.Translate(trans)))
		return
	}
	user, JWT.FreshToken, err = logic.LogIn(&p)
	if err != nil {
		zap.L().Error("login with incorrect password", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, gin.H{
		"user_id":  user.UserID, //如果id值大于 i<<53-1
		"username": user.Username,
		"token":    user.Token,
	})
	return
}
