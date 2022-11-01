package api

import (
	"context"
	"fmt"
	"github.com/bob2325168/corehero/framework"
	"github.com/bob2325168/corehero/framework/contract"
	"gorm.io/gorm"
)

type ApiService struct {
	container framework.Container
	logger    contract.Log
	ormDB     *gorm.DB // db
}

func NewApiService(params ...interface{}) (interface{}, error) {
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
	return &ApiService{container: container, logger: logger, ormDB: db}, nil
}
