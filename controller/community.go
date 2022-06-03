package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"moon/logic"
	"strconv"
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

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParma)
		return
	}
	//println("id.. ", id)
	// 查询到所有社区(id,name) 以列表形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("login.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeSererBusy)
		return
	}
	ResponseSuccess(c, data)
}
