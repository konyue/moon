package mysql

// 每一步数据库操作封装函数
// 等待logic进行调用

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
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
		return errors.New("用户已存在")
	}
	return
}

// InsertUser  向数据库中年插入新的用户记录
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

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return errors.New("密码错误")
	}
	return
}
