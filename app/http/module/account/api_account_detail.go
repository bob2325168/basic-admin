package account

import "github.com/bob2325168/corehero/framework/gin"

// AccountInfo 获取账户详情
// @Summary 获取账户详情
// @Description 获取账户详情
// @Accept  json
// @Produce  json
// @Tags 账户管理
// @Param id query int true "id"
// @Success 200 string Msg "操作成功"
// @Router /api/v1/account/detail [get]
func (api *AccounApi) AccountInfo(c *gin.Context) {
	c.ISetOkStatus().IText("操作成功")
}
