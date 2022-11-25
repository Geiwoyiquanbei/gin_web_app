package logic

import (
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//1.生成post id
	p.ID = snowflake.GetID()
	//2.保存到数据库
	mysql.CreatePost(p)
	//3.保存到redis
	err = redis.SavePostTime(p.ID, p.CommunityID)
	if err != nil {
		return err
	}
	return nil
}
func GetPostDetail(id int64) (data *models.ApiPostDetail, err error) {
	data = new(models.ApiPostDetail)
	detail, err := mysql.GetPostDetail(id)
	if err != nil {
		return nil, err
	}
	user, err := mysql.GetUserByID(detail.AuthorID)
	if err != nil {
		return nil, err
	}
	community, err := mysql.GetCommunity(detail.CommunityID)
	if err != nil {
		return nil, err
	}
	data.Post = detail
	data.CommunityDetail = community
	data.AuthuorName = user.Username
	return data, nil

}
func GetPostList(limit, offset int64) (data []*models.ApiPostDetail, err error) {
	list, err := mysql.GetPostList(limit, offset)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(list))
	for _, post := range list {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			return nil, err
		}
		community, err := mysql.GetCommunity(post.CommunityID)
		if err != nil {
			return nil, err
		}
		detail := &models.ApiPostDetail{
			AuthuorName:     user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, detail)
	}
	return data, nil
}
func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetPostIDsInOrder(p)
	if len(ids) == 0 {
		return
	}
	if err != nil {
		return
	}
	list, err := mysql.GetPostByOrder(ids)
	voteData, err := redis.GetPostVoteData(ids)
	for idx, post := range list {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			return nil, err
		}
		community, err := mysql.GetCommunity(post.CommunityID)
		if err != nil {
			return nil, err
		}
		detail := &models.ApiPostDetail{
			AuthuorName:     user.Username,
			Votes:           voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, detail)
	}
	return data, nil
}
func GetCommunityList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	orderkey := redis.GetRedisKey(redis.KeyPostTimeZSet)
	if p.Order != "time" {
		orderkey = redis.GetRedisKey(redis.KeyPostScoreZSet)
	}
	ids, err := redis.GetCommunityPostIDsInOrder(orderkey, p.CommunityID, p.Page, p.Size)
	if len(ids) == 0 {
		return
	}
	if err != nil {
		return
	}
	list, err := mysql.GetPostByOrder(ids)
	voteData, err := redis.GetPostVoteData(ids)
	for idx, post := range list {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			return nil, err
		}
		community, err := mysql.GetCommunity(post.CommunityID)
		if err != nil {
			return nil, err
		}
		detail := &models.ApiPostDetail{
			AuthuorName:     user.Username,
			Votes:           voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, detail)
	}
	return data, nil
}
