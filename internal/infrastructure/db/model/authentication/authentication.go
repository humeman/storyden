// Code generated by ent, DO NOT EDIT.

package authentication

import (
	"time"

	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the authentication type in the database.
	Label = "authentication"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldService holds the string denoting the service field in the database.
	FieldService = "service"
	// FieldIdentifier holds the string denoting the identifier field in the database.
	FieldIdentifier = "identifier"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the authentication in the database.
	Table = "authentications"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "authentications"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_authentication"
)

// Columns holds all SQL columns for authentication fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldService,
	FieldIdentifier,
	FieldToken,
	FieldMetadata,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "authentications"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_authentication",
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
	// ServiceValidator is a validator for the "service" field. It is called by the builders before save.
	ServiceValidator func(string) error
	// TokenValidator is a validator for the "token" field. It is called by the builders before save.
	TokenValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
