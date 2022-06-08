package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"moon/logic"
	"moon/models"
)

// PostVoteController 投票

func PostVoteController(c *gin.Context) {
	p := new(models.ParmaVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, CodeInvalidParma)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) //翻译并去掉错误提升中陪你的结构体标识
		ResponseErrorWithMsg(c, CodeInvalidParma, errData)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeSererBusy)
		return
	}
	ResponseSuccess(c, nil)
}
