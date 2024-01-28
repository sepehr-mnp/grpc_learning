// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmailAddress holds the string denoting the email_address field in the database.
	FieldEmailAddress = "email_address"
	// FieldAlias holds the string denoting the alias field in the database.
	FieldAlias = "alias"
	// EdgeAdministered holds the string denoting the administered edge name in mutations.
	EdgeAdministered = "administered"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AdministeredTable is the table that holds the administered relation/edge.
	AdministeredTable = "categories"
	// AdministeredInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	AdministeredInverseTable = "categories"
	// AdministeredColumn is the table column denoting the administered relation/edge.
	AdministeredColumn = "category_admin"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmailAddress,
	FieldAlias,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmailAddress orders the results by the email_address field.
func ByEmailAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailAddress, opts...).ToFunc()
}

// ByAlias orders the results by the alias field.
func ByAlias(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlias, opts...).ToFunc()
}

// ByAdministeredCount orders the results by administered count.
func ByAdministeredCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAdministeredStep(), opts...)
	}
}

// ByAdministered orders the results by administered terms.
func ByAdministered(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAdministeredStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAdministeredStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AdministeredInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AdministeredTable, AdministeredColumn),
	)
}