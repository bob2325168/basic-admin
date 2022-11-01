package log

import "github.com/bob2325168/corehero/framework/gin"

// SysOperaLogDel 删除操作记录
// @Summary 删除操作记录
// @Description 删除操作记录
// @Accept  json
// @Produce  json
// @Tags 操作日志
// @Param id query int true "id"
// @Success 200 string Msg "操作成功"
// @Router /api/v1/sys-opera-log/del [post]
func (api *logApi) SysOperaLogDel(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
