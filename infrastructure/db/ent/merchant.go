// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/ruohuaii/goddd/infrastructure/db/ent/merchant"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Merchant is the model entity for the Merchant schema.
type Merchant struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// 商户名称
	Name string `json:"name,omitempty"`
	// 商户号
	MerchantID uint64 `json:"merchant_id,omitempty"`
	// 商户邮箱号
	Email string `json:"email,omitempty"`
	// 商户手机号
	Mobile string `json:"mobile,omitempty"`
	// 密码盐
	Salt string `json:"salt,omitempty"`
	// 密码
	Password string `json:"password,omitempty"`
	// 创建时间
	CreateAt     time.Time `json:"create_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Merchant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case merchant.FieldID, merchant.FieldMerchantID:
			values[i] = new(sql.NullInt64)
		case merchant.FieldName, merchant.FieldEmail, merchant.FieldMobile, merchant.FieldSalt, merchant.FieldPassword:
			values[i] = new(sql.NullString)
		case merchant.FieldCreateAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Merchant fields.
func (m *Merchant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case merchant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = uint64(value.Int64)
		case merchant.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case merchant.FieldMerchantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_id", values[i])
			} else if value.Valid {
				m.MerchantID = uint64(value.Int64)
			}
		case merchant.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				m.Email = value.String
			}
		case merchant.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				m.Mobile = value.String
			}
		case merchant.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				m.Salt = value.String
			}
		case merchant.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				m.Password = value.String
			}
		case merchant.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				m.CreateAt = value.Time
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Merchant.
// This includes values selected through modifiers, order, etc.
func (m *Merchant) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Merchant.
// Note that you need to call Merchant.Unwrap() before calling this method if this Merchant
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Merchant) Update() *MerchantUpdateOne {
	return NewMerchantClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Merchant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Merchant) Unwrap() *Merchant {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Merchant is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Merchant) String() string {
	var builder strings.Builder
	builder.WriteString("Merchant(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("merchant_id=")
	builder.WriteString(fmt.Sprintf("%v", m.MerchantID))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(m.Email)
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(m.Mobile)
	builder.WriteString(", ")
	builder.WriteString("salt=")
	builder.WriteString(m.Salt)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(m.Password)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(m.CreateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Merchants is a parsable slice of Merchant.
type Merchants []*Merchant
