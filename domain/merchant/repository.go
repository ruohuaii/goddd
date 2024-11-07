package merchant

import "context"

// Repository 仓储接口，定义数据库操作，具体实现会在持久层实现。
type Repository interface {
	Save(ctx context.Context, merchant *Merchant) (*Merchant, error)
	FindByName(ctx context.Context, name string) (*Merchant, error)
}
