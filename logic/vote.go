package logic

import (
	"go.uber.org/zap"
	"moon/dao/redis"
	"moon/models"
	"strconv"
)

// 待实现投票算法： www.ruanyifeng.com/blog/algorithm/
// 投票功能

// 暂用算法
// 投一票+432分

// VoteForPost 为帖子投票的函数
func VoteForPost(userID int64, p *models.ParmaVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
