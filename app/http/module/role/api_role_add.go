package role

import "github.com/bob2325168/corehero/framework/gin"

// RoleAdd 创建角色
// @Summary 创建角色
// @Description 创建角色
// @Accept  json
// @Produce  json
// @Tags 角色管理
// @Success 200 string Msg "操作成功"
// @Router /api/v1/role/add [post]
func (api *RoleApi) RoleAdd(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
