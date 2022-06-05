package logic

// 存放业务逻辑代码

import (
	"moon/dao/mysql"
	"moon/models"
	"moon/pkg/jwt"
	"moon/pkg/snowflake"
)

// SignUp 用户注册
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

// Login 用户登录
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 生成jwt
	println("jwt!!!!!", user.UserID, user.Username)
	return jwt.GenToken(user.UserID, user.Username)
}
