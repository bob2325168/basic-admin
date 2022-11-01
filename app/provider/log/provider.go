package log

import "github.com/bob2325168/corehero/framework"

type OperaLogProvider struct {
	framework.ServiceProvider
	c framework.Container
}

func (sp *OperaLogProvider) Name() string {
	return OperaLogKey
}

func (sp *OperaLogProvider) Register(c framework.Container) framework.NewInstance {
	return NewOperaLogService
}

func (sp *OperaLogProvider) IsDefer() bool {
	return true
}

func (sp *OperaLogProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *OperaLogProvider) Boot(c framework.Container) error {
	return nil
}
