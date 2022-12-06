// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent/country"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent/lang"
)

// ent aliases to avoid import conflicts in user's code.
type (
	Op         = ent.Op
	Hook       = ent.Hook
	Value      = ent.Value
	Query      = ent.Query
	Policy     = ent.Policy
	Mutator    = ent.Mutator
	Mutation   = ent.Mutation
	MutateFunc = ent.MutateFunc
)

// OrderFunc applies an ordering on the sql selector.
type OrderFunc func(*sql.Selector)

// columnChecker returns a function indicates if the column exists in the given column.
func columnChecker(table string) func(string) error {
	checks := map[string]func(string) bool{
		country.Table: country.ValidColumn,
		lang.Table:    lang.ValidColumn,
	}
	check, ok := checks[table]
	if !ok {
		return func(string) error {
			return fmt.Errorf("unknown table %q", table)
		}
	}
	return func(column string) error {
		if !check(column) {
			return fmt.Errorf("unknown column %q for table %q", column, table)
		}
		return nil
	}
}

// Asc applies the given fields in ASC order.
func Asc(fields ...string) OrderFunc {
	return func(s *sql.Selector) {
		check := columnChecker(s.TableName())
		for _, f := range fields {
			if err := check(f); err != nil {
				s.AddError(&ValidationError{Name: f, err: fmt.Errorf("ent: %w", err)})
			}
			s.OrderBy(sql.Asc(s.C(f)))
		}
	}
}

// Desc applies the given fields in DESC order.
func Desc(fields ...string) OrderFunc {
	return func(s *sql.Selector) {
		check := columnChecker(s.TableName())
		for _, f := range fields {
			if err := check(f); err != nil {
				s.AddError(&ValidationError{Name: f, err: fmt.Errorf("ent: %w", err)})
			}
			s.OrderBy(sql.Desc(s.C(f)))
		}
	}
}

// AggregateFunc applies an aggregation step on the group-by traversal/selector.
type AggregateFunc func(*sql.Selector) string

// As is a pseudo aggregation function for renaming another other functions with custom names. For example:
//
//	GroupBy(field1, field2).
//	Aggregate(ent.As(ent.Sum(field1), "sum_field1"), (ent.As(ent.Sum(field2), "sum_field2")).
//	Scan(ctx, &v)
//
func As(fn AggregateFunc, end string) AggregateFunc {
	return func(s *sql.Selector) string {
		return sql.As(fn(s), end)
	}
}

// Count applies the "count" aggregation function on each group.
func Count() AggregateFunc {
	return func(s *sql.Selector) string {
		return sql.Count("*")
	}
}

// Max applies the "max" aggregation function on the given field of each group.
func Max(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Max(s.C(field))
	}
}

// Mean applies the "mean" aggregation function on the given field of each group.
func Mean(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Avg(s.C(field))
	}
}

// Min applies the "min" aggregation function on the given field of each group.
func Min(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Min(s.C(field))
	}
}

// Sum applies the "sum" aggregation function on the given field of each group.
func Sum(field string) AggregateFunc {
	return func(s *sql.Selector) string {
		check := columnChecker(s.TableName())
		if err := check(field); err != nil {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("ent: %w", err)})
			return ""
		}
		return sql.Sum(s.C(field))
	}
}

// ValidationError returns when validating a field or edge fails.
type ValidationError struct {
	Name string // Field or edge name.
	err  error
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return e.err.Error()
}

// Unwrap implements the errors.Wrapper interface.
func (e *ValidationError) Unwrap() error {
	return e.err
}

// IsValidationError returns a boolean indicating whether the error is a validation error.
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}
	var e *ValidationError
	return errors.As(err, &e)
}

// NotFoundError returns when trying to fetch a specific entity and it was not found in the database.
type NotFoundError struct {
	label string
}

// Error implements the error interface.
func (e *NotFoundError) Error() string {
	return "ent: " + e.label + " not found"
}

// IsNotFound returns a boolean indicating whether the error is a not found error.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *NotFoundError
	return errors.As(err, &e)
}

// MaskNotFound masks not found error.
func MaskNotFound(err error) error {
	if IsNotFound(err) {
		return nil
	}
	return err
}

// NotSingularError returns when trying to fetch a singular entity and more then one was found in the database.
type NotSingularError struct {
	label string
}

// Error implements the error interface.
func (e *NotSingularError) Error() string {
	return "ent: " + e.label + " not singular"
}

// IsNotSingular returns a boolean indicating whether the error is a not singular error.
func IsNotSingular(err error) bool {
	if err == nil {
		return false
	}
	var e *NotSingularError
	return errors.As(err, &e)
}

// NotLoadedError returns when trying to get a node that was not loaded by the query.
type NotLoadedError struct {
	edge string
}

// Error implements the error interface.
func (e *NotLoadedError) Error() string {
	return "ent: " + e.edge + " edge was not loaded"
}

// IsNotLoaded returns a boolean indicating whether the error is a not loaded error.
func IsNotLoaded(err error) bool {
	if err == nil {
		return false
	}
	var e *NotLoadedError
	return errors.As(err, &e)
}

// ConstraintError returns when trying to create/update one or more entities and
// one or more of their constraints failed. For example, violation of edge or
// field uniqueness.
type ConstraintError struct {
	msg  string
	wrap error
}

// Error implements the error interface.
func (e ConstraintError) Error() string {
	return "ent: constraint failed: " + e.msg
}

// Unwrap implements the errors.Wrapper interface.
func (e *ConstraintError) Unwrap() error {
	return e.wrap
}

// IsConstraintError returns a boolean indicating whether the error is a constraint failure.
func IsConstraintError(err error) bool {
	if err == nil {
		return false
	}
	var e *ConstraintError
	return errors.As(err, &e)
}

// selector embedded by the different Select/GroupBy builders.
type selector struct {
	label string
	flds  *[]string
	scan  func(context.Context, interface{}) error
}

// ScanX is like Scan, but panics if an error occurs.
func (s *selector) ScanX(ctx context.Context, v interface{}) {
	if err := s.scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (s *selector) Strings(ctx context.Context) ([]string, error) {
	if len(*s.flds) > 1 {
		return nil, errors.New("ent: Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := s.scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (s *selector) StringsX(ctx context.Context) []string {
	v, err := s.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (s *selector) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = s.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{s.label}
	default:
		err = fmt.Errorf("ent: Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (s *selector) StringX(ctx context.Context) string {
	v, err := s.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (s *selector) Ints(ctx context.Context) ([]int, error) {
	if len(*s.flds) > 1 {
		return nil, errors.New("ent: Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := s.scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (s *selector) IntsX(ctx context.Context) []int {
	v, err := s.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (s *selector) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = s.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{s.label}
	default:
		err = fmt.Errorf("ent: Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (s *selector) IntX(ctx context.Context) int {
	v, err := s.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (s *selector) Float64s(ctx context.Context) ([]float64, error) {
	if len(*s.flds) > 1 {
		return nil, errors.New("ent: Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := s.scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (s *selector) Float64sX(ctx context.Context) []float64 {
	v, err := s.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (s *selector) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = s.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{s.label}
	default:
		err = fmt.Errorf("ent: Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (s *selector) Float64X(ctx context.Context) float64 {
	v, err := s.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (s *selector) Bools(ctx context.Context) ([]bool, error) {
	if len(*s.flds) > 1 {
		return nil, errors.New("ent: Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := s.scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (s *selector) BoolsX(ctx context.Context) []bool {
	v, err := s.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (s *selector) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = s.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{s.label}
	default:
		err = fmt.Errorf("ent: Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (s *selector) BoolX(ctx context.Context) bool {
	v, err := s.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// queryHook describes an internal hook for the different sqlAll methods.
type queryHook func(context.Context, *sqlgraph.QuerySpec)
