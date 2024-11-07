// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MerchantsColumns holds the columns for the "merchants" table.
	MerchantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 16},
		{Name: "merchant_id", Type: field.TypeUint64},
		{Name: "email", Type: field.TypeString, Size: 32},
		{Name: "mobile", Type: field.TypeString, Size: 16},
		{Name: "salt", Type: field.TypeString, Size: 32},
		{Name: "password", Type: field.TypeString, Size: 128},
		{Name: "create_at", Type: field.TypeTime},
	}
	// MerchantsTable holds the schema information for the "merchants" table.
	MerchantsTable = &schema.Table{
		Name:       "merchants",
		Columns:    MerchantsColumns,
		PrimaryKey: []*schema.Column{MerchantsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "merchant_merchant_id",
				Unique:  true,
				Columns: []*schema.Column{MerchantsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MerchantsTable,
	}
)

func init() {
}
