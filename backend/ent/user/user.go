// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldIsSuperuser holds the string denoting the is_superuser field in the database.
	FieldIsSuperuser = "is_superuser"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeAuthTokens holds the string denoting the auth_tokens edge name in mutations.
	EdgeAuthTokens = "auth_tokens"
	// Table holds the table name of the user in the database.
	Table = "users"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "users"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_users"
	// AuthTokensTable is the table that holds the auth_tokens relation/edge.
	AuthTokensTable = "auth_tokens"
	// AuthTokensInverseTable is the table name for the AuthTokens entity.
	// It exists in this package in order to avoid circular dependency with the "authtokens" package.
	AuthTokensInverseTable = "auth_tokens"
	// AuthTokensColumn is the table column denoting the auth_tokens relation/edge.
	AuthTokensColumn = "user_auth_tokens"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldEmail,
	FieldPassword,
	FieldIsSuperuser,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_users",
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultIsSuperuser holds the default value on creation for the "is_superuser" field.
	DefaultIsSuperuser bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
