package role

import (
	"github.com/bob2325168/bbs/app/http/middleware/auth"
	"github.com/bob2325168/bbs/app/provider/role"
	"github.com/bob2325168/corehero/framework/gin"
)

type RoleApi struct{}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	api := &RoleApi{}
	if !r.IsBind(role.RoleKey) {
		r.Bind(&role.RoleProvider{})
	}
	roleApi := r.Group("/api/v1", auth.AuthMiddleware())
	{
		roleApi.GET("/role/list", api.RoleList)
		roleApi.GET("/role/detail", api.RoleInfo)
		roleApi.POST("/role/edit", api.RoleEdit)
		roleApi.POST("/role/del", api.RoleDel)
		roleApi.POST("/role/add", api.RoleAdd)
	}
	return nil
}
