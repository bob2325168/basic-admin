package account

import (
	"context"
	"fmt"
	"github.com/bob2325168/corehero/framework"
	"github.com/bob2325168/corehero/framework/contract"
	"gorm.io/gorm"
)

type AccountService struct {
	container framework.Container
	logger    contract.Log
	ormDB     *gorm.DB // db
}

func NewAccountService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	logger := container.MustMake(contract.LogKey).(contract.Log)
	ormService := container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		logger.Error(context.Background(), "获取gormDB错误", map[string]interface{}{
			"err": fmt.Sprintf("%+v", err),
		})
		return nil, err
	}
	return &AccountService{container: container, logger: logger, ormDB: db}, nil
}

func (as *AccountService) AddAccount(ctx context.Context, user *User) (*User, error) {
	return nil, nil
}

func (as *AccountService) EditAccount(ctx context.Context, userId int64) (*User, error) {
	return nil, nil
}

func (as *AccountService) DelAccount(ctx context.Context, userId int64) error {
	return nil
}

func (as *AccountService) GetAccountInfo(ctx context.Context, userId int64) (*User, error) {
	return nil, nil
}

func (as *AccountService) GetAccouts(ctx context.Context) ([]User, error) {
	return nil, nil
}
