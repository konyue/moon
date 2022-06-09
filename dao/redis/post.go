package redis

import "moon/models"

// GetPostIDsInOrder 按排序规则查询帖子id
func GetPostIDsInOrder(p *models.ParmaPostList) ([]string, error) {
	// 从redis获取id
	// 根据用户请求携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 确定查询的索引
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// ZREVRANGE 按分数从大到小查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}
