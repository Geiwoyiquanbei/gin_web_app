package controllers

import (
	"strconv"
	"web_app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	//查询所有的社区（community_id , community_name) 以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
func CommunityDetailHandler(c *gin.Context) {
	//1.获取社区id
	communityID := c.Param("id")
	id, err2 := strconv.ParseInt(communityID, 10, 64)
	if err2 != nil {
		zap.L().Error("id is not valid", zap.Error(err2))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunity(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
