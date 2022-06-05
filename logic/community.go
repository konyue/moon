package logic

import (
	"moon/dao/mysql"
	"moon/models"
)

// GetCommunityList 查询到所有社区(id,name) 以列表形式返回
func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库，查到所有的community并返回
	return mysql.GetCommunityList()

}

// GetCommunityDetail 根据id获取社区详情
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
