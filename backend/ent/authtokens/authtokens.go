// Code generated by ent, DO NOT EDIT.

package authtokens

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the authtokens type in the database.
	Label = "auth_tokens"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the authtokens in the database.
	Table = "auth_tokens"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "auth_tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_auth_tokens"
)

// Columns holds all SQL columns for authtokens fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldToken,
	FieldExpiresAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "auth_tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_auth_tokens",
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultExpiresAt holds the default value on creation for the "expires_at" field.
	DefaultExpiresAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
