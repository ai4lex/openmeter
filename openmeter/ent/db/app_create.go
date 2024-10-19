// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	appentitybase "github.com/openmeterio/openmeter/openmeter/app/entity/base"
	"github.com/openmeterio/openmeter/openmeter/ent/db/app"
	"github.com/openmeterio/openmeter/openmeter/ent/db/appcustomer"
)

// AppCreate is the builder for creating a App entity.
type AppCreate struct {
	config
	mutation *AppMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNamespace sets the "namespace" field.
func (ac *AppCreate) SetNamespace(s string) *AppCreate {
	ac.mutation.SetNamespace(s)
	return ac
}

// SetMetadata sets the "metadata" field.
func (ac *AppCreate) SetMetadata(m map[string]string) *AppCreate {
	ac.mutation.SetMetadata(m)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AppCreate) SetCreatedAt(t time.Time) *AppCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableCreatedAt(t *time.Time) *AppCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AppCreate) SetUpdatedAt(t time.Time) *AppCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableUpdatedAt(t *time.Time) *AppCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AppCreate) SetDeletedAt(t time.Time) *AppCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableDeletedAt(t *time.Time) *AppCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *AppCreate) SetName(s string) *AppCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *AppCreate) SetDescription(s string) *AppCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *AppCreate) SetNillableDescription(s *string) *AppCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetType sets the "type" field.
func (ac *AppCreate) SetType(at appentitybase.AppType) *AppCreate {
	ac.mutation.SetType(at)
	return ac
}

// SetStatus sets the "status" field.
func (ac *AppCreate) SetStatus(as appentitybase.AppStatus) *AppCreate {
	ac.mutation.SetStatus(as)
	return ac
}

// SetIsDefault sets the "is_default" field.
func (ac *AppCreate) SetIsDefault(b bool) *AppCreate {
	ac.mutation.SetIsDefault(b)
	return ac
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (ac *AppCreate) SetNillableIsDefault(b *bool) *AppCreate {
	if b != nil {
		ac.SetIsDefault(*b)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AppCreate) SetID(s string) *AppCreate {
	ac.mutation.SetID(s)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AppCreate) SetNillableID(s *string) *AppCreate {
	if s != nil {
		ac.SetID(*s)
	}
	return ac
}

// AddCustomerAppIDs adds the "customer_apps" edge to the AppCustomer entity by IDs.
func (ac *AppCreate) AddCustomerAppIDs(ids ...int) *AppCreate {
	ac.mutation.AddCustomerAppIDs(ids...)
	return ac
}

// AddCustomerApps adds the "customer_apps" edges to the AppCustomer entity.
func (ac *AppCreate) AddCustomerApps(a ...*AppCustomer) *AppCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddCustomerAppIDs(ids...)
}

// Mutation returns the AppMutation object of the builder.
func (ac *AppCreate) Mutation() *AppMutation {
	return ac.mutation
}

// Save creates the App in the database.
func (ac *AppCreate) Save(ctx context.Context) (*App, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AppCreate) SaveX(ctx context.Context) *App {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AppCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AppCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AppCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := app.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := app.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.IsDefault(); !ok {
		v := app.DefaultIsDefault
		ac.mutation.SetIsDefault(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := app.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AppCreate) check() error {
	if _, ok := ac.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`db: missing required field "App.namespace"`)}
	}
	if v, ok := ac.mutation.Namespace(); ok {
		if err := app.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`db: validator failed for field "App.namespace": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "App.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "App.updated_at"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "App.name"`)}
	}
	if _, ok := ac.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`db: missing required field "App.type"`)}
	}
	if _, ok := ac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`db: missing required field "App.status"`)}
	}
	if _, ok := ac.mutation.IsDefault(); !ok {
		return &ValidationError{Name: "is_default", err: errors.New(`db: missing required field "App.is_default"`)}
	}
	return nil
}

func (ac *AppCreate) sqlSave(ctx context.Context) (*App, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected App.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AppCreate) createSpec() (*App, *sqlgraph.CreateSpec) {
	var (
		_node = &App{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(app.Table, sqlgraph.NewFieldSpec(app.FieldID, field.TypeString))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Namespace(); ok {
		_spec.SetField(app.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := ac.mutation.Metadata(); ok {
		_spec.SetField(app.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(app.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(app.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(app.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(app.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(app.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := ac.mutation.GetType(); ok {
		_spec.SetField(app.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(app.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := ac.mutation.IsDefault(); ok {
		_spec.SetField(app.FieldIsDefault, field.TypeBool, value)
		_node.IsDefault = value
	}
	if nodes := ac.mutation.CustomerAppsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   app.CustomerAppsTable,
			Columns: []string{app.CustomerAppsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appcustomer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.App.Create().
//		SetNamespace(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (ac *AppCreate) OnConflict(opts ...sql.ConflictOption) *AppUpsertOne {
	ac.conflict = opts
	return &AppUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AppCreate) OnConflictColumns(columns ...string) *AppUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AppUpsertOne{
		create: ac,
	}
}

type (
	// AppUpsertOne is the builder for "upsert"-ing
	//  one App node.
	AppUpsertOne struct {
		create *AppCreate
	}

	// AppUpsert is the "OnConflict" setter.
	AppUpsert struct {
		*sql.UpdateSet
	}
)

// SetMetadata sets the "metadata" field.
func (u *AppUpsert) SetMetadata(v map[string]string) *AppUpsert {
	u.Set(app.FieldMetadata, v)
	return u
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *AppUpsert) UpdateMetadata() *AppUpsert {
	u.SetExcluded(app.FieldMetadata)
	return u
}

// ClearMetadata clears the value of the "metadata" field.
func (u *AppUpsert) ClearMetadata() *AppUpsert {
	u.SetNull(app.FieldMetadata)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUpsert) SetUpdatedAt(v time.Time) *AppUpsert {
	u.Set(app.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUpsert) UpdateUpdatedAt() *AppUpsert {
	u.SetExcluded(app.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUpsert) SetDeletedAt(v time.Time) *AppUpsert {
	u.Set(app.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUpsert) UpdateDeletedAt() *AppUpsert {
	u.SetExcluded(app.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AppUpsert) ClearDeletedAt() *AppUpsert {
	u.SetNull(app.FieldDeletedAt)
	return u
}

// SetName sets the "name" field.
func (u *AppUpsert) SetName(v string) *AppUpsert {
	u.Set(app.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsert) UpdateName() *AppUpsert {
	u.SetExcluded(app.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *AppUpsert) SetDescription(v string) *AppUpsert {
	u.Set(app.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsert) UpdateDescription() *AppUpsert {
	u.SetExcluded(app.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *AppUpsert) ClearDescription() *AppUpsert {
	u.SetNull(app.FieldDescription)
	return u
}

// SetStatus sets the "status" field.
func (u *AppUpsert) SetStatus(v appentitybase.AppStatus) *AppUpsert {
	u.Set(app.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AppUpsert) UpdateStatus() *AppUpsert {
	u.SetExcluded(app.FieldStatus)
	return u
}

// SetIsDefault sets the "is_default" field.
func (u *AppUpsert) SetIsDefault(v bool) *AppUpsert {
	u.Set(app.FieldIsDefault, v)
	return u
}

// UpdateIsDefault sets the "is_default" field to the value that was provided on create.
func (u *AppUpsert) UpdateIsDefault() *AppUpsert {
	u.SetExcluded(app.FieldIsDefault)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(app.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppUpsertOne) UpdateNewValues() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(app.FieldID)
		}
		if _, exists := u.create.mutation.Namespace(); exists {
			s.SetIgnore(app.FieldNamespace)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(app.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.GetType(); exists {
			s.SetIgnore(app.FieldType)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.App.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppUpsertOne) Ignore() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUpsertOne) DoNothing() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCreate.OnConflict
// documentation for more info.
func (u *AppUpsertOne) Update(set func(*AppUpsert)) *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUpsert{UpdateSet: update})
	}))
	return u
}

// SetMetadata sets the "metadata" field.
func (u *AppUpsertOne) SetMetadata(v map[string]string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateMetadata() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateMetadata()
	})
}

// ClearMetadata clears the value of the "metadata" field.
func (u *AppUpsertOne) ClearMetadata() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.ClearMetadata()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUpsertOne) SetUpdatedAt(v time.Time) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateUpdatedAt() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUpsertOne) SetDeletedAt(v time.Time) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateDeletedAt() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AppUpsertOne) ClearDeletedAt() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *AppUpsertOne) SetName(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateName() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *AppUpsertOne) SetDescription(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateDescription() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *AppUpsertOne) ClearDescription() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.ClearDescription()
	})
}

// SetStatus sets the "status" field.
func (u *AppUpsertOne) SetStatus(v appentitybase.AppStatus) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateStatus() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateStatus()
	})
}

// SetIsDefault sets the "is_default" field.
func (u *AppUpsertOne) SetIsDefault(v bool) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetIsDefault(v)
	})
}

// UpdateIsDefault sets the "is_default" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateIsDefault() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateIsDefault()
	})
}

// Exec executes the query.
func (u *AppUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for AppCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: AppUpsertOne.ID is not supported by MySQL driver. Use AppUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppCreateBulk is the builder for creating many App entities in bulk.
type AppCreateBulk struct {
	config
	err      error
	builders []*AppCreate
	conflict []sql.ConflictOption
}

// Save creates the App entities in the database.
func (acb *AppCreateBulk) Save(ctx context.Context) ([]*App, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*App, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AppCreateBulk) SaveX(ctx context.Context) []*App {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AppCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AppCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.App.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (acb *AppCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppUpsertBulk {
	acb.conflict = opts
	return &AppUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AppCreateBulk) OnConflictColumns(columns ...string) *AppUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AppUpsertBulk{
		create: acb,
	}
}

// AppUpsertBulk is the builder for "upsert"-ing
// a bulk of App nodes.
type AppUpsertBulk struct {
	create *AppCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(app.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppUpsertBulk) UpdateNewValues() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(app.FieldID)
			}
			if _, exists := b.mutation.Namespace(); exists {
				s.SetIgnore(app.FieldNamespace)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(app.FieldCreatedAt)
			}
			if _, exists := b.mutation.GetType(); exists {
				s.SetIgnore(app.FieldType)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppUpsertBulk) Ignore() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUpsertBulk) DoNothing() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCreateBulk.OnConflict
// documentation for more info.
func (u *AppUpsertBulk) Update(set func(*AppUpsert)) *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUpsert{UpdateSet: update})
	}))
	return u
}

// SetMetadata sets the "metadata" field.
func (u *AppUpsertBulk) SetMetadata(v map[string]string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateMetadata() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateMetadata()
	})
}

// ClearMetadata clears the value of the "metadata" field.
func (u *AppUpsertBulk) ClearMetadata() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.ClearMetadata()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUpsertBulk) SetUpdatedAt(v time.Time) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateUpdatedAt() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUpsertBulk) SetDeletedAt(v time.Time) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateDeletedAt() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AppUpsertBulk) ClearDeletedAt() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *AppUpsertBulk) SetName(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateName() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *AppUpsertBulk) SetDescription(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateDescription() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *AppUpsertBulk) ClearDescription() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.ClearDescription()
	})
}

// SetStatus sets the "status" field.
func (u *AppUpsertBulk) SetStatus(v appentitybase.AppStatus) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateStatus() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateStatus()
	})
}

// SetIsDefault sets the "is_default" field.
func (u *AppUpsertBulk) SetIsDefault(v bool) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetIsDefault(v)
	})
}

// UpdateIsDefault sets the "is_default" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateIsDefault() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateIsDefault()
	})
}

// Exec executes the query.
func (u *AppUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the AppCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for AppCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}