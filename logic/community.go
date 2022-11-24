package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() ([]models.Community, error) {
	//查找数据库  查到所有的community 并返回
	return mysql.GetCommunityList()
}
func GetCommunity(id int64) (*models.CommunityDetail, error) {
	data, err := mysql.GetCommunity(id)
	if err != nil {
		return data, err
	}
	return data, nil
}
