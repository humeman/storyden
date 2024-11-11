// Code generated by ent, DO NOT EDIT.

package event

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldIndexedAt holds the string denoting the indexed_at field in the database.
	FieldIndexedAt = "indexed_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldParticipationPolicy holds the string denoting the participation_policy field in the database.
	FieldParticipationPolicy = "participation_policy"
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// FieldLocationType holds the string denoting the location_type field in the database.
	FieldLocationType = "location_type"
	// FieldLocationName holds the string denoting the location_name field in the database.
	FieldLocationName = "location_name"
	// FieldLocationAddress holds the string denoting the location_address field in the database.
	FieldLocationAddress = "location_address"
	// FieldLocationLatitude holds the string denoting the location_latitude field in the database.
	FieldLocationLatitude = "location_latitude"
	// FieldLocationLongitude holds the string denoting the location_longitude field in the database.
	FieldLocationLongitude = "location_longitude"
	// FieldLocationURL holds the string denoting the location_url field in the database.
	FieldLocationURL = "location_url"
	// FieldCapacity holds the string denoting the capacity field in the database.
	FieldCapacity = "capacity"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// EdgeParticipants holds the string denoting the participants edge name in mutations.
	EdgeParticipants = "participants"
	// EdgeThread holds the string denoting the thread edge name in mutations.
	EdgeThread = "thread"
	// EdgePrimaryImage holds the string denoting the primary_image edge name in mutations.
	EdgePrimaryImage = "primary_image"
	// Table holds the table name of the event in the database.
	Table = "events"
	// ParticipantsTable is the table that holds the participants relation/edge.
	ParticipantsTable = "event_participants"
	// ParticipantsInverseTable is the table name for the EventParticipant entity.
	// It exists in this package in order to avoid circular dependency with the "eventparticipant" package.
	ParticipantsInverseTable = "event_participants"
	// ParticipantsColumn is the table column denoting the participants relation/edge.
	ParticipantsColumn = "event_id"
	// ThreadTable is the table that holds the thread relation/edge.
	ThreadTable = "events"
	// ThreadInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	ThreadInverseTable = "posts"
	// ThreadColumn is the table column denoting the thread relation/edge.
	ThreadColumn = "post_event"
	// PrimaryImageTable is the table that holds the primary_image relation/edge.
	PrimaryImageTable = "events"
	// PrimaryImageInverseTable is the table name for the Asset entity.
	// It exists in this package in order to avoid circular dependency with the "asset" package.
	PrimaryImageInverseTable = "assets"
	// PrimaryImageColumn is the table column denoting the primary_image relation/edge.
	PrimaryImageColumn = "asset_event"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldIndexedAt,
	FieldName,
	FieldSlug,
	FieldDescription,
	FieldStartTime,
	FieldEndTime,
	FieldParticipationPolicy,
	FieldVisibility,
	FieldLocationType,
	FieldLocationName,
	FieldLocationAddress,
	FieldLocationLatitude,
	FieldLocationLongitude,
	FieldLocationURL,
	FieldCapacity,
	FieldMetadata,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"asset_event",
	"post_event",
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Visibility defines the type for the "visibility" enum field.
type Visibility string

// VisibilityDraft is the default value of the Visibility enum.
const DefaultVisibility = VisibilityDraft

// Visibility values.
const (
	VisibilityDraft     Visibility = "draft"
	VisibilityUnlisted  Visibility = "unlisted"
	VisibilityReview    Visibility = "review"
	VisibilityPublished Visibility = "published"
)

func (v Visibility) String() string {
	return string(v)
}

// VisibilityValidator is a validator for the "visibility" field enum values. It is called by the builders before save.
func VisibilityValidator(v Visibility) error {
	switch v {
	case VisibilityDraft, VisibilityUnlisted, VisibilityReview, VisibilityPublished:
		return nil
	default:
		return fmt.Errorf("event: invalid enum value for visibility field: %q", v)
	}
}

// OrderOption defines the ordering options for the Event queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByIndexedAt orders the results by the indexed_at field.
func ByIndexedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndexedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// ByParticipationPolicy orders the results by the participation_policy field.
func ByParticipationPolicy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParticipationPolicy, opts...).ToFunc()
}

// ByVisibility orders the results by the visibility field.
func ByVisibility(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisibility, opts...).ToFunc()
}

// ByLocationType orders the results by the location_type field.
func ByLocationType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationType, opts...).ToFunc()
}

// ByLocationName orders the results by the location_name field.
func ByLocationName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationName, opts...).ToFunc()
}

// ByLocationAddress orders the results by the location_address field.
func ByLocationAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationAddress, opts...).ToFunc()
}

// ByLocationLatitude orders the results by the location_latitude field.
func ByLocationLatitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationLatitude, opts...).ToFunc()
}

// ByLocationLongitude orders the results by the location_longitude field.
func ByLocationLongitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationLongitude, opts...).ToFunc()
}

// ByLocationURL orders the results by the location_url field.
func ByLocationURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationURL, opts...).ToFunc()
}

// ByCapacity orders the results by the capacity field.
func ByCapacity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCapacity, opts...).ToFunc()
}

// ByParticipantsCount orders the results by participants count.
func ByParticipantsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newParticipantsStep(), opts...)
	}
}

// ByParticipants orders the results by participants terms.
func ByParticipants(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParticipantsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByThreadField orders the results by thread field.
func ByThreadField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newThreadStep(), sql.OrderByField(field, opts...))
	}
}

// ByPrimaryImageField orders the results by primary_image field.
func ByPrimaryImageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrimaryImageStep(), sql.OrderByField(field, opts...))
	}
}
func newParticipantsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ParticipantsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ParticipantsTable, ParticipantsColumn),
	)
}
func newThreadStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ThreadInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ThreadTable, ThreadColumn),
	)
}
func newPrimaryImageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PrimaryImageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PrimaryImageTable, PrimaryImageColumn),
	)
}
