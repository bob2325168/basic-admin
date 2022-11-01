package role

import (
	"encoding/json"
	"time"
)

const RoleKey = "role"

// Service 用户相关的服务
type Service interface {
}

// Role 角色以及权限
type Role struct {
	ID        int64     `gorm:"column:id;primaryKey"` // 代表用户id, 只有注册成功之后才有这个id，唯一表示一个用户
	UserName  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Token string `gorm:"-"` // token 可以用作注册token或者登录token
}

// MarshalBinary 实现BinaryMarshaler 接口
func (b *Role) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

// UnmarshalBinary 实现 BinaryUnMarshaler 接口
func (b *Role) UnmarshalBinary(bt []byte) error {
	return json.Unmarshal(bt, b)
}
