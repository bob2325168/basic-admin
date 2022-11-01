package wallet

import (
	"github.com/bob2325168/bbs/app/provider/wallet"
	"github.com/bob2325168/corehero/framework/gin"
)

type WalletApi struct{}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	api := &WalletApi{}
	if !r.IsBind(wallet.WalletKey) {
		r.Bind(&wallet.WalletProvider{})
	}
	r.Group("/api/v1/")
	{
		// 钱包信息
		r.GET("/user/addressinfo", api.GetAddressInfo)
	}
	return nil
}
