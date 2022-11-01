package account

import "github.com/bob2325168/corehero/framework"

type AccountProvider struct {
	framework.ServiceProvider
	c framework.Container
}

func (sp *AccountProvider) Name() string {
	return AccountKey
}

func (sp *AccountProvider) Register(c framework.Container) framework.NewInstance {
	return NewAccountService
}

func (sp *AccountProvider) IsDefer() bool {
	return true
}

func (sp *AccountProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *AccountProvider) Boot(c framework.Container) error {
	return nil
}
