// Code generated by ent, DO NOT EDIT.

package roleschema

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the roleschema type in the database.
	Label = "role_schema"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the roleschema in the database.
	Table = "Role"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "Role"
	// UserInverseTable is the table name for the UserSchema entity.
	// It exists in this package in order to avoid circular dependency with the "userschema" package.
	UserInverseTable = "User"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_schema_roles"
)

// Columns holds all SQL columns for roleschema fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "Role"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_schema_roles",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the RoleSchema queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}