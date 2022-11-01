package api

import (
	"github.com/bob2325168/bbs/app/http/middleware/auth"
	"github.com/bob2325168/bbs/app/provider/api"
	"github.com/bob2325168/corehero/framework/gin"
)

type SysApi struct{}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	sys := &SysApi{}
	if !r.IsBind(api.ApiKey) {
		r.Bind(&api.ApiProvider{})
	}
	sysApi := r.Group("/api/v1", auth.AuthMiddleware())
	{
		sysApi.GET("/sys-api/list", sys.SysApiList)
		sysApi.POST("/sys-api/edit", sys.SysApiEdit)
		sysApi.POST("/sys-api/del", sys.SysApiDel)
	}
	return nil
}
