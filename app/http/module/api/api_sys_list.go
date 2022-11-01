package api

import "github.com/bob2325168/corehero/framework/gin"

// SysApiList api接口列表
// @Summary  api接口列表
// @Description  api接口列表
// @Accept  json
// @Produce  json
// @Tags 接口管理
// @Success 200 string Msg "操作成功"
// @Router /api/v1/sys-api/list [get]
func (api *SysApi) SysApiList(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
