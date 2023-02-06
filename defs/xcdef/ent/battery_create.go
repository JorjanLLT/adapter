// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auroraride/adapter/defs/xcdef/ent/battery"
	"github.com/auroraride/adapter/defs/xcdef/ent/fault"
	"github.com/auroraride/adapter/defs/xcdef/ent/heartbeat"
	"github.com/auroraride/adapter/defs/xcdef/ent/reign"
)

// BatteryCreate is the builder for creating a Battery entity.
type BatteryCreate struct {
	config
	mutation *BatteryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (bc *BatteryCreate) SetCreatedAt(t time.Time) *BatteryCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BatteryCreate) SetNillableCreatedAt(t *time.Time) *BatteryCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BatteryCreate) SetUpdatedAt(t time.Time) *BatteryCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BatteryCreate) SetNillableUpdatedAt(t *time.Time) *BatteryCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetSn sets the "sn" field.
func (bc *BatteryCreate) SetSn(s string) *BatteryCreate {
	bc.mutation.SetSn(s)
	return bc
}

// SetSoftVersion sets the "soft_version" field.
func (bc *BatteryCreate) SetSoftVersion(u uint16) *BatteryCreate {
	bc.mutation.SetSoftVersion(u)
	return bc
}

// SetNillableSoftVersion sets the "soft_version" field if the given value is not nil.
func (bc *BatteryCreate) SetNillableSoftVersion(u *uint16) *BatteryCreate {
	if u != nil {
		bc.SetSoftVersion(*u)
	}
	return bc
}

// SetHardVersion sets the "hard_version" field.
func (bc *BatteryCreate) SetHardVersion(u uint16) *BatteryCreate {
	bc.mutation.SetHardVersion(u)
	return bc
}

// SetNillableHardVersion sets the "hard_version" field if the given value is not nil.
func (bc *BatteryCreate) SetNillableHardVersion(u *uint16) *BatteryCreate {
	if u != nil {
		bc.SetHardVersion(*u)
	}
	return bc
}

// SetSoft4gVersion sets the "soft_4g_version" field.
func (bc *BatteryCreate) SetSoft4gVersion(u uint16) *BatteryCreate {
	bc.mutation.SetSoft4gVersion(u)
	return bc
}

// SetNillableSoft4gVersion sets the "soft_4g_version" field if the given value is not nil.
func (bc *BatteryCreate) SetNillableSoft4gVersion(u *uint16) *BatteryCreate {
	if u != nil {
		bc.SetSoft4gVersion(*u)
	}
	return bc
}

// SetHard4gVersion sets the "hard_4g_version" field.
func (bc *BatteryCreate) SetHard4gVersion(u uint16) *BatteryCreate {
	bc.mutation.SetHard4gVersion(u)
	return bc
}

// SetNillableHard4gVersion sets the "hard_4g_version" field if the given value is not nil.
func (bc *BatteryCreate) SetNillableHard4gVersion(u *uint16) *BatteryCreate {
	if u != nil {
		bc.SetHard4gVersion(*u)
	}
	return bc
}

// SetSn4g sets the "sn_4g" field.
func (bc *BatteryCreate) SetSn4g(u uint64) *BatteryCreate {
	bc.mutation.SetSn4g(u)
	return bc
}

// SetNillableSn4g sets the "sn_4g" field if the given value is not nil.
func (bc *BatteryCreate) SetNillableSn4g(u *uint64) *BatteryCreate {
	if u != nil {
		bc.SetSn4g(*u)
	}
	return bc
}

// SetIccid sets the "iccid" field.
func (bc *BatteryCreate) SetIccid(s string) *BatteryCreate {
	bc.mutation.SetIccid(s)
	return bc
}

// SetNillableIccid sets the "iccid" field if the given value is not nil.
func (bc *BatteryCreate) SetNillableIccid(s *string) *BatteryCreate {
	if s != nil {
		bc.SetIccid(*s)
	}
	return bc
}

// SetSoc sets the "soc" field.
func (bc *BatteryCreate) SetSoc(u uint16) *BatteryCreate {
	bc.mutation.SetSoc(u)
	return bc
}

// SetNillableSoc sets the "soc" field if the given value is not nil.
func (bc *BatteryCreate) SetNillableSoc(u *uint16) *BatteryCreate {
	if u != nil {
		bc.SetSoc(*u)
	}
	return bc
}

// AddHeartbeatIDs adds the "heartbeats" edge to the Heartbeat entity by IDs.
func (bc *BatteryCreate) AddHeartbeatIDs(ids ...int) *BatteryCreate {
	bc.mutation.AddHeartbeatIDs(ids...)
	return bc
}

// AddHeartbeats adds the "heartbeats" edges to the Heartbeat entity.
func (bc *BatteryCreate) AddHeartbeats(h ...*Heartbeat) *BatteryCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return bc.AddHeartbeatIDs(ids...)
}

// AddReignIDs adds the "reigns" edge to the Reign entity by IDs.
func (bc *BatteryCreate) AddReignIDs(ids ...int) *BatteryCreate {
	bc.mutation.AddReignIDs(ids...)
	return bc
}

// AddReigns adds the "reigns" edges to the Reign entity.
func (bc *BatteryCreate) AddReigns(r ...*Reign) *BatteryCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return bc.AddReignIDs(ids...)
}

// AddFaultLogIDs adds the "fault_log" edge to the Fault entity by IDs.
func (bc *BatteryCreate) AddFaultLogIDs(ids ...int) *BatteryCreate {
	bc.mutation.AddFaultLogIDs(ids...)
	return bc
}

// AddFaultLog adds the "fault_log" edges to the Fault entity.
func (bc *BatteryCreate) AddFaultLog(f ...*Fault) *BatteryCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return bc.AddFaultLogIDs(ids...)
}

// Mutation returns the BatteryMutation object of the builder.
func (bc *BatteryCreate) Mutation() *BatteryMutation {
	return bc.mutation
}

// Save creates the Battery in the database.
func (bc *BatteryCreate) Save(ctx context.Context) (*Battery, error) {
	bc.defaults()
	return withHooks[*Battery, BatteryMutation](ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BatteryCreate) SaveX(ctx context.Context) *Battery {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BatteryCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BatteryCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BatteryCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := battery.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := battery.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BatteryCreate) check() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Battery.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Battery.updated_at"`)}
	}
	if _, ok := bc.mutation.Sn(); !ok {
		return &ValidationError{Name: "sn", err: errors.New(`ent: missing required field "Battery.sn"`)}
	}
	return nil
}

func (bc *BatteryCreate) sqlSave(ctx context.Context) (*Battery, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BatteryCreate) createSpec() (*Battery, *sqlgraph.CreateSpec) {
	var (
		_node = &Battery{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: battery.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: battery.FieldID,
			},
		}
	)
	_spec.OnConflict = bc.conflict
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(battery.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(battery.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.Sn(); ok {
		_spec.SetField(battery.FieldSn, field.TypeString, value)
		_node.Sn = value
	}
	if value, ok := bc.mutation.SoftVersion(); ok {
		_spec.SetField(battery.FieldSoftVersion, field.TypeUint16, value)
		_node.SoftVersion = &value
	}
	if value, ok := bc.mutation.HardVersion(); ok {
		_spec.SetField(battery.FieldHardVersion, field.TypeUint16, value)
		_node.HardVersion = &value
	}
	if value, ok := bc.mutation.Soft4gVersion(); ok {
		_spec.SetField(battery.FieldSoft4gVersion, field.TypeUint16, value)
		_node.Soft4gVersion = &value
	}
	if value, ok := bc.mutation.Hard4gVersion(); ok {
		_spec.SetField(battery.FieldHard4gVersion, field.TypeUint16, value)
		_node.Hard4gVersion = &value
	}
	if value, ok := bc.mutation.Sn4g(); ok {
		_spec.SetField(battery.FieldSn4g, field.TypeUint64, value)
		_node.Sn4g = &value
	}
	if value, ok := bc.mutation.Iccid(); ok {
		_spec.SetField(battery.FieldIccid, field.TypeString, value)
		_node.Iccid = &value
	}
	if value, ok := bc.mutation.Soc(); ok {
		_spec.SetField(battery.FieldSoc, field.TypeUint16, value)
		_node.Soc = &value
	}
	if nodes := bc.mutation.HeartbeatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   battery.HeartbeatsTable,
			Columns: []string{battery.HeartbeatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: heartbeat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.ReignsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   battery.ReignsTable,
			Columns: []string{battery.ReignsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reign.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.FaultLogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   battery.FaultLogTable,
			Columns: []string{battery.FaultLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fault.FieldID,
				},
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
//	client.Battery.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BatteryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (bc *BatteryCreate) OnConflict(opts ...sql.ConflictOption) *BatteryUpsertOne {
	bc.conflict = opts
	return &BatteryUpsertOne{
		create: bc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Battery.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bc *BatteryCreate) OnConflictColumns(columns ...string) *BatteryUpsertOne {
	bc.conflict = append(bc.conflict, sql.ConflictColumns(columns...))
	return &BatteryUpsertOne{
		create: bc,
	}
}

type (
	// BatteryUpsertOne is the builder for "upsert"-ing
	//  one Battery node.
	BatteryUpsertOne struct {
		create *BatteryCreate
	}

	// BatteryUpsert is the "OnConflict" setter.
	BatteryUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *BatteryUpsert) SetUpdatedAt(v time.Time) *BatteryUpsert {
	u.Set(battery.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BatteryUpsert) UpdateUpdatedAt() *BatteryUpsert {
	u.SetExcluded(battery.FieldUpdatedAt)
	return u
}

// SetSn sets the "sn" field.
func (u *BatteryUpsert) SetSn(v string) *BatteryUpsert {
	u.Set(battery.FieldSn, v)
	return u
}

// UpdateSn sets the "sn" field to the value that was provided on create.
func (u *BatteryUpsert) UpdateSn() *BatteryUpsert {
	u.SetExcluded(battery.FieldSn)
	return u
}

// SetSoftVersion sets the "soft_version" field.
func (u *BatteryUpsert) SetSoftVersion(v uint16) *BatteryUpsert {
	u.Set(battery.FieldSoftVersion, v)
	return u
}

// UpdateSoftVersion sets the "soft_version" field to the value that was provided on create.
func (u *BatteryUpsert) UpdateSoftVersion() *BatteryUpsert {
	u.SetExcluded(battery.FieldSoftVersion)
	return u
}

// AddSoftVersion adds v to the "soft_version" field.
func (u *BatteryUpsert) AddSoftVersion(v uint16) *BatteryUpsert {
	u.Add(battery.FieldSoftVersion, v)
	return u
}

// ClearSoftVersion clears the value of the "soft_version" field.
func (u *BatteryUpsert) ClearSoftVersion() *BatteryUpsert {
	u.SetNull(battery.FieldSoftVersion)
	return u
}

// SetHardVersion sets the "hard_version" field.
func (u *BatteryUpsert) SetHardVersion(v uint16) *BatteryUpsert {
	u.Set(battery.FieldHardVersion, v)
	return u
}

// UpdateHardVersion sets the "hard_version" field to the value that was provided on create.
func (u *BatteryUpsert) UpdateHardVersion() *BatteryUpsert {
	u.SetExcluded(battery.FieldHardVersion)
	return u
}

// AddHardVersion adds v to the "hard_version" field.
func (u *BatteryUpsert) AddHardVersion(v uint16) *BatteryUpsert {
	u.Add(battery.FieldHardVersion, v)
	return u
}

// ClearHardVersion clears the value of the "hard_version" field.
func (u *BatteryUpsert) ClearHardVersion() *BatteryUpsert {
	u.SetNull(battery.FieldHardVersion)
	return u
}

// SetSoft4gVersion sets the "soft_4g_version" field.
func (u *BatteryUpsert) SetSoft4gVersion(v uint16) *BatteryUpsert {
	u.Set(battery.FieldSoft4gVersion, v)
	return u
}

// UpdateSoft4gVersion sets the "soft_4g_version" field to the value that was provided on create.
func (u *BatteryUpsert) UpdateSoft4gVersion() *BatteryUpsert {
	u.SetExcluded(battery.FieldSoft4gVersion)
	return u
}

// AddSoft4gVersion adds v to the "soft_4g_version" field.
func (u *BatteryUpsert) AddSoft4gVersion(v uint16) *BatteryUpsert {
	u.Add(battery.FieldSoft4gVersion, v)
	return u
}

// ClearSoft4gVersion clears the value of the "soft_4g_version" field.
func (u *BatteryUpsert) ClearSoft4gVersion() *BatteryUpsert {
	u.SetNull(battery.FieldSoft4gVersion)
	return u
}

// SetHard4gVersion sets the "hard_4g_version" field.
func (u *BatteryUpsert) SetHard4gVersion(v uint16) *BatteryUpsert {
	u.Set(battery.FieldHard4gVersion, v)
	return u
}

// UpdateHard4gVersion sets the "hard_4g_version" field to the value that was provided on create.
func (u *BatteryUpsert) UpdateHard4gVersion() *BatteryUpsert {
	u.SetExcluded(battery.FieldHard4gVersion)
	return u
}

// AddHard4gVersion adds v to the "hard_4g_version" field.
func (u *BatteryUpsert) AddHard4gVersion(v uint16) *BatteryUpsert {
	u.Add(battery.FieldHard4gVersion, v)
	return u
}

// ClearHard4gVersion clears the value of the "hard_4g_version" field.
func (u *BatteryUpsert) ClearHard4gVersion() *BatteryUpsert {
	u.SetNull(battery.FieldHard4gVersion)
	return u
}

// SetSn4g sets the "sn_4g" field.
func (u *BatteryUpsert) SetSn4g(v uint64) *BatteryUpsert {
	u.Set(battery.FieldSn4g, v)
	return u
}

// UpdateSn4g sets the "sn_4g" field to the value that was provided on create.
func (u *BatteryUpsert) UpdateSn4g() *BatteryUpsert {
	u.SetExcluded(battery.FieldSn4g)
	return u
}

// AddSn4g adds v to the "sn_4g" field.
func (u *BatteryUpsert) AddSn4g(v uint64) *BatteryUpsert {
	u.Add(battery.FieldSn4g, v)
	return u
}

// ClearSn4g clears the value of the "sn_4g" field.
func (u *BatteryUpsert) ClearSn4g() *BatteryUpsert {
	u.SetNull(battery.FieldSn4g)
	return u
}

// SetIccid sets the "iccid" field.
func (u *BatteryUpsert) SetIccid(v string) *BatteryUpsert {
	u.Set(battery.FieldIccid, v)
	return u
}

// UpdateIccid sets the "iccid" field to the value that was provided on create.
func (u *BatteryUpsert) UpdateIccid() *BatteryUpsert {
	u.SetExcluded(battery.FieldIccid)
	return u
}

// ClearIccid clears the value of the "iccid" field.
func (u *BatteryUpsert) ClearIccid() *BatteryUpsert {
	u.SetNull(battery.FieldIccid)
	return u
}

// SetSoc sets the "soc" field.
func (u *BatteryUpsert) SetSoc(v uint16) *BatteryUpsert {
	u.Set(battery.FieldSoc, v)
	return u
}

// UpdateSoc sets the "soc" field to the value that was provided on create.
func (u *BatteryUpsert) UpdateSoc() *BatteryUpsert {
	u.SetExcluded(battery.FieldSoc)
	return u
}

// AddSoc adds v to the "soc" field.
func (u *BatteryUpsert) AddSoc(v uint16) *BatteryUpsert {
	u.Add(battery.FieldSoc, v)
	return u
}

// ClearSoc clears the value of the "soc" field.
func (u *BatteryUpsert) ClearSoc() *BatteryUpsert {
	u.SetNull(battery.FieldSoc)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Battery.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *BatteryUpsertOne) UpdateNewValues() *BatteryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(battery.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Battery.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BatteryUpsertOne) Ignore() *BatteryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BatteryUpsertOne) DoNothing() *BatteryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BatteryCreate.OnConflict
// documentation for more info.
func (u *BatteryUpsertOne) Update(set func(*BatteryUpsert)) *BatteryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BatteryUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *BatteryUpsertOne) SetUpdatedAt(v time.Time) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BatteryUpsertOne) UpdateUpdatedAt() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetSn sets the "sn" field.
func (u *BatteryUpsertOne) SetSn(v string) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.SetSn(v)
	})
}

// UpdateSn sets the "sn" field to the value that was provided on create.
func (u *BatteryUpsertOne) UpdateSn() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateSn()
	})
}

// SetSoftVersion sets the "soft_version" field.
func (u *BatteryUpsertOne) SetSoftVersion(v uint16) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.SetSoftVersion(v)
	})
}

// AddSoftVersion adds v to the "soft_version" field.
func (u *BatteryUpsertOne) AddSoftVersion(v uint16) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.AddSoftVersion(v)
	})
}

// UpdateSoftVersion sets the "soft_version" field to the value that was provided on create.
func (u *BatteryUpsertOne) UpdateSoftVersion() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateSoftVersion()
	})
}

// ClearSoftVersion clears the value of the "soft_version" field.
func (u *BatteryUpsertOne) ClearSoftVersion() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearSoftVersion()
	})
}

// SetHardVersion sets the "hard_version" field.
func (u *BatteryUpsertOne) SetHardVersion(v uint16) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.SetHardVersion(v)
	})
}

// AddHardVersion adds v to the "hard_version" field.
func (u *BatteryUpsertOne) AddHardVersion(v uint16) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.AddHardVersion(v)
	})
}

// UpdateHardVersion sets the "hard_version" field to the value that was provided on create.
func (u *BatteryUpsertOne) UpdateHardVersion() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateHardVersion()
	})
}

// ClearHardVersion clears the value of the "hard_version" field.
func (u *BatteryUpsertOne) ClearHardVersion() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearHardVersion()
	})
}

// SetSoft4gVersion sets the "soft_4g_version" field.
func (u *BatteryUpsertOne) SetSoft4gVersion(v uint16) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.SetSoft4gVersion(v)
	})
}

// AddSoft4gVersion adds v to the "soft_4g_version" field.
func (u *BatteryUpsertOne) AddSoft4gVersion(v uint16) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.AddSoft4gVersion(v)
	})
}

// UpdateSoft4gVersion sets the "soft_4g_version" field to the value that was provided on create.
func (u *BatteryUpsertOne) UpdateSoft4gVersion() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateSoft4gVersion()
	})
}

// ClearSoft4gVersion clears the value of the "soft_4g_version" field.
func (u *BatteryUpsertOne) ClearSoft4gVersion() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearSoft4gVersion()
	})
}

// SetHard4gVersion sets the "hard_4g_version" field.
func (u *BatteryUpsertOne) SetHard4gVersion(v uint16) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.SetHard4gVersion(v)
	})
}

// AddHard4gVersion adds v to the "hard_4g_version" field.
func (u *BatteryUpsertOne) AddHard4gVersion(v uint16) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.AddHard4gVersion(v)
	})
}

// UpdateHard4gVersion sets the "hard_4g_version" field to the value that was provided on create.
func (u *BatteryUpsertOne) UpdateHard4gVersion() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateHard4gVersion()
	})
}

// ClearHard4gVersion clears the value of the "hard_4g_version" field.
func (u *BatteryUpsertOne) ClearHard4gVersion() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearHard4gVersion()
	})
}

// SetSn4g sets the "sn_4g" field.
func (u *BatteryUpsertOne) SetSn4g(v uint64) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.SetSn4g(v)
	})
}

// AddSn4g adds v to the "sn_4g" field.
func (u *BatteryUpsertOne) AddSn4g(v uint64) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.AddSn4g(v)
	})
}

// UpdateSn4g sets the "sn_4g" field to the value that was provided on create.
func (u *BatteryUpsertOne) UpdateSn4g() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateSn4g()
	})
}

// ClearSn4g clears the value of the "sn_4g" field.
func (u *BatteryUpsertOne) ClearSn4g() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearSn4g()
	})
}

// SetIccid sets the "iccid" field.
func (u *BatteryUpsertOne) SetIccid(v string) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.SetIccid(v)
	})
}

// UpdateIccid sets the "iccid" field to the value that was provided on create.
func (u *BatteryUpsertOne) UpdateIccid() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateIccid()
	})
}

// ClearIccid clears the value of the "iccid" field.
func (u *BatteryUpsertOne) ClearIccid() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearIccid()
	})
}

// SetSoc sets the "soc" field.
func (u *BatteryUpsertOne) SetSoc(v uint16) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.SetSoc(v)
	})
}

// AddSoc adds v to the "soc" field.
func (u *BatteryUpsertOne) AddSoc(v uint16) *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.AddSoc(v)
	})
}

// UpdateSoc sets the "soc" field to the value that was provided on create.
func (u *BatteryUpsertOne) UpdateSoc() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateSoc()
	})
}

// ClearSoc clears the value of the "soc" field.
func (u *BatteryUpsertOne) ClearSoc() *BatteryUpsertOne {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearSoc()
	})
}

// Exec executes the query.
func (u *BatteryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BatteryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BatteryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BatteryUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BatteryUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BatteryCreateBulk is the builder for creating many Battery entities in bulk.
type BatteryCreateBulk struct {
	config
	builders []*BatteryCreate
	conflict []sql.ConflictOption
}

// Save creates the Battery entities in the database.
func (bcb *BatteryCreateBulk) Save(ctx context.Context) ([]*Battery, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Battery, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BatteryMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BatteryCreateBulk) SaveX(ctx context.Context) []*Battery {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BatteryCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BatteryCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Battery.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BatteryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (bcb *BatteryCreateBulk) OnConflict(opts ...sql.ConflictOption) *BatteryUpsertBulk {
	bcb.conflict = opts
	return &BatteryUpsertBulk{
		create: bcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Battery.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bcb *BatteryCreateBulk) OnConflictColumns(columns ...string) *BatteryUpsertBulk {
	bcb.conflict = append(bcb.conflict, sql.ConflictColumns(columns...))
	return &BatteryUpsertBulk{
		create: bcb,
	}
}

// BatteryUpsertBulk is the builder for "upsert"-ing
// a bulk of Battery nodes.
type BatteryUpsertBulk struct {
	create *BatteryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Battery.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *BatteryUpsertBulk) UpdateNewValues() *BatteryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(battery.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Battery.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BatteryUpsertBulk) Ignore() *BatteryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BatteryUpsertBulk) DoNothing() *BatteryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BatteryCreateBulk.OnConflict
// documentation for more info.
func (u *BatteryUpsertBulk) Update(set func(*BatteryUpsert)) *BatteryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BatteryUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *BatteryUpsertBulk) SetUpdatedAt(v time.Time) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BatteryUpsertBulk) UpdateUpdatedAt() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetSn sets the "sn" field.
func (u *BatteryUpsertBulk) SetSn(v string) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.SetSn(v)
	})
}

// UpdateSn sets the "sn" field to the value that was provided on create.
func (u *BatteryUpsertBulk) UpdateSn() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateSn()
	})
}

// SetSoftVersion sets the "soft_version" field.
func (u *BatteryUpsertBulk) SetSoftVersion(v uint16) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.SetSoftVersion(v)
	})
}

// AddSoftVersion adds v to the "soft_version" field.
func (u *BatteryUpsertBulk) AddSoftVersion(v uint16) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.AddSoftVersion(v)
	})
}

// UpdateSoftVersion sets the "soft_version" field to the value that was provided on create.
func (u *BatteryUpsertBulk) UpdateSoftVersion() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateSoftVersion()
	})
}

// ClearSoftVersion clears the value of the "soft_version" field.
func (u *BatteryUpsertBulk) ClearSoftVersion() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearSoftVersion()
	})
}

// SetHardVersion sets the "hard_version" field.
func (u *BatteryUpsertBulk) SetHardVersion(v uint16) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.SetHardVersion(v)
	})
}

// AddHardVersion adds v to the "hard_version" field.
func (u *BatteryUpsertBulk) AddHardVersion(v uint16) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.AddHardVersion(v)
	})
}

// UpdateHardVersion sets the "hard_version" field to the value that was provided on create.
func (u *BatteryUpsertBulk) UpdateHardVersion() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateHardVersion()
	})
}

// ClearHardVersion clears the value of the "hard_version" field.
func (u *BatteryUpsertBulk) ClearHardVersion() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearHardVersion()
	})
}

// SetSoft4gVersion sets the "soft_4g_version" field.
func (u *BatteryUpsertBulk) SetSoft4gVersion(v uint16) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.SetSoft4gVersion(v)
	})
}

// AddSoft4gVersion adds v to the "soft_4g_version" field.
func (u *BatteryUpsertBulk) AddSoft4gVersion(v uint16) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.AddSoft4gVersion(v)
	})
}

// UpdateSoft4gVersion sets the "soft_4g_version" field to the value that was provided on create.
func (u *BatteryUpsertBulk) UpdateSoft4gVersion() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateSoft4gVersion()
	})
}

// ClearSoft4gVersion clears the value of the "soft_4g_version" field.
func (u *BatteryUpsertBulk) ClearSoft4gVersion() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearSoft4gVersion()
	})
}

// SetHard4gVersion sets the "hard_4g_version" field.
func (u *BatteryUpsertBulk) SetHard4gVersion(v uint16) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.SetHard4gVersion(v)
	})
}

// AddHard4gVersion adds v to the "hard_4g_version" field.
func (u *BatteryUpsertBulk) AddHard4gVersion(v uint16) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.AddHard4gVersion(v)
	})
}

// UpdateHard4gVersion sets the "hard_4g_version" field to the value that was provided on create.
func (u *BatteryUpsertBulk) UpdateHard4gVersion() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateHard4gVersion()
	})
}

// ClearHard4gVersion clears the value of the "hard_4g_version" field.
func (u *BatteryUpsertBulk) ClearHard4gVersion() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearHard4gVersion()
	})
}

// SetSn4g sets the "sn_4g" field.
func (u *BatteryUpsertBulk) SetSn4g(v uint64) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.SetSn4g(v)
	})
}

// AddSn4g adds v to the "sn_4g" field.
func (u *BatteryUpsertBulk) AddSn4g(v uint64) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.AddSn4g(v)
	})
}

// UpdateSn4g sets the "sn_4g" field to the value that was provided on create.
func (u *BatteryUpsertBulk) UpdateSn4g() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateSn4g()
	})
}

// ClearSn4g clears the value of the "sn_4g" field.
func (u *BatteryUpsertBulk) ClearSn4g() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearSn4g()
	})
}

// SetIccid sets the "iccid" field.
func (u *BatteryUpsertBulk) SetIccid(v string) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.SetIccid(v)
	})
}

// UpdateIccid sets the "iccid" field to the value that was provided on create.
func (u *BatteryUpsertBulk) UpdateIccid() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateIccid()
	})
}

// ClearIccid clears the value of the "iccid" field.
func (u *BatteryUpsertBulk) ClearIccid() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearIccid()
	})
}

// SetSoc sets the "soc" field.
func (u *BatteryUpsertBulk) SetSoc(v uint16) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.SetSoc(v)
	})
}

// AddSoc adds v to the "soc" field.
func (u *BatteryUpsertBulk) AddSoc(v uint16) *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.AddSoc(v)
	})
}

// UpdateSoc sets the "soc" field to the value that was provided on create.
func (u *BatteryUpsertBulk) UpdateSoc() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.UpdateSoc()
	})
}

// ClearSoc clears the value of the "soc" field.
func (u *BatteryUpsertBulk) ClearSoc() *BatteryUpsertBulk {
	return u.Update(func(s *BatteryUpsert) {
		s.ClearSoc()
	})
}

// Exec executes the query.
func (u *BatteryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BatteryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BatteryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BatteryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
