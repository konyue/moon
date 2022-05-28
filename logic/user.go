package logic

// 存放业务逻辑代码

import (
	"moon/dao/mysql"
	"moon/pkg/snowflake"
)

func SignUp() {
	// 判断用户是否存在
	mysql.QueryUserByUsername()
	// 生成uid
	snowflake.GenID()

	// 保存进数据库
	mysql.InsertUser()

}
