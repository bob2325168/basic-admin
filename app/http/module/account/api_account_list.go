package account

import "github.com/bob2325168/corehero/framework/gin"

// AccountList 获取账户列表
// @Summary 获取账户列表
// @Description 获取账户列表
// @Accept  json
// @Produce  json
// @Tags 账户管理
// @Success 200 string Msg "操作成功"
// @Router /api/v1/account/list [get]
func (api *AccounApi) AccountList(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
