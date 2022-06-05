package routes

import (
	"github.com/gin-gonic/gin"
	"moon/controller"
	"moon/logger"
	"moon/middlewares"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) //用中间件
	{
		// 社区列表查询
		v1.GET("/community", controller.CommunityHandler)
		// 社区id查询
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		// 创建帖子
		v1.POST("/post", controller.CreatePostHandler)
		// 根据id查询帖子详情
		v1.GET("/post/:id", controller.GetPostDetailHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
