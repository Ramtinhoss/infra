// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/env"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/internal"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/predicate"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/snapshot"
)

// SnapshotUpdate is the builder for updating Snapshot entities.
type SnapshotUpdate struct {
	config
	hooks     []Hook
	mutation  *SnapshotMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SnapshotUpdate builder.
func (su *SnapshotUpdate) Where(ps ...predicate.Snapshot) *SnapshotUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetPausedAt sets the "paused_at" field.
func (su *SnapshotUpdate) SetPausedAt(t time.Time) *SnapshotUpdate {
	su.mutation.SetPausedAt(t)
	return su
}

// SetNillablePausedAt sets the "paused_at" field if the given value is not nil.
func (su *SnapshotUpdate) SetNillablePausedAt(t *time.Time) *SnapshotUpdate {
	if t != nil {
		su.SetPausedAt(*t)
	}
	return su
}

// ClearPausedAt clears the value of the "paused_at" field.
func (su *SnapshotUpdate) ClearPausedAt() *SnapshotUpdate {
	su.mutation.ClearPausedAt()
	return su
}

// SetSandboxStartedAt sets the "sandbox_started_at" field.
func (su *SnapshotUpdate) SetSandboxStartedAt(t time.Time) *SnapshotUpdate {
	su.mutation.SetSandboxStartedAt(t)
	return su
}

// SetNillableSandboxStartedAt sets the "sandbox_started_at" field if the given value is not nil.
func (su *SnapshotUpdate) SetNillableSandboxStartedAt(t *time.Time) *SnapshotUpdate {
	if t != nil {
		su.SetSandboxStartedAt(*t)
	}
	return su
}

// SetBaseEnvID sets the "base_env_id" field.
func (su *SnapshotUpdate) SetBaseEnvID(s string) *SnapshotUpdate {
	su.mutation.SetBaseEnvID(s)
	return su
}

// SetNillableBaseEnvID sets the "base_env_id" field if the given value is not nil.
func (su *SnapshotUpdate) SetNillableBaseEnvID(s *string) *SnapshotUpdate {
	if s != nil {
		su.SetBaseEnvID(*s)
	}
	return su
}

// SetEnvID sets the "env_id" field.
func (su *SnapshotUpdate) SetEnvID(s string) *SnapshotUpdate {
	su.mutation.SetEnvID(s)
	return su
}

// SetNillableEnvID sets the "env_id" field if the given value is not nil.
func (su *SnapshotUpdate) SetNillableEnvID(s *string) *SnapshotUpdate {
	if s != nil {
		su.SetEnvID(*s)
	}
	return su
}

// SetSandboxID sets the "sandbox_id" field.
func (su *SnapshotUpdate) SetSandboxID(s string) *SnapshotUpdate {
	su.mutation.SetSandboxID(s)
	return su
}

// SetNillableSandboxID sets the "sandbox_id" field if the given value is not nil.
func (su *SnapshotUpdate) SetNillableSandboxID(s *string) *SnapshotUpdate {
	if s != nil {
		su.SetSandboxID(*s)
	}
	return su
}

// SetMetadata sets the "metadata" field.
func (su *SnapshotUpdate) SetMetadata(m map[string]string) *SnapshotUpdate {
	su.mutation.SetMetadata(m)
	return su
}

// SetEnv sets the "env" edge to the Env entity.
func (su *SnapshotUpdate) SetEnv(e *Env) *SnapshotUpdate {
	return su.SetEnvID(e.ID)
}

// Mutation returns the SnapshotMutation object of the builder.
func (su *SnapshotUpdate) Mutation() *SnapshotMutation {
	return su.mutation
}

// ClearEnv clears the "env" edge to the Env entity.
func (su *SnapshotUpdate) ClearEnv() *SnapshotUpdate {
	su.mutation.ClearEnv()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SnapshotUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SnapshotUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SnapshotUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SnapshotUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SnapshotUpdate) check() error {
	if _, ok := su.mutation.EnvID(); su.mutation.EnvCleared() && !ok {
		return errors.New(`models: clearing a required unique edge "Snapshot.env"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SnapshotUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SnapshotUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SnapshotUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(snapshot.Table, snapshot.Columns, sqlgraph.NewFieldSpec(snapshot.FieldID, field.TypeUUID))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.PausedAt(); ok {
		_spec.SetField(snapshot.FieldPausedAt, field.TypeTime, value)
	}
	if su.mutation.PausedAtCleared() {
		_spec.ClearField(snapshot.FieldPausedAt, field.TypeTime)
	}
	if value, ok := su.mutation.SandboxStartedAt(); ok {
		_spec.SetField(snapshot.FieldSandboxStartedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.BaseEnvID(); ok {
		_spec.SetField(snapshot.FieldBaseEnvID, field.TypeString, value)
	}
	if value, ok := su.mutation.SandboxID(); ok {
		_spec.SetField(snapshot.FieldSandboxID, field.TypeString, value)
	}
	if value, ok := su.mutation.Metadata(); ok {
		_spec.SetField(snapshot.FieldMetadata, field.TypeJSON, value)
	}
	if su.mutation.EnvCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   snapshot.EnvTable,
			Columns: []string{snapshot.EnvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = su.schemaConfig.Snapshot
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.EnvIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   snapshot.EnvTable,
			Columns: []string{snapshot.EnvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = su.schemaConfig.Snapshot
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = su.schemaConfig.Snapshot
	ctx = internal.NewSchemaConfigContext(ctx, su.schemaConfig)
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{snapshot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SnapshotUpdateOne is the builder for updating a single Snapshot entity.
type SnapshotUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SnapshotMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetPausedAt sets the "paused_at" field.
func (suo *SnapshotUpdateOne) SetPausedAt(t time.Time) *SnapshotUpdateOne {
	suo.mutation.SetPausedAt(t)
	return suo
}

// SetNillablePausedAt sets the "paused_at" field if the given value is not nil.
func (suo *SnapshotUpdateOne) SetNillablePausedAt(t *time.Time) *SnapshotUpdateOne {
	if t != nil {
		suo.SetPausedAt(*t)
	}
	return suo
}

// ClearPausedAt clears the value of the "paused_at" field.
func (suo *SnapshotUpdateOne) ClearPausedAt() *SnapshotUpdateOne {
	suo.mutation.ClearPausedAt()
	return suo
}

// SetSandboxStartedAt sets the "sandbox_started_at" field.
func (suo *SnapshotUpdateOne) SetSandboxStartedAt(t time.Time) *SnapshotUpdateOne {
	suo.mutation.SetSandboxStartedAt(t)
	return suo
}

// SetNillableSandboxStartedAt sets the "sandbox_started_at" field if the given value is not nil.
func (suo *SnapshotUpdateOne) SetNillableSandboxStartedAt(t *time.Time) *SnapshotUpdateOne {
	if t != nil {
		suo.SetSandboxStartedAt(*t)
	}
	return suo
}

// SetBaseEnvID sets the "base_env_id" field.
func (suo *SnapshotUpdateOne) SetBaseEnvID(s string) *SnapshotUpdateOne {
	suo.mutation.SetBaseEnvID(s)
	return suo
}

// SetNillableBaseEnvID sets the "base_env_id" field if the given value is not nil.
func (suo *SnapshotUpdateOne) SetNillableBaseEnvID(s *string) *SnapshotUpdateOne {
	if s != nil {
		suo.SetBaseEnvID(*s)
	}
	return suo
}

// SetEnvID sets the "env_id" field.
func (suo *SnapshotUpdateOne) SetEnvID(s string) *SnapshotUpdateOne {
	suo.mutation.SetEnvID(s)
	return suo
}

// SetNillableEnvID sets the "env_id" field if the given value is not nil.
func (suo *SnapshotUpdateOne) SetNillableEnvID(s *string) *SnapshotUpdateOne {
	if s != nil {
		suo.SetEnvID(*s)
	}
	return suo
}

// SetSandboxID sets the "sandbox_id" field.
func (suo *SnapshotUpdateOne) SetSandboxID(s string) *SnapshotUpdateOne {
	suo.mutation.SetSandboxID(s)
	return suo
}

// SetNillableSandboxID sets the "sandbox_id" field if the given value is not nil.
func (suo *SnapshotUpdateOne) SetNillableSandboxID(s *string) *SnapshotUpdateOne {
	if s != nil {
		suo.SetSandboxID(*s)
	}
	return suo
}

// SetMetadata sets the "metadata" field.
func (suo *SnapshotUpdateOne) SetMetadata(m map[string]string) *SnapshotUpdateOne {
	suo.mutation.SetMetadata(m)
	return suo
}

// SetEnv sets the "env" edge to the Env entity.
func (suo *SnapshotUpdateOne) SetEnv(e *Env) *SnapshotUpdateOne {
	return suo.SetEnvID(e.ID)
}

// Mutation returns the SnapshotMutation object of the builder.
func (suo *SnapshotUpdateOne) Mutation() *SnapshotMutation {
	return suo.mutation
}

// ClearEnv clears the "env" edge to the Env entity.
func (suo *SnapshotUpdateOne) ClearEnv() *SnapshotUpdateOne {
	suo.mutation.ClearEnv()
	return suo
}

// Where appends a list predicates to the SnapshotUpdate builder.
func (suo *SnapshotUpdateOne) Where(ps ...predicate.Snapshot) *SnapshotUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SnapshotUpdateOne) Select(field string, fields ...string) *SnapshotUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Snapshot entity.
func (suo *SnapshotUpdateOne) Save(ctx context.Context) (*Snapshot, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SnapshotUpdateOne) SaveX(ctx context.Context) *Snapshot {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SnapshotUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SnapshotUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SnapshotUpdateOne) check() error {
	if _, ok := suo.mutation.EnvID(); suo.mutation.EnvCleared() && !ok {
		return errors.New(`models: clearing a required unique edge "Snapshot.env"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SnapshotUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SnapshotUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SnapshotUpdateOne) sqlSave(ctx context.Context) (_node *Snapshot, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(snapshot.Table, snapshot.Columns, sqlgraph.NewFieldSpec(snapshot.FieldID, field.TypeUUID))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "Snapshot.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, snapshot.FieldID)
		for _, f := range fields {
			if !snapshot.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != snapshot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.PausedAt(); ok {
		_spec.SetField(snapshot.FieldPausedAt, field.TypeTime, value)
	}
	if suo.mutation.PausedAtCleared() {
		_spec.ClearField(snapshot.FieldPausedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.SandboxStartedAt(); ok {
		_spec.SetField(snapshot.FieldSandboxStartedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.BaseEnvID(); ok {
		_spec.SetField(snapshot.FieldBaseEnvID, field.TypeString, value)
	}
	if value, ok := suo.mutation.SandboxID(); ok {
		_spec.SetField(snapshot.FieldSandboxID, field.TypeString, value)
	}
	if value, ok := suo.mutation.Metadata(); ok {
		_spec.SetField(snapshot.FieldMetadata, field.TypeJSON, value)
	}
	if suo.mutation.EnvCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   snapshot.EnvTable,
			Columns: []string{snapshot.EnvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = suo.schemaConfig.Snapshot
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.EnvIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   snapshot.EnvTable,
			Columns: []string{snapshot.EnvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = suo.schemaConfig.Snapshot
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = suo.schemaConfig.Snapshot
	ctx = internal.NewSchemaConfigContext(ctx, suo.schemaConfig)
	_spec.AddModifiers(suo.modifiers...)
	_node = &Snapshot{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{snapshot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
