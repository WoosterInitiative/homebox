// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent/authtokens"
	"github.com/hay-kot/content/backend/ent/predicate"
	"github.com/hay-kot/content/backend/ent/user"
)

// AuthTokensQuery is the builder for querying AuthTokens entities.
type AuthTokensQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AuthTokens
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthTokensQuery builder.
func (atq *AuthTokensQuery) Where(ps ...predicate.AuthTokens) *AuthTokensQuery {
	atq.predicates = append(atq.predicates, ps...)
	return atq
}

// Limit adds a limit step to the query.
func (atq *AuthTokensQuery) Limit(limit int) *AuthTokensQuery {
	atq.limit = &limit
	return atq
}

// Offset adds an offset step to the query.
func (atq *AuthTokensQuery) Offset(offset int) *AuthTokensQuery {
	atq.offset = &offset
	return atq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (atq *AuthTokensQuery) Unique(unique bool) *AuthTokensQuery {
	atq.unique = &unique
	return atq
}

// Order adds an order step to the query.
func (atq *AuthTokensQuery) Order(o ...OrderFunc) *AuthTokensQuery {
	atq.order = append(atq.order, o...)
	return atq
}

// QueryUser chains the current query on the "user" edge.
func (atq *AuthTokensQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: atq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := atq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := atq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authtokens.Table, authtokens.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, authtokens.UserTable, authtokens.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(atq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AuthTokens entity from the query.
// Returns a *NotFoundError when no AuthTokens was found.
func (atq *AuthTokensQuery) First(ctx context.Context) (*AuthTokens, error) {
	nodes, err := atq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authtokens.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atq *AuthTokensQuery) FirstX(ctx context.Context) *AuthTokens {
	node, err := atq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthTokens ID from the query.
// Returns a *NotFoundError when no AuthTokens ID was found.
func (atq *AuthTokensQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authtokens.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atq *AuthTokensQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := atq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthTokens entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AuthTokens entity is found.
// Returns a *NotFoundError when no AuthTokens entities are found.
func (atq *AuthTokensQuery) Only(ctx context.Context) (*AuthTokens, error) {
	nodes, err := atq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authtokens.Label}
	default:
		return nil, &NotSingularError{authtokens.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atq *AuthTokensQuery) OnlyX(ctx context.Context) *AuthTokens {
	node, err := atq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthTokens ID in the query.
// Returns a *NotSingularError when more than one AuthTokens ID is found.
// Returns a *NotFoundError when no entities are found.
func (atq *AuthTokensQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = atq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authtokens.Label}
	default:
		err = &NotSingularError{authtokens.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atq *AuthTokensQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := atq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthTokensSlice.
func (atq *AuthTokensQuery) All(ctx context.Context) ([]*AuthTokens, error) {
	if err := atq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return atq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (atq *AuthTokensQuery) AllX(ctx context.Context) []*AuthTokens {
	nodes, err := atq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthTokens IDs.
func (atq *AuthTokensQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := atq.Select(authtokens.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atq *AuthTokensQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := atq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atq *AuthTokensQuery) Count(ctx context.Context) (int, error) {
	if err := atq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return atq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (atq *AuthTokensQuery) CountX(ctx context.Context) int {
	count, err := atq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atq *AuthTokensQuery) Exist(ctx context.Context) (bool, error) {
	if err := atq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return atq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (atq *AuthTokensQuery) ExistX(ctx context.Context) bool {
	exist, err := atq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthTokensQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atq *AuthTokensQuery) Clone() *AuthTokensQuery {
	if atq == nil {
		return nil
	}
	return &AuthTokensQuery{
		config:     atq.config,
		limit:      atq.limit,
		offset:     atq.offset,
		order:      append([]OrderFunc{}, atq.order...),
		predicates: append([]predicate.AuthTokens{}, atq.predicates...),
		withUser:   atq.withUser.Clone(),
		// clone intermediate query.
		sql:    atq.sql.Clone(),
		path:   atq.path,
		unique: atq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (atq *AuthTokensQuery) WithUser(opts ...func(*UserQuery)) *AuthTokensQuery {
	query := &UserQuery{config: atq.config}
	for _, opt := range opts {
		opt(query)
	}
	atq.withUser = query
	return atq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthTokens.Query().
//		GroupBy(authtokens.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (atq *AuthTokensQuery) GroupBy(field string, fields ...string) *AuthTokensGroupBy {
	grbuild := &AuthTokensGroupBy{config: atq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := atq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return atq.sqlQuery(ctx), nil
	}
	grbuild.label = authtokens.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.AuthTokens.Query().
//		Select(authtokens.FieldCreatedAt).
//		Scan(ctx, &v)
func (atq *AuthTokensQuery) Select(fields ...string) *AuthTokensSelect {
	atq.fields = append(atq.fields, fields...)
	selbuild := &AuthTokensSelect{AuthTokensQuery: atq}
	selbuild.label = authtokens.Label
	selbuild.flds, selbuild.scan = &atq.fields, selbuild.Scan
	return selbuild
}

func (atq *AuthTokensQuery) prepareQuery(ctx context.Context) error {
	for _, f := range atq.fields {
		if !authtokens.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if atq.path != nil {
		prev, err := atq.path(ctx)
		if err != nil {
			return err
		}
		atq.sql = prev
	}
	return nil
}

func (atq *AuthTokensQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AuthTokens, error) {
	var (
		nodes       = []*AuthTokens{}
		withFKs     = atq.withFKs
		_spec       = atq.querySpec()
		loadedTypes = [1]bool{
			atq.withUser != nil,
		}
	)
	if atq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, authtokens.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AuthTokens).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AuthTokens{config: atq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, atq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := atq.withUser; query != nil {
		if err := atq.loadUser(ctx, query, nodes, nil,
			func(n *AuthTokens, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (atq *AuthTokensQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*AuthTokens, init func(*AuthTokens), assign func(*AuthTokens, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AuthTokens)
	for i := range nodes {
		if nodes[i].user_auth_tokens == nil {
			continue
		}
		fk := *nodes[i].user_auth_tokens
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_auth_tokens" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (atq *AuthTokensQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atq.querySpec()
	_spec.Node.Columns = atq.fields
	if len(atq.fields) > 0 {
		_spec.Unique = atq.unique != nil && *atq.unique
	}
	return sqlgraph.CountNodes(ctx, atq.driver, _spec)
}

func (atq *AuthTokensQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := atq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (atq *AuthTokensQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authtokens.Table,
			Columns: authtokens.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: authtokens.FieldID,
			},
		},
		From:   atq.sql,
		Unique: true,
	}
	if unique := atq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := atq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authtokens.FieldID)
		for i := range fields {
			if fields[i] != authtokens.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := atq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (atq *AuthTokensQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(atq.driver.Dialect())
	t1 := builder.Table(authtokens.Table)
	columns := atq.fields
	if len(columns) == 0 {
		columns = authtokens.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if atq.sql != nil {
		selector = atq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if atq.unique != nil && *atq.unique {
		selector.Distinct()
	}
	for _, p := range atq.predicates {
		p(selector)
	}
	for _, p := range atq.order {
		p(selector)
	}
	if offset := atq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthTokensGroupBy is the group-by builder for AuthTokens entities.
type AuthTokensGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atgb *AuthTokensGroupBy) Aggregate(fns ...AggregateFunc) *AuthTokensGroupBy {
	atgb.fns = append(atgb.fns, fns...)
	return atgb
}

// Scan applies the group-by query and scans the result into the given value.
func (atgb *AuthTokensGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := atgb.path(ctx)
	if err != nil {
		return err
	}
	atgb.sql = query
	return atgb.sqlScan(ctx, v)
}

func (atgb *AuthTokensGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range atgb.fields {
		if !authtokens.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := atgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (atgb *AuthTokensGroupBy) sqlQuery() *sql.Selector {
	selector := atgb.sql.Select()
	aggregation := make([]string, 0, len(atgb.fns))
	for _, fn := range atgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(atgb.fields)+len(atgb.fns))
		for _, f := range atgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(atgb.fields...)...)
}

// AuthTokensSelect is the builder for selecting fields of AuthTokens entities.
type AuthTokensSelect struct {
	*AuthTokensQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ats *AuthTokensSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ats.prepareQuery(ctx); err != nil {
		return err
	}
	ats.sql = ats.AuthTokensQuery.sqlQuery(ctx)
	return ats.sqlScan(ctx, v)
}

func (ats *AuthTokensSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ats.sql.Query()
	if err := ats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
