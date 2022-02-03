// Code generated by entc, DO NOT EDIT.

package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Z00mZE/url-shortner/ent/service/shorturl"
)

// ShortUrlCreate is the builder for creating a ShortUrl entity.
type ShortUrlCreate struct {
	config
	mutation *ShortUrlMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetURL sets the "url" field.
func (suc *ShortUrlCreate) SetURL(s string) *ShortUrlCreate {
	suc.mutation.SetURL(s)
	return suc
}

// SetExpiredAt sets the "expired_at" field.
func (suc *ShortUrlCreate) SetExpiredAt(t time.Time) *ShortUrlCreate {
	suc.mutation.SetExpiredAt(t)
	return suc
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (suc *ShortUrlCreate) SetNillableExpiredAt(t *time.Time) *ShortUrlCreate {
	if t != nil {
		suc.SetExpiredAt(*t)
	}
	return suc
}

// Mutation returns the ShortUrlMutation object of the builder.
func (suc *ShortUrlCreate) Mutation() *ShortUrlMutation {
	return suc.mutation
}

// Save creates the ShortUrl in the database.
func (suc *ShortUrlCreate) Save(ctx context.Context) (*ShortUrl, error) {
	var (
		err  error
		node *ShortUrl
	)
	suc.defaults()
	if len(suc.hooks) == 0 {
		if err = suc.check(); err != nil {
			return nil, err
		}
		node, err = suc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShortUrlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suc.check(); err != nil {
				return nil, err
			}
			suc.mutation = mutation
			if node, err = suc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(suc.hooks) - 1; i >= 0; i-- {
			if suc.hooks[i] == nil {
				return nil, fmt.Errorf("service: uninitialized hook (forgotten import service/runtime?)")
			}
			mut = suc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (suc *ShortUrlCreate) SaveX(ctx context.Context) *ShortUrl {
	v, err := suc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (suc *ShortUrlCreate) Exec(ctx context.Context) error {
	_, err := suc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suc *ShortUrlCreate) ExecX(ctx context.Context) {
	if err := suc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suc *ShortUrlCreate) defaults() {
	if _, ok := suc.mutation.ExpiredAt(); !ok {
		v := shorturl.DefaultExpiredAt()
		suc.mutation.SetExpiredAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suc *ShortUrlCreate) check() error {
	if _, ok := suc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`service: missing required field "ShortUrl.url"`)}
	}
	if v, ok := suc.mutation.URL(); ok {
		if err := shorturl.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`service: validator failed for field "ShortUrl.url": %w`, err)}
		}
	}
	if _, ok := suc.mutation.ExpiredAt(); !ok {
		return &ValidationError{Name: "expired_at", err: errors.New(`service: missing required field "ShortUrl.expired_at"`)}
	}
	return nil
}

func (suc *ShortUrlCreate) sqlSave(ctx context.Context) (*ShortUrl, error) {
	_node, _spec := suc.createSpec()
	if err := sqlgraph.CreateNode(ctx, suc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (suc *ShortUrlCreate) createSpec() (*ShortUrl, *sqlgraph.CreateSpec) {
	var (
		_node = &ShortUrl{config: suc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shorturl.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shorturl.FieldID,
			},
		}
	)
	_spec.OnConflict = suc.conflict
	if value, ok := suc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shorturl.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := suc.mutation.ExpiredAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shorturl.FieldExpiredAt,
		})
		_node.ExpiredAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ShortUrl.Create().
//		SetURL(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ShortUrlUpsert) {
//			SetURL(v+v).
//		}).
//		Exec(ctx)
//
func (suc *ShortUrlCreate) OnConflict(opts ...sql.ConflictOption) *ShortUrlUpsertOne {
	suc.conflict = opts
	return &ShortUrlUpsertOne{
		create: suc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ShortUrl.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (suc *ShortUrlCreate) OnConflictColumns(columns ...string) *ShortUrlUpsertOne {
	suc.conflict = append(suc.conflict, sql.ConflictColumns(columns...))
	return &ShortUrlUpsertOne{
		create: suc,
	}
}

type (
	// ShortUrlUpsertOne is the builder for "upsert"-ing
	//  one ShortUrl node.
	ShortUrlUpsertOne struct {
		create *ShortUrlCreate
	}

	// ShortUrlUpsert is the "OnConflict" setter.
	ShortUrlUpsert struct {
		*sql.UpdateSet
	}
)

// SetURL sets the "url" field.
func (u *ShortUrlUpsert) SetURL(v string) *ShortUrlUpsert {
	u.Set(shorturl.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ShortUrlUpsert) UpdateURL() *ShortUrlUpsert {
	u.SetExcluded(shorturl.FieldURL)
	return u
}

// SetExpiredAt sets the "expired_at" field.
func (u *ShortUrlUpsert) SetExpiredAt(v time.Time) *ShortUrlUpsert {
	u.Set(shorturl.FieldExpiredAt, v)
	return u
}

// UpdateExpiredAt sets the "expired_at" field to the value that was provided on create.
func (u *ShortUrlUpsert) UpdateExpiredAt() *ShortUrlUpsert {
	u.SetExcluded(shorturl.FieldExpiredAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.ShortUrl.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ShortUrlUpsertOne) UpdateNewValues() *ShortUrlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ShortUrl.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ShortUrlUpsertOne) Ignore() *ShortUrlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ShortUrlUpsertOne) DoNothing() *ShortUrlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ShortUrlCreate.OnConflict
// documentation for more info.
func (u *ShortUrlUpsertOne) Update(set func(*ShortUrlUpsert)) *ShortUrlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ShortUrlUpsert{UpdateSet: update})
	}))
	return u
}

// SetURL sets the "url" field.
func (u *ShortUrlUpsertOne) SetURL(v string) *ShortUrlUpsertOne {
	return u.Update(func(s *ShortUrlUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ShortUrlUpsertOne) UpdateURL() *ShortUrlUpsertOne {
	return u.Update(func(s *ShortUrlUpsert) {
		s.UpdateURL()
	})
}

// SetExpiredAt sets the "expired_at" field.
func (u *ShortUrlUpsertOne) SetExpiredAt(v time.Time) *ShortUrlUpsertOne {
	return u.Update(func(s *ShortUrlUpsert) {
		s.SetExpiredAt(v)
	})
}

// UpdateExpiredAt sets the "expired_at" field to the value that was provided on create.
func (u *ShortUrlUpsertOne) UpdateExpiredAt() *ShortUrlUpsertOne {
	return u.Update(func(s *ShortUrlUpsert) {
		s.UpdateExpiredAt()
	})
}

// Exec executes the query.
func (u *ShortUrlUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("service: missing options for ShortUrlCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ShortUrlUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ShortUrlUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ShortUrlUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ShortUrlCreateBulk is the builder for creating many ShortUrl entities in bulk.
type ShortUrlCreateBulk struct {
	config
	builders []*ShortUrlCreate
	conflict []sql.ConflictOption
}

// Save creates the ShortUrl entities in the database.
func (sucb *ShortUrlCreateBulk) Save(ctx context.Context) ([]*ShortUrl, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sucb.builders))
	nodes := make([]*ShortUrl, len(sucb.builders))
	mutators := make([]Mutator, len(sucb.builders))
	for i := range sucb.builders {
		func(i int, root context.Context) {
			builder := sucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShortUrlMutation)
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
					_, err = mutators[i+1].Mutate(root, sucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sucb *ShortUrlCreateBulk) SaveX(ctx context.Context) []*ShortUrl {
	v, err := sucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sucb *ShortUrlCreateBulk) Exec(ctx context.Context) error {
	_, err := sucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sucb *ShortUrlCreateBulk) ExecX(ctx context.Context) {
	if err := sucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ShortUrl.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ShortUrlUpsert) {
//			SetURL(v+v).
//		}).
//		Exec(ctx)
//
func (sucb *ShortUrlCreateBulk) OnConflict(opts ...sql.ConflictOption) *ShortUrlUpsertBulk {
	sucb.conflict = opts
	return &ShortUrlUpsertBulk{
		create: sucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ShortUrl.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sucb *ShortUrlCreateBulk) OnConflictColumns(columns ...string) *ShortUrlUpsertBulk {
	sucb.conflict = append(sucb.conflict, sql.ConflictColumns(columns...))
	return &ShortUrlUpsertBulk{
		create: sucb,
	}
}

// ShortUrlUpsertBulk is the builder for "upsert"-ing
// a bulk of ShortUrl nodes.
type ShortUrlUpsertBulk struct {
	create *ShortUrlCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ShortUrl.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ShortUrlUpsertBulk) UpdateNewValues() *ShortUrlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ShortUrl.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ShortUrlUpsertBulk) Ignore() *ShortUrlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ShortUrlUpsertBulk) DoNothing() *ShortUrlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ShortUrlCreateBulk.OnConflict
// documentation for more info.
func (u *ShortUrlUpsertBulk) Update(set func(*ShortUrlUpsert)) *ShortUrlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ShortUrlUpsert{UpdateSet: update})
	}))
	return u
}

// SetURL sets the "url" field.
func (u *ShortUrlUpsertBulk) SetURL(v string) *ShortUrlUpsertBulk {
	return u.Update(func(s *ShortUrlUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *ShortUrlUpsertBulk) UpdateURL() *ShortUrlUpsertBulk {
	return u.Update(func(s *ShortUrlUpsert) {
		s.UpdateURL()
	})
}

// SetExpiredAt sets the "expired_at" field.
func (u *ShortUrlUpsertBulk) SetExpiredAt(v time.Time) *ShortUrlUpsertBulk {
	return u.Update(func(s *ShortUrlUpsert) {
		s.SetExpiredAt(v)
	})
}

// UpdateExpiredAt sets the "expired_at" field to the value that was provided on create.
func (u *ShortUrlUpsertBulk) UpdateExpiredAt() *ShortUrlUpsertBulk {
	return u.Update(func(s *ShortUrlUpsert) {
		s.UpdateExpiredAt()
	})
}

// Exec executes the query.
func (u *ShortUrlUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("service: OnConflict was set for builder %d. Set it on the ShortUrlCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("service: missing options for ShortUrlCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ShortUrlUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}