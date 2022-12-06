// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent/appcountry"
	"github.com/google/uuid"
)

// AppCountryCreate is the builder for creating a AppCountry entity.
type AppCountryCreate struct {
	config
	mutation *AppCountryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (acc *AppCountryCreate) SetCreatedAt(u uint32) *AppCountryCreate {
	acc.mutation.SetCreatedAt(u)
	return acc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acc *AppCountryCreate) SetNillableCreatedAt(u *uint32) *AppCountryCreate {
	if u != nil {
		acc.SetCreatedAt(*u)
	}
	return acc
}

// SetUpdatedAt sets the "updated_at" field.
func (acc *AppCountryCreate) SetUpdatedAt(u uint32) *AppCountryCreate {
	acc.mutation.SetUpdatedAt(u)
	return acc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (acc *AppCountryCreate) SetNillableUpdatedAt(u *uint32) *AppCountryCreate {
	if u != nil {
		acc.SetUpdatedAt(*u)
	}
	return acc
}

// SetDeletedAt sets the "deleted_at" field.
func (acc *AppCountryCreate) SetDeletedAt(u uint32) *AppCountryCreate {
	acc.mutation.SetDeletedAt(u)
	return acc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (acc *AppCountryCreate) SetNillableDeletedAt(u *uint32) *AppCountryCreate {
	if u != nil {
		acc.SetDeletedAt(*u)
	}
	return acc
}

// SetAppID sets the "app_id" field.
func (acc *AppCountryCreate) SetAppID(u uuid.UUID) *AppCountryCreate {
	acc.mutation.SetAppID(u)
	return acc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (acc *AppCountryCreate) SetNillableAppID(u *uuid.UUID) *AppCountryCreate {
	if u != nil {
		acc.SetAppID(*u)
	}
	return acc
}

// SetCountryID sets the "country_id" field.
func (acc *AppCountryCreate) SetCountryID(u uuid.UUID) *AppCountryCreate {
	acc.mutation.SetCountryID(u)
	return acc
}

// SetNillableCountryID sets the "country_id" field if the given value is not nil.
func (acc *AppCountryCreate) SetNillableCountryID(u *uuid.UUID) *AppCountryCreate {
	if u != nil {
		acc.SetCountryID(*u)
	}
	return acc
}

// SetID sets the "id" field.
func (acc *AppCountryCreate) SetID(u uuid.UUID) *AppCountryCreate {
	acc.mutation.SetID(u)
	return acc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (acc *AppCountryCreate) SetNillableID(u *uuid.UUID) *AppCountryCreate {
	if u != nil {
		acc.SetID(*u)
	}
	return acc
}

// Mutation returns the AppCountryMutation object of the builder.
func (acc *AppCountryCreate) Mutation() *AppCountryMutation {
	return acc.mutation
}

// Save creates the AppCountry in the database.
func (acc *AppCountryCreate) Save(ctx context.Context) (*AppCountry, error) {
	var (
		err  error
		node *AppCountry
	)
	if err := acc.defaults(); err != nil {
		return nil, err
	}
	if len(acc.hooks) == 0 {
		if err = acc.check(); err != nil {
			return nil, err
		}
		node, err = acc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppCountryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acc.check(); err != nil {
				return nil, err
			}
			acc.mutation = mutation
			if node, err = acc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(acc.hooks) - 1; i >= 0; i-- {
			if acc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, acc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppCountry)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppCountryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (acc *AppCountryCreate) SaveX(ctx context.Context) *AppCountry {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *AppCountryCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *AppCountryCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acc *AppCountryCreate) defaults() error {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		if appcountry.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appcountry.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appcountry.DefaultCreatedAt()
		acc.mutation.SetCreatedAt(v)
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		if appcountry.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appcountry.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appcountry.DefaultUpdatedAt()
		acc.mutation.SetUpdatedAt(v)
	}
	if _, ok := acc.mutation.DeletedAt(); !ok {
		if appcountry.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appcountry.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appcountry.DefaultDeletedAt()
		acc.mutation.SetDeletedAt(v)
	}
	if _, ok := acc.mutation.AppID(); !ok {
		if appcountry.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized appcountry.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := appcountry.DefaultAppID()
		acc.mutation.SetAppID(v)
	}
	if _, ok := acc.mutation.CountryID(); !ok {
		if appcountry.DefaultCountryID == nil {
			return fmt.Errorf("ent: uninitialized appcountry.DefaultCountryID (forgotten import ent/runtime?)")
		}
		v := appcountry.DefaultCountryID()
		acc.mutation.SetCountryID(v)
	}
	if _, ok := acc.mutation.ID(); !ok {
		if appcountry.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized appcountry.DefaultID (forgotten import ent/runtime?)")
		}
		v := appcountry.DefaultID()
		acc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (acc *AppCountryCreate) check() error {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppCountry.created_at"`)}
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppCountry.updated_at"`)}
	}
	if _, ok := acc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppCountry.deleted_at"`)}
	}
	return nil
}

func (acc *AppCountryCreate) sqlSave(ctx context.Context) (*AppCountry, error) {
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (acc *AppCountryCreate) createSpec() (*AppCountry, *sqlgraph.CreateSpec) {
	var (
		_node = &AppCountry{config: acc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appcountry.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcountry.FieldID,
			},
		}
	)
	_spec.OnConflict = acc.conflict
	if id, ok := acc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := acc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := acc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := acc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appcountry.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := acc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcountry.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := acc.mutation.CountryID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appcountry.FieldCountryID,
		})
		_node.CountryID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppCountry.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppCountryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (acc *AppCountryCreate) OnConflict(opts ...sql.ConflictOption) *AppCountryUpsertOne {
	acc.conflict = opts
	return &AppCountryUpsertOne{
		create: acc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppCountry.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acc *AppCountryCreate) OnConflictColumns(columns ...string) *AppCountryUpsertOne {
	acc.conflict = append(acc.conflict, sql.ConflictColumns(columns...))
	return &AppCountryUpsertOne{
		create: acc,
	}
}

type (
	// AppCountryUpsertOne is the builder for "upsert"-ing
	//  one AppCountry node.
	AppCountryUpsertOne struct {
		create *AppCountryCreate
	}

	// AppCountryUpsert is the "OnConflict" setter.
	AppCountryUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppCountryUpsert) SetCreatedAt(v uint32) *AppCountryUpsert {
	u.Set(appcountry.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppCountryUpsert) UpdateCreatedAt() *AppCountryUpsert {
	u.SetExcluded(appcountry.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppCountryUpsert) AddCreatedAt(v uint32) *AppCountryUpsert {
	u.Add(appcountry.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppCountryUpsert) SetUpdatedAt(v uint32) *AppCountryUpsert {
	u.Set(appcountry.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppCountryUpsert) UpdateUpdatedAt() *AppCountryUpsert {
	u.SetExcluded(appcountry.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppCountryUpsert) AddUpdatedAt(v uint32) *AppCountryUpsert {
	u.Add(appcountry.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppCountryUpsert) SetDeletedAt(v uint32) *AppCountryUpsert {
	u.Set(appcountry.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppCountryUpsert) UpdateDeletedAt() *AppCountryUpsert {
	u.SetExcluded(appcountry.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppCountryUpsert) AddDeletedAt(v uint32) *AppCountryUpsert {
	u.Add(appcountry.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppCountryUpsert) SetAppID(v uuid.UUID) *AppCountryUpsert {
	u.Set(appcountry.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppCountryUpsert) UpdateAppID() *AppCountryUpsert {
	u.SetExcluded(appcountry.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppCountryUpsert) ClearAppID() *AppCountryUpsert {
	u.SetNull(appcountry.FieldAppID)
	return u
}

// SetCountryID sets the "country_id" field.
func (u *AppCountryUpsert) SetCountryID(v uuid.UUID) *AppCountryUpsert {
	u.Set(appcountry.FieldCountryID, v)
	return u
}

// UpdateCountryID sets the "country_id" field to the value that was provided on create.
func (u *AppCountryUpsert) UpdateCountryID() *AppCountryUpsert {
	u.SetExcluded(appcountry.FieldCountryID)
	return u
}

// ClearCountryID clears the value of the "country_id" field.
func (u *AppCountryUpsert) ClearCountryID() *AppCountryUpsert {
	u.SetNull(appcountry.FieldCountryID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppCountry.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appcountry.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppCountryUpsertOne) UpdateNewValues() *AppCountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appcountry.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppCountry.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppCountryUpsertOne) Ignore() *AppCountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppCountryUpsertOne) DoNothing() *AppCountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCountryCreate.OnConflict
// documentation for more info.
func (u *AppCountryUpsertOne) Update(set func(*AppCountryUpsert)) *AppCountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppCountryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppCountryUpsertOne) SetCreatedAt(v uint32) *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppCountryUpsertOne) AddCreatedAt(v uint32) *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppCountryUpsertOne) UpdateCreatedAt() *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppCountryUpsertOne) SetUpdatedAt(v uint32) *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppCountryUpsertOne) AddUpdatedAt(v uint32) *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppCountryUpsertOne) UpdateUpdatedAt() *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppCountryUpsertOne) SetDeletedAt(v uint32) *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppCountryUpsertOne) AddDeletedAt(v uint32) *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppCountryUpsertOne) UpdateDeletedAt() *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppCountryUpsertOne) SetAppID(v uuid.UUID) *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppCountryUpsertOne) UpdateAppID() *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppCountryUpsertOne) ClearAppID() *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.ClearAppID()
	})
}

// SetCountryID sets the "country_id" field.
func (u *AppCountryUpsertOne) SetCountryID(v uuid.UUID) *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.SetCountryID(v)
	})
}

// UpdateCountryID sets the "country_id" field to the value that was provided on create.
func (u *AppCountryUpsertOne) UpdateCountryID() *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.UpdateCountryID()
	})
}

// ClearCountryID clears the value of the "country_id" field.
func (u *AppCountryUpsertOne) ClearCountryID() *AppCountryUpsertOne {
	return u.Update(func(s *AppCountryUpsert) {
		s.ClearCountryID()
	})
}

// Exec executes the query.
func (u *AppCountryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppCountryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppCountryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppCountryUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppCountryUpsertOne.ID is not supported by MySQL driver. Use AppCountryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppCountryUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppCountryCreateBulk is the builder for creating many AppCountry entities in bulk.
type AppCountryCreateBulk struct {
	config
	builders []*AppCountryCreate
	conflict []sql.ConflictOption
}

// Save creates the AppCountry entities in the database.
func (accb *AppCountryCreateBulk) Save(ctx context.Context) ([]*AppCountry, error) {
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*AppCountry, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppCountryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = accb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *AppCountryCreateBulk) SaveX(ctx context.Context) []*AppCountry {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *AppCountryCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *AppCountryCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppCountry.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppCountryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (accb *AppCountryCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppCountryUpsertBulk {
	accb.conflict = opts
	return &AppCountryUpsertBulk{
		create: accb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppCountry.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (accb *AppCountryCreateBulk) OnConflictColumns(columns ...string) *AppCountryUpsertBulk {
	accb.conflict = append(accb.conflict, sql.ConflictColumns(columns...))
	return &AppCountryUpsertBulk{
		create: accb,
	}
}

// AppCountryUpsertBulk is the builder for "upsert"-ing
// a bulk of AppCountry nodes.
type AppCountryUpsertBulk struct {
	create *AppCountryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppCountry.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appcountry.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppCountryUpsertBulk) UpdateNewValues() *AppCountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appcountry.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppCountry.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppCountryUpsertBulk) Ignore() *AppCountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppCountryUpsertBulk) DoNothing() *AppCountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCountryCreateBulk.OnConflict
// documentation for more info.
func (u *AppCountryUpsertBulk) Update(set func(*AppCountryUpsert)) *AppCountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppCountryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppCountryUpsertBulk) SetCreatedAt(v uint32) *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppCountryUpsertBulk) AddCreatedAt(v uint32) *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppCountryUpsertBulk) UpdateCreatedAt() *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppCountryUpsertBulk) SetUpdatedAt(v uint32) *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppCountryUpsertBulk) AddUpdatedAt(v uint32) *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppCountryUpsertBulk) UpdateUpdatedAt() *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppCountryUpsertBulk) SetDeletedAt(v uint32) *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppCountryUpsertBulk) AddDeletedAt(v uint32) *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppCountryUpsertBulk) UpdateDeletedAt() *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppCountryUpsertBulk) SetAppID(v uuid.UUID) *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppCountryUpsertBulk) UpdateAppID() *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppCountryUpsertBulk) ClearAppID() *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.ClearAppID()
	})
}

// SetCountryID sets the "country_id" field.
func (u *AppCountryUpsertBulk) SetCountryID(v uuid.UUID) *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.SetCountryID(v)
	})
}

// UpdateCountryID sets the "country_id" field to the value that was provided on create.
func (u *AppCountryUpsertBulk) UpdateCountryID() *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.UpdateCountryID()
	})
}

// ClearCountryID clears the value of the "country_id" field.
func (u *AppCountryUpsertBulk) ClearCountryID() *AppCountryUpsertBulk {
	return u.Update(func(s *AppCountryUpsert) {
		s.ClearCountryID()
	})
}

// Exec executes the query.
func (u *AppCountryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppCountryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppCountryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppCountryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
