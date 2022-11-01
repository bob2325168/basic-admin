package api

import "github.com/bob2325168/corehero/framework/gin"

// SysApiEdit 更新api接口
// @Summary 更新api接口
// @Description 更新api接口
// @Accept  json
// @Produce  json
// @Tags 接口管理
// @Param id query int true "id"
// @Success 200 string Msg "操作成功"
// @Router /api/v1/sys-api/edit [post]
func (api *SysApi) SysApiEdit(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
