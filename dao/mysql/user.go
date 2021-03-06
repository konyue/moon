package mysql

// 每一步数据库操作封装函数
// 等待logic进行调用

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"moon/models"
)

const secret = "konyue.com"

// CheckUserExist 通过用户名查重
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username= ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser  向数据库中插入新的用户记录
func InsertUser(user *models.User) (err error) {
	// 密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行sql语句
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Login 用户登录
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	//println(password, user.Password)
	return
}

// GetUserById 根据id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id=?`
	err = db.Get(user, sqlStr, uid)
	return
}
