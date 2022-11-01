package account

import "github.com/bob2325168/corehero/framework/gin"

// AccountDel 删除账户
// @Summary 删除账户
// @Description 删除账户
// @Accept  json
// @Produce  json
// @Tags 账户管理
// @Param id query int true "id"
// @Success 200 string Msg "操作成功"
// @Router /api/v1/account/del [post]
func (api *AccounApi) AccountDel(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
