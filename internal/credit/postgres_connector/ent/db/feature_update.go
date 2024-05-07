// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/creditentry"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/feature"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/predicate"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/pgulid"
)

// FeatureUpdate is the builder for updating Feature entities.
type FeatureUpdate struct {
	config
	hooks    []Hook
	mutation *FeatureMutation
}

// Where appends a list predicates to the FeatureUpdate builder.
func (fu *FeatureUpdate) Where(ps ...predicate.Feature) *FeatureUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FeatureUpdate) SetUpdatedAt(t time.Time) *FeatureUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetName sets the "name" field.
func (fu *FeatureUpdate) SetName(s string) *FeatureUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fu *FeatureUpdate) SetNillableName(s *string) *FeatureUpdate {
	if s != nil {
		fu.SetName(*s)
	}
	return fu
}

// SetMeterGroupByFilters sets the "meter_group_by_filters" field.
func (fu *FeatureUpdate) SetMeterGroupByFilters(m map[string]string) *FeatureUpdate {
	fu.mutation.SetMeterGroupByFilters(m)
	return fu
}

// ClearMeterGroupByFilters clears the value of the "meter_group_by_filters" field.
func (fu *FeatureUpdate) ClearMeterGroupByFilters() *FeatureUpdate {
	fu.mutation.ClearMeterGroupByFilters()
	return fu
}

// SetArchived sets the "archived" field.
func (fu *FeatureUpdate) SetArchived(b bool) *FeatureUpdate {
	fu.mutation.SetArchived(b)
	return fu
}

// SetNillableArchived sets the "archived" field if the given value is not nil.
func (fu *FeatureUpdate) SetNillableArchived(b *bool) *FeatureUpdate {
	if b != nil {
		fu.SetArchived(*b)
	}
	return fu
}

// AddCreditGrantIDs adds the "credit_grants" edge to the CreditEntry entity by IDs.
func (fu *FeatureUpdate) AddCreditGrantIDs(ids ...pgulid.ULID) *FeatureUpdate {
	fu.mutation.AddCreditGrantIDs(ids...)
	return fu
}

// AddCreditGrants adds the "credit_grants" edges to the CreditEntry entity.
func (fu *FeatureUpdate) AddCreditGrants(c ...*CreditEntry) *FeatureUpdate {
	ids := make([]pgulid.ULID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return fu.AddCreditGrantIDs(ids...)
}

// Mutation returns the FeatureMutation object of the builder.
func (fu *FeatureUpdate) Mutation() *FeatureMutation {
	return fu.mutation
}

// ClearCreditGrants clears all "credit_grants" edges to the CreditEntry entity.
func (fu *FeatureUpdate) ClearCreditGrants() *FeatureUpdate {
	fu.mutation.ClearCreditGrants()
	return fu
}

// RemoveCreditGrantIDs removes the "credit_grants" edge to CreditEntry entities by IDs.
func (fu *FeatureUpdate) RemoveCreditGrantIDs(ids ...pgulid.ULID) *FeatureUpdate {
	fu.mutation.RemoveCreditGrantIDs(ids...)
	return fu
}

// RemoveCreditGrants removes "credit_grants" edges to CreditEntry entities.
func (fu *FeatureUpdate) RemoveCreditGrants(c ...*CreditEntry) *FeatureUpdate {
	ids := make([]pgulid.ULID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return fu.RemoveCreditGrantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FeatureUpdate) Save(ctx context.Context) (int, error) {
	fu.defaults()
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FeatureUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FeatureUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FeatureUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FeatureUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := feature.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FeatureUpdate) check() error {
	if v, ok := fu.mutation.Name(); ok {
		if err := feature.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "Feature.name": %w`, err)}
		}
	}
	return nil
}

func (fu *FeatureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(feature.Table, feature.Columns, sqlgraph.NewFieldSpec(feature.FieldID, field.TypeOther))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(feature.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(feature.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.MeterGroupByFilters(); ok {
		_spec.SetField(feature.FieldMeterGroupByFilters, field.TypeJSON, value)
	}
	if fu.mutation.MeterGroupByFiltersCleared() {
		_spec.ClearField(feature.FieldMeterGroupByFilters, field.TypeJSON)
	}
	if value, ok := fu.mutation.Archived(); ok {
		_spec.SetField(feature.FieldArchived, field.TypeBool, value)
	}
	if fu.mutation.CreditGrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feature.CreditGrantsTable,
			Columns: []string{feature.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedCreditGrantsIDs(); len(nodes) > 0 && !fu.mutation.CreditGrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feature.CreditGrantsTable,
			Columns: []string{feature.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.CreditGrantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feature.CreditGrantsTable,
			Columns: []string{feature.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feature.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FeatureUpdateOne is the builder for updating a single Feature entity.
type FeatureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FeatureMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FeatureUpdateOne) SetUpdatedAt(t time.Time) *FeatureUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetName sets the "name" field.
func (fuo *FeatureUpdateOne) SetName(s string) *FeatureUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fuo *FeatureUpdateOne) SetNillableName(s *string) *FeatureUpdateOne {
	if s != nil {
		fuo.SetName(*s)
	}
	return fuo
}

// SetMeterGroupByFilters sets the "meter_group_by_filters" field.
func (fuo *FeatureUpdateOne) SetMeterGroupByFilters(m map[string]string) *FeatureUpdateOne {
	fuo.mutation.SetMeterGroupByFilters(m)
	return fuo
}

// ClearMeterGroupByFilters clears the value of the "meter_group_by_filters" field.
func (fuo *FeatureUpdateOne) ClearMeterGroupByFilters() *FeatureUpdateOne {
	fuo.mutation.ClearMeterGroupByFilters()
	return fuo
}

// SetArchived sets the "archived" field.
func (fuo *FeatureUpdateOne) SetArchived(b bool) *FeatureUpdateOne {
	fuo.mutation.SetArchived(b)
	return fuo
}

// SetNillableArchived sets the "archived" field if the given value is not nil.
func (fuo *FeatureUpdateOne) SetNillableArchived(b *bool) *FeatureUpdateOne {
	if b != nil {
		fuo.SetArchived(*b)
	}
	return fuo
}

// AddCreditGrantIDs adds the "credit_grants" edge to the CreditEntry entity by IDs.
func (fuo *FeatureUpdateOne) AddCreditGrantIDs(ids ...pgulid.ULID) *FeatureUpdateOne {
	fuo.mutation.AddCreditGrantIDs(ids...)
	return fuo
}

// AddCreditGrants adds the "credit_grants" edges to the CreditEntry entity.
func (fuo *FeatureUpdateOne) AddCreditGrants(c ...*CreditEntry) *FeatureUpdateOne {
	ids := make([]pgulid.ULID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return fuo.AddCreditGrantIDs(ids...)
}

// Mutation returns the FeatureMutation object of the builder.
func (fuo *FeatureUpdateOne) Mutation() *FeatureMutation {
	return fuo.mutation
}

// ClearCreditGrants clears all "credit_grants" edges to the CreditEntry entity.
func (fuo *FeatureUpdateOne) ClearCreditGrants() *FeatureUpdateOne {
	fuo.mutation.ClearCreditGrants()
	return fuo
}

// RemoveCreditGrantIDs removes the "credit_grants" edge to CreditEntry entities by IDs.
func (fuo *FeatureUpdateOne) RemoveCreditGrantIDs(ids ...pgulid.ULID) *FeatureUpdateOne {
	fuo.mutation.RemoveCreditGrantIDs(ids...)
	return fuo
}

// RemoveCreditGrants removes "credit_grants" edges to CreditEntry entities.
func (fuo *FeatureUpdateOne) RemoveCreditGrants(c ...*CreditEntry) *FeatureUpdateOne {
	ids := make([]pgulid.ULID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return fuo.RemoveCreditGrantIDs(ids...)
}

// Where appends a list predicates to the FeatureUpdate builder.
func (fuo *FeatureUpdateOne) Where(ps ...predicate.Feature) *FeatureUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FeatureUpdateOne) Select(field string, fields ...string) *FeatureUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Feature entity.
func (fuo *FeatureUpdateOne) Save(ctx context.Context) (*Feature, error) {
	fuo.defaults()
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FeatureUpdateOne) SaveX(ctx context.Context) *Feature {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FeatureUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FeatureUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FeatureUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := feature.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FeatureUpdateOne) check() error {
	if v, ok := fuo.mutation.Name(); ok {
		if err := feature.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "Feature.name": %w`, err)}
		}
	}
	return nil
}

func (fuo *FeatureUpdateOne) sqlSave(ctx context.Context) (_node *Feature, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(feature.Table, feature.Columns, sqlgraph.NewFieldSpec(feature.FieldID, field.TypeOther))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Feature.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feature.FieldID)
		for _, f := range fields {
			if !feature.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != feature.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(feature.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(feature.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.MeterGroupByFilters(); ok {
		_spec.SetField(feature.FieldMeterGroupByFilters, field.TypeJSON, value)
	}
	if fuo.mutation.MeterGroupByFiltersCleared() {
		_spec.ClearField(feature.FieldMeterGroupByFilters, field.TypeJSON)
	}
	if value, ok := fuo.mutation.Archived(); ok {
		_spec.SetField(feature.FieldArchived, field.TypeBool, value)
	}
	if fuo.mutation.CreditGrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feature.CreditGrantsTable,
			Columns: []string{feature.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedCreditGrantsIDs(); len(nodes) > 0 && !fuo.mutation.CreditGrantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feature.CreditGrantsTable,
			Columns: []string{feature.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.CreditGrantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feature.CreditGrantsTable,
			Columns: []string{feature.CreditGrantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeOther),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Feature{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feature.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
