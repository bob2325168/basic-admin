package log

import "github.com/bob2325168/corehero/framework/gin"

// SysOperaLogInfo 操作记录详情
// @Summary 操作记录详情
// @Description 操作记录详情
// @Accept  json
// @Produce  json
// @Tags 操作日志
// @Param id query int true "id"
// @Success 200 string Msg "操作成功"
// @Router /api/v1/sys-opera-log/detail [get]
func (api *logApi) SysOperaLogInfo(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
