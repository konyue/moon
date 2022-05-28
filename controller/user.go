package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"moon/logic"
	"moon/models"
	"net/http"
)

// SignUpHandler 处理注册请求
func SignUpHandler(c *gin.Context) {
	// 获取参数和参数校验
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，返回相应
		zap.L().Error("SignUp with invalid parma", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	// 对请求参数进行详细的业务规则校验
	// 。。。。。代码

	fmt.Println(p)
	// 业务处理
	logic.SignUp()
	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
