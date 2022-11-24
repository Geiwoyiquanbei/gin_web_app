package redis

import (
	"strconv"
	"time"
	"web_app/models"

	"github.com/go-redis/redis"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	return client.ZRevRange(key, start, end).Result()
}
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := GetRedisKey(KeyPostTimeZSet)
	if p.Order == "score" {
		key = GetRedisKey(KeyPostScoreZSet)
	}
	//2.确定查询的索引起始点
	return getIDsFormKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	pipeline := client.Pipeline()
	for index, _ := range ids {
		pipeline.ZCount(GetRedisKey(KeyPostVotedPrefix+ids[index]), "1", "1")
	}
	exec, err := pipeline.Exec()
	if err != nil {
		return nil, nil
	}
	data = make([]int64, 0, len(ids))
	for _, cmder := range exec {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

//按社区id查找帖子的ids
func GetCommunityPostIDsInOrder(orderKey string, communtityID int64, page, size int64) ([]string, error) {
	//使用 zinterstore 把分区的帖子与帖子分数的 zset 生成新的 zset
	//利用缓存key 减少zinterstore 的执行次数
	ckey := GetRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communtityID)))
	key := orderKey + strconv.Itoa(int(communtityID))
	if client.Exists(key).Val() < 1 {
		client.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, ckey, orderKey)
		client.Expire(key, 60*time.Second) //设置超时时间
	}
	return getIDsFormKey(key, page, size)
}
