package redis

// redis key

const (
	Prefix                 = "moon:"
	KeyPostTimeZSet        = "post:time"   // zset; 帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  // zset; 帖子及投票分数
	KeyPostVotedZSetPrefix = "post:voted:" //zset; 记录用户和投票类型;参数是post_id
	KeyCommunitySetPrefix  = "community:"  //set; 保存每个分区下帖子的id
)

// getRedisKey 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
