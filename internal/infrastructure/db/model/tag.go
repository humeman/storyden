// Code generated by ent, DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/internal/infrastructure/db/model/tag"
	"github.com/rs/xid"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TagQuery when eager-loading is set.
	Edges TagEdges `json:"edges"`
}

// TagEdges holds the relations/edges for other nodes in the graph.
type TagEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// Accounts holds the value of the accounts edge.
	Accounts []*Account `json:"accounts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// AccountsOrErr returns the Accounts value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) AccountsOrErr() ([]*Account, error) {
	if e.loadedTypes[1] {
		return e.Accounts, nil
	}
	return nil, &NotLoadedError{edge: "accounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldName:
			values[i] = new(sql.NullString)
		case tag.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case tag.FieldID:
			values[i] = new(xid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case tag.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case tag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		}
	}
	return nil
}

// QueryPosts queries the "posts" edge of the Tag entity.
func (t *Tag) QueryPosts() *PostQuery {
	return (&TagClient{config: t.config}).QueryPosts(t)
}

// QueryAccounts queries the "accounts" edge of the Tag entity.
func (t *Tag) QueryAccounts() *AccountQuery {
	return (&TagClient{config: t.config}).QueryAccounts(t)
}

// Update returns a builder for updating this Tag.
// Note that you need to call Tag.Unwrap() before calling this method if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return (&TagClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Tag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("model: Tag is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Tags is a parsable slice of Tag.
type Tags []*Tag

func (t Tags) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
