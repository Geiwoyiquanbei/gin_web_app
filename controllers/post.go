package controllers

import (
	"fmt"
	"strconv"
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const (
	orderTime  = "time"
	orderScore = "score"
)

//CreatePostHandler  创建帖子接口
// @Summary 创建帖子列表接口
// @Description 可创建帖子
// @Tags 创建帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.Post false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSuccessful
// @Router /post [post]
func CreatePostHandler(c *gin.Context) {
	//1.获取参数及参数校验
	p := new(models.Post)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeneedLogin)
		return
	}
	p.AuthorID = id
	//2.创建帖子
	err = logic.CreatePost(p)
	if err != nil {
		zap.L().Error("logic.CreatePost(p) field", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

//获取帖子详情

//GetPostDetailHandler  升级版列表接口
// @Summary 获取帖子详细内容
// @Description 可按帖子id号来查询帖子内容
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /post/:id [get]
func GetPostDetailHandler(c *gin.Context) {
	//1.获取参数（从url中获取id）
	id := c.Param("id")
	pid, err2 := strconv.ParseInt(id, 10, 64)
	if err2 != nil {
		zap.L().Error("get post detail with invalid param1", zap.Error(nil))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if len(id) == 0 {
		zap.L().Error("get post detail with invalid param2", zap.Error(nil))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostDetail(pid)
	if err != nil {
		zap.L().Error("get post detail with invalid param3", zap.Error(nil))
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c, data)
}

//GetPostDetailHandler  升级版列表接口
// @Summary 获取帖子列表
// @Description 可按页码来查询帖子内容
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts [get]
func GetPostListHandler(c *gin.Context) {
	page, size := GetPageInfo(c)
	data, err := logic.GetPostList(size, page)
	if err != nil {
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}

//根据前端传来的参数的动态来获取帖子列表
//按创建时间排序 或者按照分数排序
//1.获取参数
//2.去redis查询id 列表
//3.根据id去数据库查询帖子的详细信息

//GetPostListHandler2  升级版列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object}  _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	//get 请求参数 ：/api/v1/post2?offset=1&limit=10&order=time
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: orderTime,
	}
	err2 := c.ShouldBindQuery(p)
	fmt.Println(p)
	if err2 != nil {
		zap.L().Error("GetPostListHandler2 with invalid param", zap.Error(err2))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if p.CommunityID == 0 {
		//c.ShouldBindJSON() 如果请求中携带的json格式数据，才能用这个方法获取数据
		data, err := logic.GetPostList2(p)
		if err != nil {
			ResponseError(c, CodeServerBusy)
		}
		ResponseSuccess(c, data)
	} else {
		data, err := logic.GetCommunityList2(p)
		if err != nil {
			ResponseError(c, CodeServerBusy)
		}
		ResponseSuccess(c, data)
	}

}
