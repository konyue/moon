package models

// 定义请求的参数结构体
const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParmaVoteData 投票数据
type ParmaVoteData struct {
	// UserID 从请求中获取
	PostID    string `json:"post_id" binding:"required"`              //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票(1) 反对票(-1)
}

// ParmaPostList 获取帖子列表query string参数
type ParmaPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"` //可以为空
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
	Order       string `json:"order" form:"order"`
}
