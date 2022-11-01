package role

import "github.com/bob2325168/corehero/framework/gin"

// RoleEdit 编辑角色
// @Summary 编辑角色
// @Description 编辑角色
// @Accept  json
// @Produce  json
// @Tags 角色管理
// @Param id query int true "id"
// @Success 200 string Msg "操作成功"
// @Router /api/v1/role/edit [post]
func (api *RoleApi) RoleEdit(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
