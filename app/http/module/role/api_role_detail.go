package role

import "github.com/bob2325168/corehero/framework/gin"

// RoleInfo 查看角色详情
// @Summary 查看角色详情
// @Description 查看角色详情
// @Accept  json
// @Produce  json
// @Tags 角色管理
// @Param id query int true "id"
// @Success 200 string Msg "操作成功"
// @Router /api/v1/role/detail [get]
func (api *RoleApi) RoleInfo(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
