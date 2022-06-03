package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"moon/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}