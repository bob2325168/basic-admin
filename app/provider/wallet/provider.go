package wallet

import "github.com/bob2325168/corehero/framework"

type WalletProvider struct {
	framework.ServiceProvider
	c framework.Container
}

func (sp *WalletProvider) Name() string {
	return WalletKey
}

func (sp *WalletProvider) Register(c framework.Container) framework.NewInstance {
	return NewWalletService
}

func (sp *WalletProvider) IsDefer() bool {
	return true
}

func (sp *WalletProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *WalletProvider) Boot(c framework.Container) error {
	return nil
}
