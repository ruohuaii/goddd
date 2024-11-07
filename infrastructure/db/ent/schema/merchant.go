package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Merchant struct {
	ent.Schema
}

func (Merchant) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").MaxLen(16).Comment("商户名称"),
		field.Uint64("merchant_id").Comment("商户号"),
		field.String("email").MaxLen(32).Comment("商户邮箱号"),
		field.String("mobile").MaxLen(16).Comment("商户手机号"),
		field.String("salt").MaxLen(32).Comment("密码盐"),
		field.String("password").MaxLen(128).Comment("密码"),
		field.Time("create_at").Comment("创建时间"),
	}
}

func (Merchant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("merchant_id").Unique(),
	}
}
