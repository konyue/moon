package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"moon/logic"
	"moon/models"
)

func CreatePostHandler(c *gin.Context) {
	// 获取参数及校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p)  error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParma)
		return
	}
	// 从c取到当前发起请求的用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorId = userID
	//创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeSererBusy)
		return
	}
	//返回相应

	ResponseSuccess(c, nil)
}
