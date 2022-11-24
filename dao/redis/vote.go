package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

/* 投票的几种情况
direction = 1 时，有两种情况：
	1.之前没有投过票，现在投赞成票  -->更新分数和投票记录   差值的绝对值: 1   +432
	2.之前投过反对票，现在改投赞成票-->更新分数和投票记录﹐差值的绝对值:2      +432*2
direction = 0 时 ，有两种情况:
	1，之前投过赞成票，现在要取消投票 -->更新分数和投票记录差值的绝对值:1      -432
	2．之前投过反对票，现在要取消投票 -->更新分数和投票记录﹐差值的绝对值:1     +432
direction=-1时，有两种情况:
	1．之前没有投过票，现在投反对票-->更新分数和投票记录差值的绝对值:1         -432
	2．之前投赞成票，现在改投反对票-->更新分数和投票记录差值的绝对值:2         -432*2
投票限制：
每个贴子自发表之日起一个星期之内允许用户投票，超过一个暴期就不允许再投票了。
1。到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
2，到期之后删除那个KeyPostVotedzSetPF
*/
const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 //每一票占多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")

	ErrVoteRepeat = errors.New("请勿重复投票")
)

func SavePostTime(postID, communityID int64) (err error) {
	pipeline := client.TxPipeline()
	//储存帖子时间
	pipeline.ZAdd(GetRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//储存帖子分数
	pipeline.ZAdd(GetRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//储存帖子id
	ckey := GetRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(ckey, postID)
	_, err = pipeline.Exec()
	if err != nil {
		return err
	}
	return nil
}
func VoteForPost(userID, postID string, value float64) error {
	pipeline := client.TxPipeline() //放在同一事务中
	//1.判断投票的限制  //去redis 取发帖时间
	postTime := client.ZScore(GetRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//2.更新分数
	//先查当前用户之前的投票记录
	ovalue := client.ZScore(GetRedisKey(KeyPostVotedPrefix+postID), userID).Val()
	if ovalue == value {
		return ErrVoteRepeat
	}
	var op float64 = 0
	if value > ovalue {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ovalue - value)
	pipeline.ZIncrBy(GetRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	//3.记录该用户操作的记录
	if value == 0 {
		pipeline.ZRem(GetRedisKey(KeyPostVotedPrefix+postID), userID)
	} else {
		pipeline.ZAdd(GetRedisKey(KeyPostVotedPrefix+postID), redis.Z{
			value,
			userID,
		})
	}
	_, err := pipeline.Exec()
	if err != nil {
		if ErrVoteTimeExpire != nil {
			return err
		}
		return err
	}
	return nil
}
