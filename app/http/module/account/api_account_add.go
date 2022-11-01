package account

import (
	provider "github.com/bob2325168/bbs/app/provider/account"
	"github.com/bob2325168/corehero/framework/gin"
)

// AccountAdd 创建账户
// @Summary 创建账户
// @Description 创建账户
// @Accept  json
// @Produce  json
// @Tags 账户管理
// @Success 200 string Msg "操作成功"
// @Router /api/v1/account/add [post]
func (api *AccounApi) AccountAdd(c *gin.Context) {
	accService := c.MustMake(provider.AccountKey).(provider.Service)
	if _, err := accService.AddAccount(c, nil); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	c.ISetOkStatus().IText("操作成功")
}
