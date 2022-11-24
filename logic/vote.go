package logic

import (
	"strconv"
	"web_app/dao/redis"
	"web_app/models"
)

//投票功能
//1.用户投票

/* 投票的几种情况
direction = 1 时，有两种情况：
	1.之前没有投过票，现在投赞成票
	2.之前投过反对票，现在改投赞成票
direction = 0 时 ，有两种情况:
	1，之前投过赞成票，现在要取消投票
	2．之前投过反对票，现在要取消投票
direction=-1时，有两种情况:
	1．之前没有投过票，现在投反对票
	2．之前投赞成票，现在改投反对票
投票限制：
每个贴子自发表之日起一个星期之内允许用户投票，超过一个暴期就不允许再投票了。
1。到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
2，到期之后删除那个KeyPostVotedzSetPF
*/
func PostVote(userID int64, p *models.ParamVoteData) error {
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
	//1.判断投票的限制
	//2.更新分数
	//3.记录该用户操作的记录
}
