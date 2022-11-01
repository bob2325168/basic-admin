package wallet

import (
	provider "github.com/bob2325168/bbs/app/provider/user"
	"github.com/bob2325168/corehero/framework/gin"
)

/**
// GetAddressInfo 获取钱包信息
// @Summary 获取钱包信息
// @Description 获取钱包信息
// @Accept  json
// @Produce  json
// @Tags 用户操作
// @Param address query string true "钱包地址"
// @Success 200 {string} Message "成功获取钱包信息"
// @Router /api/v1/user/addressinfo [get]
*/

func (api *WalletApi) GetAddressInfo(c *gin.Context) {
	// 验证参数
	userService := c.MustMake(provider.UserKey).(provider.Service)
	token := c.Query("token")
	if token == "" {
		c.ISetStatus(404).IText("参数错误")
		return
	}

	verified, err := userService.VerifyRegister(c, token)
	if err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}

	if !verified {
		c.ISetStatus(500).IText("验证错误")
		return
	}

	// 输出
	c.ISetOkStatus().IText("注册成功，请进入登录页面")
}
