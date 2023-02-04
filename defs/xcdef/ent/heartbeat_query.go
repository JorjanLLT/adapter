// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/auroraride/adapter/defs/xcdef/ent/battery"
	"github.com/auroraride/adapter/defs/xcdef/ent/heartbeat"
	"github.com/auroraride/adapter/defs/xcdef/ent/predicate"
)

// HeartbeatQuery is the builder for querying Heartbeat entities.
type HeartbeatQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.Heartbeat
	withBattery *BatteryQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HeartbeatQuery builder.
func (hq *HeartbeatQuery) Where(ps ...predicate.Heartbeat) *HeartbeatQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit the number of records to be returned by this query.
func (hq *HeartbeatQuery) Limit(limit int) *HeartbeatQuery {
	hq.ctx.Limit = &limit
	return hq
}

// Offset to start from.
func (hq *HeartbeatQuery) Offset(offset int) *HeartbeatQuery {
	hq.ctx.Offset = &offset
	return hq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hq *HeartbeatQuery) Unique(unique bool) *HeartbeatQuery {
	hq.ctx.Unique = &unique
	return hq
}

// Order specifies how the records should be ordered.
func (hq *HeartbeatQuery) Order(o ...OrderFunc) *HeartbeatQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// QueryBattery chains the current query on the "battery" edge.
func (hq *HeartbeatQuery) QueryBattery() *BatteryQuery {
	query := (&BatteryClient{config: hq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(heartbeat.Table, heartbeat.FieldID, selector),
			sqlgraph.To(battery.Table, battery.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, heartbeat.BatteryTable, heartbeat.BatteryColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Heartbeat entity from the query.
// Returns a *NotFoundError when no Heartbeat was found.
func (hq *HeartbeatQuery) First(ctx context.Context) (*Heartbeat, error) {
	nodes, err := hq.Limit(1).All(setContextOp(ctx, hq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{heartbeat.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HeartbeatQuery) FirstX(ctx context.Context) *Heartbeat {
	node, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Heartbeat ID from the query.
// Returns a *NotFoundError when no Heartbeat ID was found.
func (hq *HeartbeatQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(1).IDs(setContextOp(ctx, hq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{heartbeat.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hq *HeartbeatQuery) FirstIDX(ctx context.Context) int {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Heartbeat entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Heartbeat entity is found.
// Returns a *NotFoundError when no Heartbeat entities are found.
func (hq *HeartbeatQuery) Only(ctx context.Context) (*Heartbeat, error) {
	nodes, err := hq.Limit(2).All(setContextOp(ctx, hq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{heartbeat.Label}
	default:
		return nil, &NotSingularError{heartbeat.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HeartbeatQuery) OnlyX(ctx context.Context) *Heartbeat {
	node, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Heartbeat ID in the query.
// Returns a *NotSingularError when more than one Heartbeat ID is found.
// Returns a *NotFoundError when no entities are found.
func (hq *HeartbeatQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(2).IDs(setContextOp(ctx, hq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{heartbeat.Label}
	default:
		err = &NotSingularError{heartbeat.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hq *HeartbeatQuery) OnlyIDX(ctx context.Context) int {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Heartbeats.
func (hq *HeartbeatQuery) All(ctx context.Context) ([]*Heartbeat, error) {
	ctx = setContextOp(ctx, hq.ctx, "All")
	if err := hq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Heartbeat, *HeartbeatQuery]()
	return withInterceptors[[]*Heartbeat](ctx, hq, qr, hq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hq *HeartbeatQuery) AllX(ctx context.Context) []*Heartbeat {
	nodes, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Heartbeat IDs.
func (hq *HeartbeatQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, hq.ctx, "IDs")
	if err := hq.Select(heartbeat.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HeartbeatQuery) IDsX(ctx context.Context) []int {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HeartbeatQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hq.ctx, "Count")
	if err := hq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hq, querierCount[*HeartbeatQuery](), hq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HeartbeatQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HeartbeatQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hq.ctx, "Exist")
	switch _, err := hq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HeartbeatQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HeartbeatQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HeartbeatQuery) Clone() *HeartbeatQuery {
	if hq == nil {
		return nil
	}
	return &HeartbeatQuery{
		config:      hq.config,
		ctx:         hq.ctx.Clone(),
		order:       append([]OrderFunc{}, hq.order...),
		inters:      append([]Interceptor{}, hq.inters...),
		predicates:  append([]predicate.Heartbeat{}, hq.predicates...),
		withBattery: hq.withBattery.Clone(),
		// clone intermediate query.
		sql:  hq.sql.Clone(),
		path: hq.path,
	}
}

// WithBattery tells the query-builder to eager-load the nodes that are connected to
// the "battery" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HeartbeatQuery) WithBattery(opts ...func(*BatteryQuery)) *HeartbeatQuery {
	query := (&BatteryClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hq.withBattery = query
	return hq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Sn string `json:"sn,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Heartbeat.Query().
//		GroupBy(heartbeat.FieldSn).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (hq *HeartbeatQuery) GroupBy(field string, fields ...string) *HeartbeatGroupBy {
	hq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HeartbeatGroupBy{build: hq}
	grbuild.flds = &hq.ctx.Fields
	grbuild.label = heartbeat.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Sn string `json:"sn,omitempty"`
//	}
//
//	client.Heartbeat.Query().
//		Select(heartbeat.FieldSn).
//		Scan(ctx, &v)
func (hq *HeartbeatQuery) Select(fields ...string) *HeartbeatSelect {
	hq.ctx.Fields = append(hq.ctx.Fields, fields...)
	sbuild := &HeartbeatSelect{HeartbeatQuery: hq}
	sbuild.label = heartbeat.Label
	sbuild.flds, sbuild.scan = &hq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HeartbeatSelect configured with the given aggregations.
func (hq *HeartbeatQuery) Aggregate(fns ...AggregateFunc) *HeartbeatSelect {
	return hq.Select().Aggregate(fns...)
}

func (hq *HeartbeatQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hq); err != nil {
				return err
			}
		}
	}
	for _, f := range hq.ctx.Fields {
		if !heartbeat.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hq.path != nil {
		prev, err := hq.path(ctx)
		if err != nil {
			return err
		}
		hq.sql = prev
	}
	return nil
}

func (hq *HeartbeatQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Heartbeat, error) {
	var (
		nodes       = []*Heartbeat{}
		_spec       = hq.querySpec()
		loadedTypes = [1]bool{
			hq.withBattery != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Heartbeat).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Heartbeat{config: hq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(hq.modifiers) > 0 {
		_spec.Modifiers = hq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hq.withBattery; query != nil {
		if err := hq.loadBattery(ctx, query, nodes, nil,
			func(n *Heartbeat, e *Battery) { n.Edges.Battery = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hq *HeartbeatQuery) loadBattery(ctx context.Context, query *BatteryQuery, nodes []*Heartbeat, init func(*Heartbeat), assign func(*Heartbeat, *Battery)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Heartbeat)
	for i := range nodes {
		fk := nodes[i].BatteryID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(battery.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "battery_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (hq *HeartbeatQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	if len(hq.modifiers) > 0 {
		_spec.Modifiers = hq.modifiers
	}
	_spec.Node.Columns = hq.ctx.Fields
	if len(hq.ctx.Fields) > 0 {
		_spec.Unique = hq.ctx.Unique != nil && *hq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HeartbeatQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   heartbeat.Table,
			Columns: heartbeat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: heartbeat.FieldID,
			},
		},
		From:   hq.sql,
		Unique: true,
	}
	if unique := hq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := hq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, heartbeat.FieldID)
		for i := range fields {
			if fields[i] != heartbeat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hq *HeartbeatQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(heartbeat.Table)
	columns := hq.ctx.Fields
	if len(columns) == 0 {
		columns = heartbeat.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hq.ctx.Unique != nil && *hq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range hq.modifiers {
		m(selector)
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector)
	}
	if offset := hq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (hq *HeartbeatQuery) Modify(modifiers ...func(s *sql.Selector)) *HeartbeatSelect {
	hq.modifiers = append(hq.modifiers, modifiers...)
	return hq.Select()
}

// HeartbeatGroupBy is the group-by builder for Heartbeat entities.
type HeartbeatGroupBy struct {
	selector
	build *HeartbeatQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HeartbeatGroupBy) Aggregate(fns ...AggregateFunc) *HeartbeatGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the selector query and scans the result into the given value.
func (hgb *HeartbeatGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hgb.build.ctx, "GroupBy")
	if err := hgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HeartbeatQuery, *HeartbeatGroupBy](ctx, hgb.build, hgb, hgb.build.inters, v)
}

func (hgb *HeartbeatGroupBy) sqlScan(ctx context.Context, root *HeartbeatQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hgb.fns))
	for _, fn := range hgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hgb.flds)+len(hgb.fns))
		for _, f := range *hgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HeartbeatSelect is the builder for selecting fields of Heartbeat entities.
type HeartbeatSelect struct {
	*HeartbeatQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hs *HeartbeatSelect) Aggregate(fns ...AggregateFunc) *HeartbeatSelect {
	hs.fns = append(hs.fns, fns...)
	return hs
}

// Scan applies the selector query and scans the result into the given value.
func (hs *HeartbeatSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hs.ctx, "Select")
	if err := hs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HeartbeatQuery, *HeartbeatSelect](ctx, hs.HeartbeatQuery, hs, hs.inters, v)
}

func (hs *HeartbeatSelect) sqlScan(ctx context.Context, root *HeartbeatQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hs.fns))
	for _, fn := range hs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (hs *HeartbeatSelect) Modify(modifiers ...func(s *sql.Selector)) *HeartbeatSelect {
	hs.modifiers = append(hs.modifiers, modifiers...)
	return hs
}
