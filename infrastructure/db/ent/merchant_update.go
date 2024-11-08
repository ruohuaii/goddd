// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/ruohuaii/goddd/infrastructure/db/ent/merchant"
	"github.com/ruohuaii/goddd/infrastructure/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MerchantUpdate is the builder for updating Merchant entities.
type MerchantUpdate struct {
	config
	hooks    []Hook
	mutation *MerchantMutation
}

// Where appends a list predicates to the MerchantUpdate builder.
func (mu *MerchantUpdate) Where(ps ...predicate.Merchant) *MerchantUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *MerchantUpdate) SetName(s string) *MerchantUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableName(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetMerchantID sets the "merchant_id" field.
func (mu *MerchantUpdate) SetMerchantID(u uint64) *MerchantUpdate {
	mu.mutation.ResetMerchantID()
	mu.mutation.SetMerchantID(u)
	return mu
}

// SetNillableMerchantID sets the "merchant_id" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableMerchantID(u *uint64) *MerchantUpdate {
	if u != nil {
		mu.SetMerchantID(*u)
	}
	return mu
}

// AddMerchantID adds u to the "merchant_id" field.
func (mu *MerchantUpdate) AddMerchantID(u int64) *MerchantUpdate {
	mu.mutation.AddMerchantID(u)
	return mu
}

// SetEmail sets the "email" field.
func (mu *MerchantUpdate) SetEmail(s string) *MerchantUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableEmail(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetEmail(*s)
	}
	return mu
}

// SetMobile sets the "mobile" field.
func (mu *MerchantUpdate) SetMobile(s string) *MerchantUpdate {
	mu.mutation.SetMobile(s)
	return mu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableMobile(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetMobile(*s)
	}
	return mu
}

// SetSalt sets the "salt" field.
func (mu *MerchantUpdate) SetSalt(s string) *MerchantUpdate {
	mu.mutation.SetSalt(s)
	return mu
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableSalt(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetSalt(*s)
	}
	return mu
}

// SetPassword sets the "password" field.
func (mu *MerchantUpdate) SetPassword(s string) *MerchantUpdate {
	mu.mutation.SetPassword(s)
	return mu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillablePassword(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetPassword(*s)
	}
	return mu
}

// SetCreateAt sets the "create_at" field.
func (mu *MerchantUpdate) SetCreateAt(t time.Time) *MerchantUpdate {
	mu.mutation.SetCreateAt(t)
	return mu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableCreateAt(t *time.Time) *MerchantUpdate {
	if t != nil {
		mu.SetCreateAt(*t)
	}
	return mu
}

// Mutation returns the MerchantMutation object of the builder.
func (mu *MerchantUpdate) Mutation() *MerchantMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MerchantUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MerchantUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MerchantUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MerchantUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MerchantUpdate) check() error {
	if v, ok := mu.mutation.Name(); ok {
		if err := merchant.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Merchant.name": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Email(); ok {
		if err := merchant.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Merchant.email": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Mobile(); ok {
		if err := merchant.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "Merchant.mobile": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Salt(); ok {
		if err := merchant.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "Merchant.salt": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Password(); ok {
		if err := merchant.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Merchant.password": %w`, err)}
		}
	}
	return nil
}

func (mu *MerchantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(merchant.Table, merchant.Columns, sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(merchant.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.MerchantID(); ok {
		_spec.SetField(merchant.FieldMerchantID, field.TypeUint64, value)
	}
	if value, ok := mu.mutation.AddedMerchantID(); ok {
		_spec.AddField(merchant.FieldMerchantID, field.TypeUint64, value)
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.SetField(merchant.FieldEmail, field.TypeString, value)
	}
	if value, ok := mu.mutation.Mobile(); ok {
		_spec.SetField(merchant.FieldMobile, field.TypeString, value)
	}
	if value, ok := mu.mutation.Salt(); ok {
		_spec.SetField(merchant.FieldSalt, field.TypeString, value)
	}
	if value, ok := mu.mutation.Password(); ok {
		_spec.SetField(merchant.FieldPassword, field.TypeString, value)
	}
	if value, ok := mu.mutation.CreateAt(); ok {
		_spec.SetField(merchant.FieldCreateAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MerchantUpdateOne is the builder for updating a single Merchant entity.
type MerchantUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MerchantMutation
}

// SetName sets the "name" field.
func (muo *MerchantUpdateOne) SetName(s string) *MerchantUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableName(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetMerchantID sets the "merchant_id" field.
func (muo *MerchantUpdateOne) SetMerchantID(u uint64) *MerchantUpdateOne {
	muo.mutation.ResetMerchantID()
	muo.mutation.SetMerchantID(u)
	return muo
}

// SetNillableMerchantID sets the "merchant_id" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableMerchantID(u *uint64) *MerchantUpdateOne {
	if u != nil {
		muo.SetMerchantID(*u)
	}
	return muo
}

// AddMerchantID adds u to the "merchant_id" field.
func (muo *MerchantUpdateOne) AddMerchantID(u int64) *MerchantUpdateOne {
	muo.mutation.AddMerchantID(u)
	return muo
}

// SetEmail sets the "email" field.
func (muo *MerchantUpdateOne) SetEmail(s string) *MerchantUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableEmail(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetEmail(*s)
	}
	return muo
}

// SetMobile sets the "mobile" field.
func (muo *MerchantUpdateOne) SetMobile(s string) *MerchantUpdateOne {
	muo.mutation.SetMobile(s)
	return muo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableMobile(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetMobile(*s)
	}
	return muo
}

// SetSalt sets the "salt" field.
func (muo *MerchantUpdateOne) SetSalt(s string) *MerchantUpdateOne {
	muo.mutation.SetSalt(s)
	return muo
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableSalt(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetSalt(*s)
	}
	return muo
}

// SetPassword sets the "password" field.
func (muo *MerchantUpdateOne) SetPassword(s string) *MerchantUpdateOne {
	muo.mutation.SetPassword(s)
	return muo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillablePassword(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetPassword(*s)
	}
	return muo
}

// SetCreateAt sets the "create_at" field.
func (muo *MerchantUpdateOne) SetCreateAt(t time.Time) *MerchantUpdateOne {
	muo.mutation.SetCreateAt(t)
	return muo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableCreateAt(t *time.Time) *MerchantUpdateOne {
	if t != nil {
		muo.SetCreateAt(*t)
	}
	return muo
}

// Mutation returns the MerchantMutation object of the builder.
func (muo *MerchantUpdateOne) Mutation() *MerchantMutation {
	return muo.mutation
}

// Where appends a list predicates to the MerchantUpdate builder.
func (muo *MerchantUpdateOne) Where(ps ...predicate.Merchant) *MerchantUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MerchantUpdateOne) Select(field string, fields ...string) *MerchantUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Merchant entity.
func (muo *MerchantUpdateOne) Save(ctx context.Context) (*Merchant, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MerchantUpdateOne) SaveX(ctx context.Context) *Merchant {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MerchantUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MerchantUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MerchantUpdateOne) check() error {
	if v, ok := muo.mutation.Name(); ok {
		if err := merchant.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Merchant.name": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Email(); ok {
		if err := merchant.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Merchant.email": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Mobile(); ok {
		if err := merchant.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "Merchant.mobile": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Salt(); ok {
		if err := merchant.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "Merchant.salt": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Password(); ok {
		if err := merchant.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Merchant.password": %w`, err)}
		}
	}
	return nil
}

func (muo *MerchantUpdateOne) sqlSave(ctx context.Context) (_node *Merchant, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(merchant.Table, merchant.Columns, sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Merchant.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchant.FieldID)
		for _, f := range fields {
			if !merchant.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != merchant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(merchant.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.MerchantID(); ok {
		_spec.SetField(merchant.FieldMerchantID, field.TypeUint64, value)
	}
	if value, ok := muo.mutation.AddedMerchantID(); ok {
		_spec.AddField(merchant.FieldMerchantID, field.TypeUint64, value)
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.SetField(merchant.FieldEmail, field.TypeString, value)
	}
	if value, ok := muo.mutation.Mobile(); ok {
		_spec.SetField(merchant.FieldMobile, field.TypeString, value)
	}
	if value, ok := muo.mutation.Salt(); ok {
		_spec.SetField(merchant.FieldSalt, field.TypeString, value)
	}
	if value, ok := muo.mutation.Password(); ok {
		_spec.SetField(merchant.FieldPassword, field.TypeString, value)
	}
	if value, ok := muo.mutation.CreateAt(); ok {
		_spec.SetField(merchant.FieldCreateAt, field.TypeTime, value)
	}
	_node = &Merchant{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
