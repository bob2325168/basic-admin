package user

import (
	provider "github.com/bob2325168/bbs/app/provider/user"
	"github.com/bob2325168/corehero/framework/gin"
)

type loginAddressParam struct {
	Address   string `json:"address" binding:"required"`
	Message   string `json:"message" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

// LoginWithAddress 钱包登录
// @Summary 钱包登录
// @Description 钱包登录
// @Accept  json
// @Produce  json
// @Tags 用户操作
// @Param loginParam body loginAddressParam  true "login with param"
// @Success 200 {string} Token "token"
// @Router /api/v1/user/login/address [post]
func (api *UserApi) LoginWithAddress(c *gin.Context) {
	// 验证参数
	userService := c.MustMake(provider.UserKey).(provider.Service)

	param := &loginAddressParam{}
	if err := c.ShouldBind(param); err != nil {
		c.ISetStatus(404).IText("参数错误")
		return
	}

	// 登录
	model := &provider.Wallet{Address: param.Address}
	userWithToken, err := userService.LoginWithAddress(c, model)
	if err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	if userWithToken == nil {
		c.ISetStatus(500).IText("用户不存在")
		return
	}

	// 输出
	c.ISetOkStatus().IText(userWithToken.Token)
	return
}
