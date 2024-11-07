// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/ruohuaii/goddd/infrastructure/db/ent/merchant"
	"github.com/ruohuaii/goddd/infrastructure/db/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMerchant = "Merchant"
)

// MerchantMutation represents an operation that mutates the Merchant nodes in the graph.
type MerchantMutation struct {
	config
	op             Op
	typ            string
	id             *uint64
	name           *string
	merchant_id    *uint64
	addmerchant_id *int64
	email          *string
	mobile         *string
	salt           *string
	password       *string
	create_at      *time.Time
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Merchant, error)
	predicates     []predicate.Merchant
}

var _ ent.Mutation = (*MerchantMutation)(nil)

// merchantOption allows management of the mutation configuration using functional options.
type merchantOption func(*MerchantMutation)

// newMerchantMutation creates new mutation for the Merchant entity.
func newMerchantMutation(c config, op Op, opts ...merchantOption) *MerchantMutation {
	m := &MerchantMutation{
		config:        c,
		op:            op,
		typ:           TypeMerchant,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMerchantID sets the ID field of the mutation.
func withMerchantID(id uint64) merchantOption {
	return func(m *MerchantMutation) {
		var (
			err   error
			once  sync.Once
			value *Merchant
		)
		m.oldValue = func(ctx context.Context) (*Merchant, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Merchant.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMerchant sets the old Merchant of the mutation.
func withMerchant(node *Merchant) merchantOption {
	return func(m *MerchantMutation) {
		m.oldValue = func(context.Context) (*Merchant, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MerchantMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MerchantMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Merchant entities.
func (m *MerchantMutation) SetID(id uint64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MerchantMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MerchantMutation) IDs(ctx context.Context) ([]uint64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Merchant.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *MerchantMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *MerchantMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Merchant entity.
// If the Merchant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MerchantMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *MerchantMutation) ResetName() {
	m.name = nil
}

// SetMerchantID sets the "merchant_id" field.
func (m *MerchantMutation) SetMerchantID(u uint64) {
	m.merchant_id = &u
	m.addmerchant_id = nil
}

// MerchantID returns the value of the "merchant_id" field in the mutation.
func (m *MerchantMutation) MerchantID() (r uint64, exists bool) {
	v := m.merchant_id
	if v == nil {
		return
	}
	return *v, true
}

// OldMerchantID returns the old "merchant_id" field's value of the Merchant entity.
// If the Merchant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MerchantMutation) OldMerchantID(ctx context.Context) (v uint64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMerchantID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMerchantID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMerchantID: %w", err)
	}
	return oldValue.MerchantID, nil
}

// AddMerchantID adds u to the "merchant_id" field.
func (m *MerchantMutation) AddMerchantID(u int64) {
	if m.addmerchant_id != nil {
		*m.addmerchant_id += u
	} else {
		m.addmerchant_id = &u
	}
}

// AddedMerchantID returns the value that was added to the "merchant_id" field in this mutation.
func (m *MerchantMutation) AddedMerchantID() (r int64, exists bool) {
	v := m.addmerchant_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetMerchantID resets all changes to the "merchant_id" field.
func (m *MerchantMutation) ResetMerchantID() {
	m.merchant_id = nil
	m.addmerchant_id = nil
}

// SetEmail sets the "email" field.
func (m *MerchantMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *MerchantMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Merchant entity.
// If the Merchant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MerchantMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *MerchantMutation) ResetEmail() {
	m.email = nil
}

// SetMobile sets the "mobile" field.
func (m *MerchantMutation) SetMobile(s string) {
	m.mobile = &s
}

// Mobile returns the value of the "mobile" field in the mutation.
func (m *MerchantMutation) Mobile() (r string, exists bool) {
	v := m.mobile
	if v == nil {
		return
	}
	return *v, true
}

// OldMobile returns the old "mobile" field's value of the Merchant entity.
// If the Merchant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MerchantMutation) OldMobile(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMobile is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMobile requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMobile: %w", err)
	}
	return oldValue.Mobile, nil
}

// ResetMobile resets all changes to the "mobile" field.
func (m *MerchantMutation) ResetMobile() {
	m.mobile = nil
}

// SetSalt sets the "salt" field.
func (m *MerchantMutation) SetSalt(s string) {
	m.salt = &s
}

// Salt returns the value of the "salt" field in the mutation.
func (m *MerchantMutation) Salt() (r string, exists bool) {
	v := m.salt
	if v == nil {
		return
	}
	return *v, true
}

// OldSalt returns the old "salt" field's value of the Merchant entity.
// If the Merchant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MerchantMutation) OldSalt(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSalt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSalt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSalt: %w", err)
	}
	return oldValue.Salt, nil
}

// ResetSalt resets all changes to the "salt" field.
func (m *MerchantMutation) ResetSalt() {
	m.salt = nil
}

// SetPassword sets the "password" field.
func (m *MerchantMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *MerchantMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the Merchant entity.
// If the Merchant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MerchantMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *MerchantMutation) ResetPassword() {
	m.password = nil
}

// SetCreateAt sets the "create_at" field.
func (m *MerchantMutation) SetCreateAt(t time.Time) {
	m.create_at = &t
}

// CreateAt returns the value of the "create_at" field in the mutation.
func (m *MerchantMutation) CreateAt() (r time.Time, exists bool) {
	v := m.create_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateAt returns the old "create_at" field's value of the Merchant entity.
// If the Merchant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MerchantMutation) OldCreateAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateAt: %w", err)
	}
	return oldValue.CreateAt, nil
}

// ResetCreateAt resets all changes to the "create_at" field.
func (m *MerchantMutation) ResetCreateAt() {
	m.create_at = nil
}

// Where appends a list predicates to the MerchantMutation builder.
func (m *MerchantMutation) Where(ps ...predicate.Merchant) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the MerchantMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *MerchantMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Merchant, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *MerchantMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *MerchantMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Merchant).
func (m *MerchantMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MerchantMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.name != nil {
		fields = append(fields, merchant.FieldName)
	}
	if m.merchant_id != nil {
		fields = append(fields, merchant.FieldMerchantID)
	}
	if m.email != nil {
		fields = append(fields, merchant.FieldEmail)
	}
	if m.mobile != nil {
		fields = append(fields, merchant.FieldMobile)
	}
	if m.salt != nil {
		fields = append(fields, merchant.FieldSalt)
	}
	if m.password != nil {
		fields = append(fields, merchant.FieldPassword)
	}
	if m.create_at != nil {
		fields = append(fields, merchant.FieldCreateAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MerchantMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case merchant.FieldName:
		return m.Name()
	case merchant.FieldMerchantID:
		return m.MerchantID()
	case merchant.FieldEmail:
		return m.Email()
	case merchant.FieldMobile:
		return m.Mobile()
	case merchant.FieldSalt:
		return m.Salt()
	case merchant.FieldPassword:
		return m.Password()
	case merchant.FieldCreateAt:
		return m.CreateAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MerchantMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case merchant.FieldName:
		return m.OldName(ctx)
	case merchant.FieldMerchantID:
		return m.OldMerchantID(ctx)
	case merchant.FieldEmail:
		return m.OldEmail(ctx)
	case merchant.FieldMobile:
		return m.OldMobile(ctx)
	case merchant.FieldSalt:
		return m.OldSalt(ctx)
	case merchant.FieldPassword:
		return m.OldPassword(ctx)
	case merchant.FieldCreateAt:
		return m.OldCreateAt(ctx)
	}
	return nil, fmt.Errorf("unknown Merchant field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MerchantMutation) SetField(name string, value ent.Value) error {
	switch name {
	case merchant.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case merchant.FieldMerchantID:
		v, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMerchantID(v)
		return nil
	case merchant.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case merchant.FieldMobile:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMobile(v)
		return nil
	case merchant.FieldSalt:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSalt(v)
		return nil
	case merchant.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case merchant.FieldCreateAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateAt(v)
		return nil
	}
	return fmt.Errorf("unknown Merchant field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MerchantMutation) AddedFields() []string {
	var fields []string
	if m.addmerchant_id != nil {
		fields = append(fields, merchant.FieldMerchantID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MerchantMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case merchant.FieldMerchantID:
		return m.AddedMerchantID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MerchantMutation) AddField(name string, value ent.Value) error {
	switch name {
	case merchant.FieldMerchantID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddMerchantID(v)
		return nil
	}
	return fmt.Errorf("unknown Merchant numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MerchantMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MerchantMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MerchantMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Merchant nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MerchantMutation) ResetField(name string) error {
	switch name {
	case merchant.FieldName:
		m.ResetName()
		return nil
	case merchant.FieldMerchantID:
		m.ResetMerchantID()
		return nil
	case merchant.FieldEmail:
		m.ResetEmail()
		return nil
	case merchant.FieldMobile:
		m.ResetMobile()
		return nil
	case merchant.FieldSalt:
		m.ResetSalt()
		return nil
	case merchant.FieldPassword:
		m.ResetPassword()
		return nil
	case merchant.FieldCreateAt:
		m.ResetCreateAt()
		return nil
	}
	return fmt.Errorf("unknown Merchant field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MerchantMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MerchantMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MerchantMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MerchantMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MerchantMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MerchantMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MerchantMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Merchant unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MerchantMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Merchant edge %s", name)
}
