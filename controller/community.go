package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"moon/logic"
)

// --- 社区相关的

func CommunityHandler(c *gin.Context) {
	// 查询到所有社区(id,name) 以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("login.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeSererBusy)
		return
	}
	ResponseSuccess(c, data)
}
