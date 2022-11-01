package dto

import (
	"github.com/bob2325168/bbs/framework/go-admin/common/models"
	"github.com/gin-gonic/gin"
)

type Index interface {
	Generate() Index
	Bind(ctx *gin.Context) error
	GetPageIndex() int
	GetPageSize() int
	GetNeedSearch() interface{}
}

type Control interface {
	Generate() Control
	Bind(ctx *gin.Context) error
	GenerateM() (models.ActiveRecord, error)
	GetId() interface{}
}
