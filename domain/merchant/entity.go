package merchant

import "time"

// Merchant 领域实体，包含业务属性和行为。
type Merchant struct {
	MerchantID uint64
	Name       string
	Email      string
	Mobile     string
	Password   string
	Salt       string
	CreatedAt  time.Time
}

func (m *Merchant) Check() error {
	return nil
}
