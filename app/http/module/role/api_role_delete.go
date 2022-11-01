package role

import "github.com/bob2325168/corehero/framework/gin"

// RoleDel 删除角色
// @Summary 删除角色
// @Description 删除角色
// @Accept  json
// @Produce  json
// @Tags 角色管理
// @Param id query int true "删除id"
// @Success 200 string Msg "操作成功"
// @Router /api/v1/role/del [post]
func (api *RoleApi) RoleDel(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
