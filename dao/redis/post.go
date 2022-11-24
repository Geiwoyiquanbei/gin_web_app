package redis

import (
	"web_app/models"

	"github.com/go-redis/redis"
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := GetRedisKey(KeyPostTimeZSet)
	if p.Order == "score" {
		key = GetRedisKey(KeyPostScoreZSet)
	}
	//2.确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//3. zrevrange  按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	pipeline := client.Pipeline()
	for index, _ := range ids {
		pipeline.ZCount(KeyPostVotedPrefix+ids[index], "1", "1")
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
