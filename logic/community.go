package logic

import (
	"moon/dao/mysql"
	"moon/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库，查到所有的community并返回
	return mysql.GetCommunityList()

}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
