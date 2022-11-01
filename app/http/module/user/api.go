package user

import (
	"github.com/bob2325168/bbs/app/http/middleware/auth"
	"github.com/bob2325168/bbs/app/provider/user"
	"github.com/bob2325168/corehero/framework/gin"
)

type UserApi struct{}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	api := &UserApi{}
	if !r.IsBind(user.UserKey) {
		r.Bind(&user.UserProvider{})
	}
	r.Group("/api/v1", nil)
	{
		// 账密登录
		r.POST("/user/login/username", api.LoginWithUsername)
		// 钱包登录
		r.POST("/user/login/address", api.LoginWithAddress)
		// 登出
		r.GET("/user/logout", auth.AuthMiddleware(), api.Logout)
		// 注册
		r.POST("/user/register", api.Register)
		// 注册验证
		r.GET("/user/register/verify", api.Verify)
		// 用户的详情
		r.GET("/user/detail", api.GetUserInfo)
	}
	return nil
}
