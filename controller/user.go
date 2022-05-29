package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"moon/logic"
	"moon/models"
	"net/http"
)

// SignUpHandler 处理注册请求
func SignUpHandler(c *gin.Context) {
	// 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，返回相应错误
		zap.L().Error("SignUp with invalid parma", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译
		})
		return
	}

	fmt.Println(p)
	// 业务处理
	if err := logic.SignUp(p); err != nil {
		//println(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
