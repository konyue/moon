package logic

// 存放业务逻辑代码

import (
	"moon/dao/mysql"
	"moon/models"
	"moon/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成uid
	userID := snowflake.GenID()
	// 构造一个user示例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库
	return mysql.InsertUser(user)

}
