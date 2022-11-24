package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserID = "uerid"

var ERRorUserNotLogin = errors.New("用户未登录")

func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserID)
	if !ok {
		err = ERRorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ERRorUserNotLogin
		return
	}
	return
}

func GetPageInfo(c *gin.Context) (int64, int64) {
	//获取分页参数
	pagesetStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		page int64
		size int64
	)
	page, _ = strconv.ParseInt(pagesetStr, 10, 64)
	size, _ = strconv.ParseInt(sizeStr, 10, 64)
	return page, size
}
