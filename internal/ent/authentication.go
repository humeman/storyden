// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/authentication"
	"github.com/rs/xid"
)

// Authentication is the model entity for the Authentication schema.
type Authentication struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The authentication service name, such as GitHub, Twitter, Discord, etc. Or, 'password' for password auth and 'api_token' for token auth
	Service string `json:"service,omitempty"`
	// The identifier, usually a user/account ID on some OAuth service or API token name. If it's a password, this is blank.
	Identifier string `json:"identifier,omitempty"`
	// The actual authentication token/password/key/etc. If OAuth, it'll be the access_token value, if it's a password, a hash and if it's an api_token type then the API token string.
	Token string `json:"-"`
	// A human-readable name for the authentication method. For WebAuthn, this may be the device OS or nickname.
	Name *string `json:"name,omitempty"`
	// Any necessary metadata specific to the authentication method.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuthenticationQuery when eager-loading is set.
	Edges                  AuthenticationEdges `json:"edges"`
	account_authentication *xid.ID
	selectValues           sql.SelectValues
}

// AuthenticationEdges holds the relations/edges for other nodes in the graph.
type AuthenticationEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// EmailAddress holds the value of the email_address edge.
	EmailAddress []*Email `json:"email_address,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuthenticationEdges) AccountOrErr() (*Account, error) {
	if e.Account != nil {
		return e.Account, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: account.Label}
	}
	return nil, &NotLoadedError{edge: "account"}
}

// EmailAddressOrErr returns the EmailAddress value or an error if the edge
// was not loaded in eager-loading.
func (e AuthenticationEdges) EmailAddressOrErr() ([]*Email, error) {
	if e.loadedTypes[1] {
		return e.EmailAddress, nil
	}
	return nil, &NotLoadedError{edge: "email_address"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Authentication) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case authentication.FieldMetadata:
			values[i] = new([]byte)
		case authentication.FieldService, authentication.FieldIdentifier, authentication.FieldToken, authentication.FieldName:
			values[i] = new(sql.NullString)
		case authentication.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case authentication.FieldID:
			values[i] = new(xid.ID)
		case authentication.ForeignKeys[0]: // account_authentication
			values[i] = &sql.NullScanner{S: new(xid.ID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Authentication fields.
func (a *Authentication) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authentication.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case authentication.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case authentication.FieldService:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service", values[i])
			} else if value.Valid {
				a.Service = value.String
			}
		case authentication.FieldIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[i])
			} else if value.Valid {
				a.Identifier = value.String
			}
		case authentication.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				a.Token = value.String
			}
		case authentication.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = new(string)
				*a.Name = value.String
			}
		case authentication.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case authentication.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field account_authentication", values[i])
			} else if value.Valid {
				a.account_authentication = new(xid.ID)
				*a.account_authentication = *value.S.(*xid.ID)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Authentication.
// This includes values selected through modifiers, order, etc.
func (a *Authentication) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryAccount queries the "account" edge of the Authentication entity.
func (a *Authentication) QueryAccount() *AccountQuery {
	return NewAuthenticationClient(a.config).QueryAccount(a)
}

// QueryEmailAddress queries the "email_address" edge of the Authentication entity.
func (a *Authentication) QueryEmailAddress() *EmailQuery {
	return NewAuthenticationClient(a.config).QueryEmailAddress(a)
}

// Update returns a builder for updating this Authentication.
// Note that you need to call Authentication.Unwrap() before calling this method if this Authentication
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Authentication) Update() *AuthenticationUpdateOne {
	return NewAuthenticationClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Authentication entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Authentication) Unwrap() *Authentication {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Authentication is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Authentication) String() string {
	var builder strings.Builder
	builder.WriteString("Authentication(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("service=")
	builder.WriteString(a.Service)
	builder.WriteString(", ")
	builder.WriteString("identifier=")
	builder.WriteString(a.Identifier)
	builder.WriteString(", ")
	builder.WriteString("token=<sensitive>")
	builder.WriteString(", ")
	if v := a.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", a.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// Authentications is a parsable slice of Authentication.
type Authentications []*Authentication
