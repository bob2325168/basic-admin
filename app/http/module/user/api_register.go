package user

import (
	"fmt"
	provider "github.com/bob2325168/bbs/app/provider/user"
	"github.com/bob2325168/corehero/framework/contract"
	"github.com/bob2325168/corehero/framework/gin"
	"time"
)

type registerParam struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,gte=6"`
	Email    string `json:"email" binding:"required,gte=6"`
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Accept  json
// @Produce  json
// @Tags 用户操作
// @Param registerParam body registerParam true "注册参数"
// @Success 200 {string} Message "注册成功"
// @Router /api/v1/user/register [post]
func (api *UserApi) Register(c *gin.Context) {
	// 验证参数
	userService := c.MustMake(provider.UserKey).(provider.Service)
	logger := c.MustMake(contract.LogKey).(contract.Log)

	param := &registerParam{}
	// gin框架自带的验证参数
	if err := c.ShouldBind(param); err != nil {
		c.ISetStatus(404).IText("参数错误")
		return
	}

	model := &provider.User{
		UserName:  param.UserName,
		Password:  param.Password,
		Email:     param.Email,
		CreatedAt: time.Now(),
	}
	// 注册
	userWithToken, err := userService.Register(c, model)
	if err != nil {
		logger.Error(c, err.Error(), map[string]interface{}{
			"stack": fmt.Sprintf("%+v", err),
		})
		c.ISetStatus(500).IText(err.Error())
		return
	}
	if userWithToken == nil {
		c.ISetStatus(500).IText("注册失败")
		return
	}

	if err := userService.SendRegisterMail(c, userWithToken); err != nil {

		c.ISetStatus(500).IText("发送电子邮件失败")
		return
	}

	c.ISetOkStatus().IText("注册成功，请前往邮箱查看邮件")
	return
}
