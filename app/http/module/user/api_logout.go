package user

import (
	"github.com/bob2325168/bbs/app/http/middleware/auth"
	provider "github.com/bob2325168/bbs/app/provider/user"
	"github.com/bob2325168/corehero/framework/gin"
)

// Logout 用户登出
// @Summary 用户登出
// @Description 用户登出
// @Accept  json
// @Produce  json
// @Tags 用户操作
// @Success 200 {string} Message "用户登出成功"
// @Router /api/v1/user/logout [get]
func (api *UserApi) Logout(c *gin.Context) {
	authUser := auth.GetAuthUser(c)
	if authUser == nil {
		c.ISetStatus(500).IText("用户未登录")
		return
	}

	userService := c.MustMake(provider.UserKey).(provider.Service)
	if err := userService.Logout(c, authUser); err != nil {
		c.ISetStatus(500).IText(err.Error())
		return
	}
	c.ISetOkStatus().IText("用户登出成功")
	return
}
