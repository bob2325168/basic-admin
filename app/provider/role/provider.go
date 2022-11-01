package role

import "github.com/bob2325168/corehero/framework"

type RoleProvider struct {
	framework.ServiceProvider
	c framework.Container
}

func (sp *RoleProvider) Name() string {
	return RoleKey
}

func (sp *RoleProvider) Register(c framework.Container) framework.NewInstance {
	return NewRoleService
}

func (sp *RoleProvider) IsDefer() bool {
	return true
}

func (sp *RoleProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *RoleProvider) Boot(c framework.Container) error {
	return nil
}
