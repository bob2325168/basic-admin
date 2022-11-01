package account

import (
	"github.com/bob2325168/bbs/app/http/middleware/auth"
	"github.com/bob2325168/bbs/app/provider/account"
	"github.com/bob2325168/corehero/framework/gin"
)

type AccounApi struct{}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	api := &AccounApi{}
	if !r.IsBind(account.AccountKey) {
		r.Bind(&account.AccountProvider{})
	}
	accApi := r.Group("/api/v1", auth.AuthMiddleware())
	{
		accApi.GET("/account/list", api.AccountList)
		accApi.GET("/account/detail", api.AccountInfo)
		accApi.POST("/account/edit", api.AccountEdit)
		accApi.POST("/account/add", api.AccountAdd)
		accApi.POST("/account/del", api.AccountDel)
	}
	return nil
}
