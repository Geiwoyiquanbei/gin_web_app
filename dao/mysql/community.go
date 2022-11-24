package mysql

import (
	"database/sql"
	"web_app/models"

	"go.uber.org/zap"
)

func GetCommunityList() ([]models.Community, error) {
	sqlStr := "select community_id,community_name from community"
	communityList := []models.Community{}
	err := db.Select(&communityList, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("There is no community in db")
			err = nil
		}
	}
	return communityList, nil
}
func GetCommunity(id int64) (*models.CommunityDetail, error) {
	var community = &models.CommunityDetail{}
	sqlStr := `select community_id,community_name,introduction,create_time from community where community_id = ?`
	err := db.Get(community, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("There is no community in db")
			err = nil
		}
	}
	return community, nil
}
