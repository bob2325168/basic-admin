package log

import (
	"github.com/bob2325168/bbs/app/http/middleware/auth"
	"github.com/bob2325168/bbs/app/provider/log"
	"github.com/bob2325168/corehero/framework/gin"
)

type logApi struct{}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	api := &logApi{}
	if !r.IsBind(log.OperaLogKey) {
		r.Bind(&log.OperaLogProvider{})
	}
	sysApi := r.Group("/api/v1", auth.AuthMiddleware())
	{
		sysApi.GET("/sys-opera-log/list", api.SysOperaLogList)
		sysApi.GET("/sys-opera-log/detail", api.SysOperaLogInfo)
		sysApi.POST("/sys-opera-log/del", api.SysOperaLogDel)
	}
	return nil
}
