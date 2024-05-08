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
	"github.com/Southclaws/storyden/internal/ent/item"
	"github.com/rs/xid"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Content holds the value of the "content" field.
	Content *string `json:"content,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID xid.ID `json:"account_id,omitempty"`
	// Visibility holds the value of the "visibility" field.
	Visibility item.Visibility `json:"visibility,omitempty"`
	// Properties holds the value of the "properties" field.
	Properties any `json:"properties,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemQuery when eager-loading is set.
	Edges        ItemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ItemEdges holds the relations/edges for other nodes in the graph.
type ItemEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Account `json:"owner,omitempty"`
	// Clusters holds the value of the clusters edge.
	Clusters []*Cluster `json:"clusters,omitempty"`
	// Assets holds the value of the assets edge.
	Assets []*Asset `json:"assets,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// Links holds the value of the links edge.
	Links []*Link `json:"links,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemEdges) OwnerOrErr() (*Account, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: account.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ClustersOrErr returns the Clusters value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) ClustersOrErr() ([]*Cluster, error) {
	if e.loadedTypes[1] {
		return e.Clusters, nil
	}
	return nil, &NotLoadedError{edge: "clusters"}
}

// AssetsOrErr returns the Assets value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) AssetsOrErr() ([]*Asset, error) {
	if e.loadedTypes[2] {
		return e.Assets, nil
	}
	return nil, &NotLoadedError{edge: "assets"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[3] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// LinksOrErr returns the Links value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) LinksOrErr() ([]*Link, error) {
	if e.loadedTypes[4] {
		return e.Links, nil
	}
	return nil, &NotLoadedError{edge: "links"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case item.FieldProperties:
			values[i] = new([]byte)
		case item.FieldName, item.FieldSlug, item.FieldDescription, item.FieldContent, item.FieldVisibility:
			values[i] = new(sql.NullString)
		case item.FieldCreatedAt, item.FieldUpdatedAt, item.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case item.FieldID, item.FieldAccountID:
			values[i] = new(xid.ID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case item.FieldID:
			if value, ok := values[j].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case item.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case item.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case item.FieldDeletedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[j])
			} else if value.Valid {
				i.DeletedAt = new(time.Time)
				*i.DeletedAt = value.Time
			}
		case item.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case item.FieldSlug:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[j])
			} else if value.Valid {
				i.Slug = value.String
			}
		case item.FieldDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[j])
			} else if value.Valid {
				i.Description = value.String
			}
		case item.FieldContent:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[j])
			} else if value.Valid {
				i.Content = new(string)
				*i.Content = value.String
			}
		case item.FieldAccountID:
			if value, ok := values[j].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[j])
			} else if value != nil {
				i.AccountID = *value
			}
		case item.FieldVisibility:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field visibility", values[j])
			} else if value.Valid {
				i.Visibility = item.Visibility(value.String)
			}
		case item.FieldProperties:
			if value, ok := values[j].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field properties", values[j])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &i.Properties); err != nil {
					return fmt.Errorf("unmarshal field properties: %w", err)
				}
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Item.
// This includes values selected through modifiers, order, etc.
func (i *Item) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Item entity.
func (i *Item) QueryOwner() *AccountQuery {
	return NewItemClient(i.config).QueryOwner(i)
}

// QueryClusters queries the "clusters" edge of the Item entity.
func (i *Item) QueryClusters() *ClusterQuery {
	return NewItemClient(i.config).QueryClusters(i)
}

// QueryAssets queries the "assets" edge of the Item entity.
func (i *Item) QueryAssets() *AssetQuery {
	return NewItemClient(i.config).QueryAssets(i)
}

// QueryTags queries the "tags" edge of the Item entity.
func (i *Item) QueryTags() *TagQuery {
	return NewItemClient(i.config).QueryTags(i)
}

// QueryLinks queries the "links" edge of the Item entity.
func (i *Item) QueryLinks() *LinkQuery {
	return NewItemClient(i.config).QueryLinks(i)
}

// Update returns a builder for updating this Item.
// Note that you need to call Item.Unwrap() before calling this method if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return NewItemClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := i.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(i.Slug)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(i.Description)
	builder.WriteString(", ")
	if v := i.Content; v != nil {
		builder.WriteString("content=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", i.AccountID))
	builder.WriteString(", ")
	builder.WriteString("visibility=")
	builder.WriteString(fmt.Sprintf("%v", i.Visibility))
	builder.WriteString(", ")
	builder.WriteString("properties=")
	builder.WriteString(fmt.Sprintf("%v", i.Properties))
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item
