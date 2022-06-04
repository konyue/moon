package logic

import (
	"moon/dao/mysql"
	"moon/models"
	"moon/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 生成post id
	p.ID = snowflake.GenID()
	//保存到数据库
	return mysql.CreatePost(p)
	//返回
}
