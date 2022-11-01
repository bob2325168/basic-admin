package wallet

import (
	"context"
	"encoding/json"
	"time"
)

const WalletKey = "wallet"

// Service 钱包相关的服务
type Service interface {
	// GetWallet 获取钱包信息
	GetWallet(ctx context.Context, addr string) (*Wallet, error)
}

type Wallet struct {
	ID        int64     `gorm:"column:id;primaryKey"` // 代表用户的钱包id
	Address   string    `gorm:"column:address"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Token string `gorm:"-"` // token 可以用作注册token或者登录token
}

// MarshalBinary 实现BinaryMarshaler 接口
func (w *Wallet) MarshalBinary() ([]byte, error) {
	return json.Marshal(w)
}

// UnmarshalBinary 实现 BinaryUnMarshaler 接口
func (w *Wallet) UnmarshalBinary(bt []byte) error {
	return json.Unmarshal(bt, w)
}
