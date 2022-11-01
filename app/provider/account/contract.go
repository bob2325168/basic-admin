package account

import (
	"context"
	"encoding/json"
	"time"
)

const AccountKey = "account"

// Service 用户相关的服务
type Service interface {
	// AddAccount 添加用户
	AddAccount(ctx context.Context, user *User) (*User, error)
	// EditAccount 编辑账户
	EditAccount(ctx context.Context, userId int64) (*User, error)
	// DelAccount 删除账户
	DelAccount(ctx context.Context, userId int64) error
	// GetAccountInfo 获取账户信息
	GetAccountInfo(ctx context.Context, userId int64) (*User, error)
	// GetAccouts 获取账户列表
	GetAccouts(ctx context.Context) ([]User, error)
}

// User 代表一个用户，注意这里的用户信息字段在不同接口和参数可能为空
type User struct {
	ID        int64     `gorm:"column:id;primaryKey"` // 代表用户id, 只有注册成功之后才有这个id，唯一表示一个用户
	UserName  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Token string `gorm:"-"` // token 可以用作注册token或者登录token
}

type Wallet struct {
	ID        int64     `gorm:"column:id;primaryKey"` // 代表用户id, 只有注册成功之后才有这个id，唯一表示一个用户
	Address   string    `gorm:"column:address"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Token string `gorm:"-"` // token 可以用作注册token或者登录token
}

// MarshalBinary 实现BinaryMarshaler 接口
func (b *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

// UnmarshalBinary 实现 BinaryUnMarshaler 接口
func (b *User) UnmarshalBinary(bt []byte) error {
	return json.Unmarshal(bt, b)
}
