package http

import (
	"github.com/bob2325168/bbs/app/http/module/account"
	"github.com/bob2325168/bbs/app/http/module/api"
	"github.com/bob2325168/bbs/app/http/module/role"
	"github.com/bob2325168/bbs/app/http/module/user"
	"github.com/bob2325168/corehero/framework/contract"
	"github.com/bob2325168/corehero/framework/gin"
	ginSwagger "github.com/bob2325168/corehero/framework/middleware/gin-swagger"
	"github.com/bob2325168/corehero/framework/middleware/gin-swagger/swaggerFiles"
	"github.com/bob2325168/corehero/framework/middleware/static"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {
	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// /路径先去./dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	// 如果配置了swagger，则显示swagger的中间件
	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// 用户模块路由
	user.RegisterRoutes(r)
	// 角色模块路由
	role.RegisterRoutes(r)
	// 账户模块路由
	account.RegisterRoutes(r)
	// api接口路由
	api.RegisterRoutes(r)
}
