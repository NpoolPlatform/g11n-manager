// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent/country"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent/lang"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   country.Table,
			Columns: country.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: country.FieldID,
			},
		},
		Type: "Country",
		Fields: map[string]*sqlgraph.FieldSpec{
			country.FieldCreatedAt: {Type: field.TypeUint32, Column: country.FieldCreatedAt},
			country.FieldUpdatedAt: {Type: field.TypeUint32, Column: country.FieldUpdatedAt},
			country.FieldDeletedAt: {Type: field.TypeUint32, Column: country.FieldDeletedAt},
			country.FieldCountry:   {Type: field.TypeString, Column: country.FieldCountry},
			country.FieldFlag:      {Type: field.TypeString, Column: country.FieldFlag},
			country.FieldCode:      {Type: field.TypeString, Column: country.FieldCode},
			country.FieldShort:     {Type: field.TypeString, Column: country.FieldShort},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   lang.Table,
			Columns: lang.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: lang.FieldID,
			},
		},
		Type: "Lang",
		Fields: map[string]*sqlgraph.FieldSpec{
			lang.FieldCreatedAt: {Type: field.TypeUint32, Column: lang.FieldCreatedAt},
			lang.FieldUpdatedAt: {Type: field.TypeUint32, Column: lang.FieldUpdatedAt},
			lang.FieldDeletedAt: {Type: field.TypeUint32, Column: lang.FieldDeletedAt},
			lang.FieldLang:      {Type: field.TypeString, Column: lang.FieldLang},
			lang.FieldLogo:      {Type: field.TypeString, Column: lang.FieldLogo},
			lang.FieldName:      {Type: field.TypeString, Column: lang.FieldName},
			lang.FieldShort:     {Type: field.TypeString, Column: lang.FieldShort},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *CountryQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CountryQuery builder.
func (cq *CountryQuery) Filter() *CountryFilter {
	return &CountryFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CountryMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CountryMutation builder.
func (m *CountryMutation) Filter() *CountryFilter {
	return &CountryFilter{config: m.config, predicateAdder: m}
}

// CountryFilter provides a generic filtering capability at runtime for CountryQuery.
type CountryFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CountryFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CountryFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(country.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CountryFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(country.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CountryFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(country.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CountryFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(country.FieldDeletedAt))
}

// WhereCountry applies the entql string predicate on the country field.
func (f *CountryFilter) WhereCountry(p entql.StringP) {
	f.Where(p.Field(country.FieldCountry))
}

// WhereFlag applies the entql string predicate on the flag field.
func (f *CountryFilter) WhereFlag(p entql.StringP) {
	f.Where(p.Field(country.FieldFlag))
}

// WhereCode applies the entql string predicate on the code field.
func (f *CountryFilter) WhereCode(p entql.StringP) {
	f.Where(p.Field(country.FieldCode))
}

// WhereShort applies the entql string predicate on the short field.
func (f *CountryFilter) WhereShort(p entql.StringP) {
	f.Where(p.Field(country.FieldShort))
}

// addPredicate implements the predicateAdder interface.
func (lq *LangQuery) addPredicate(pred func(s *sql.Selector)) {
	lq.predicates = append(lq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the LangQuery builder.
func (lq *LangQuery) Filter() *LangFilter {
	return &LangFilter{config: lq.config, predicateAdder: lq}
}

// addPredicate implements the predicateAdder interface.
func (m *LangMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the LangMutation builder.
func (m *LangMutation) Filter() *LangFilter {
	return &LangFilter{config: m.config, predicateAdder: m}
}

// LangFilter provides a generic filtering capability at runtime for LangQuery.
type LangFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *LangFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *LangFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(lang.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *LangFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(lang.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *LangFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(lang.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *LangFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(lang.FieldDeletedAt))
}

// WhereLang applies the entql string predicate on the lang field.
func (f *LangFilter) WhereLang(p entql.StringP) {
	f.Where(p.Field(lang.FieldLang))
}

// WhereLogo applies the entql string predicate on the logo field.
func (f *LangFilter) WhereLogo(p entql.StringP) {
	f.Where(p.Field(lang.FieldLogo))
}

// WhereName applies the entql string predicate on the name field.
func (f *LangFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(lang.FieldName))
}

// WhereShort applies the entql string predicate on the short field.
func (f *LangFilter) WhereShort(p entql.StringP) {
	f.Where(p.Field(lang.FieldShort))
}
