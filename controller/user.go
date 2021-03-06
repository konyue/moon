package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"moon/dao/mysql"
	"moon/logic"
	"moon/models"
)

// SignUpHandler 处理注册请求
// @Summary 注册接口
// @Description 处理用户注册请求
// @Tags 注册相关接口
// @Accept application/json
// @Produce application/json
// @Param object  query models.ParamSignUp false "注册参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	// 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，返回相应错误
		zap.L().Error("SignUp with invalid parma", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParma)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParma, removeTopStruct(errs.Translate(trans)))
		return
	}

	//fmt.Println(p)
	// 业务处理
	if err := logic.SignUp(p); err != nil {
		//println(err)
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeSererBusy)
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "注册失败",
		//})
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求
func LoginHandler(c *gin.Context) {
	// 获取参数和参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid parma", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParma)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParma, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 业务处理
	user, err := logic.Login(p)
	//println(token, err, "ddd ~!~~~")
	if err != nil {
		zap.L().Error("login.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	// 返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
