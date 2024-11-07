package persistence

import (
	"context"

	"github.com/ruohuaii/goddd/domain/merchant"
	"github.com/ruohuaii/goddd/infrastructure/db"
	"github.com/ruohuaii/goddd/infrastructure/db/ent"
	where "github.com/ruohuaii/goddd/infrastructure/db/ent/merchant"
	"github.com/ruohuaii/goddd/infrastructure/faults"
)

type merchantRepository struct {
	client *ent.Client
}

func NewMerchantRepository() merchant.Repository {
	return &merchantRepository{client: db.Client()}
}

func (m *merchantRepository) Save(ctx context.Context, input *merchant.Merchant) (*merchant.Merchant, error) {
	if input == nil {
		return nil, faults.MerchantInfoError
	}
	record, _ := m.FindByName(ctx, input.Name)
	if record != nil {
		return record, faults.MerchantExists
	}
	data, err := m.client.Merchant.Create().
		SetMerchantID(input.MerchantID).
		SetName(input.Name).
		SetEmail(input.Email).
		SetMobile(input.Mobile).
		SetSalt(input.Salt).
		SetPassword(input.Password).
		SetCreateAt(input.CreatedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &merchant.Merchant{
		MerchantID: data.MerchantID,
		Name:       data.Name,
		Email:      data.Email,
		Mobile:     data.Mobile,
		Password:   data.Password,
		Salt:       data.Salt,
		CreatedAt:  data.CreateAt,
	}, nil
}

func (m *merchantRepository) FindByName(ctx context.Context, name string) (*merchant.Merchant, error) {
	data, _ := m.client.Merchant.Query().Where(where.Name(name)).Only(ctx)
	if data == nil {
		return nil, faults.MerchantNotFound
	}
	return &merchant.Merchant{
		MerchantID: data.MerchantID,
		Name:       data.Name,
		Email:      data.Email,
		Mobile:     data.Mobile,
		Password:   data.Password,
		Salt:       data.Salt,
		CreatedAt:  data.CreateAt,
	}, nil
}
