package controllers

import "web_app/models"

type _ResponsePostList struct {
	Msg  string                  `json:"msg"`  //信息
	Code ResCode                 `json:"code"` //状态码
	Data []*models.ApiPostDetail `json:"data"` //数据内容
}
type _ResponseSuccessful struct {
	Msg  string      `json:"msg"`  //信息
	Code ResCode     `json:"code"` //状态码
	Data interface{} `json:"data"` //数据内容
}
