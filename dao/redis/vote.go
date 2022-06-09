package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

/*
投票情况：
direction=1
	1. 之前没用投过票，投赞成
	2. 之前投反对，投赞成
direction=0
	1. 之前投过赞成，取消投票
	2. 之前投过反对，取消投票
direction=-1
	1. 之前没用投过票，投反对
	2. 之前投赞成，投反对

投票限制：
每个帖子自发表起一周内投票，超过不允许投票
	1. 到期侯redis持久化到mysql中
	2. 到期后删除 KeyPostVotedZSetPrefix
*/
const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 //每一票占多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func CreatePost(postID, communityID int64) error {
	//println(postID)
	pipeline := client.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 把帖子id加到社区set
	ckey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(communityID)))
	pipeline.SAdd(ckey, postID)
	_, err := pipeline.Exec()
	return err

}
func VoteForPost(userID, postID string, value float64) error {
	// 判断投票限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()

	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	// 更新分数
	oldValue := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	// 如果这次投票值和原始值保存一致，提示不允许重复投票
	if oldValue == value {
		return ErrVoteRepeated
	}
	var op float64
	if value > oldValue {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(oldValue - value)

	pipeline := client.TxPipeline()

	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)

	// 记录用户为该帖子投过票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
