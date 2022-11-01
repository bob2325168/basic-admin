package api

import "github.com/bob2325168/corehero/framework"

type ApiProvider struct {
	framework.ServiceProvider
	c framework.Container
}

func (sp *ApiProvider) Name() string {
	return ApiKey
}

func (sp *ApiProvider) Register(c framework.Container) framework.NewInstance {
	return NewApiService
}

func (sp *ApiProvider) IsDefer() bool {
	return true
}

func (sp *ApiProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *ApiProvider) Boot(c framework.Container) error {
	return nil
}
