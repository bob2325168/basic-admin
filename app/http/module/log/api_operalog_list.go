package log

import "github.com/bob2325168/corehero/framework/gin"

// SysOperaLogList 操作记录列表
// @Summary 操作记录列表
// @Description 操作记录列表
// @Accept  json
// @Produce  json
// @Tags 操作日志
// @Success 200 string Msg "操作成功"
// @Router /api/v1/sys-opera-log/list [get]
func (api *logApi) SysOperaLogList(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
