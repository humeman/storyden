// Code generated by ent, DO NOT EDIT.

package account

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...xid.ID) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id xid.ID) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldDeletedAt, v))
}

// Handle applies equality check predicate on the "handle" field. It's identical to HandleEQ.
func Handle(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldHandle, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldName, v))
}

// Bio applies equality check predicate on the "bio" field. It's identical to BioEQ.
func Bio(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldBio, v))
}

// Admin applies equality check predicate on the "admin" field. It's identical to AdminEQ.
func Admin(v bool) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAdmin, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldDeletedAt))
}

// HandleEQ applies the EQ predicate on the "handle" field.
func HandleEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldHandle, v))
}

// HandleNEQ applies the NEQ predicate on the "handle" field.
func HandleNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldHandle, v))
}

// HandleIn applies the In predicate on the "handle" field.
func HandleIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldHandle, vs...))
}

// HandleNotIn applies the NotIn predicate on the "handle" field.
func HandleNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldHandle, vs...))
}

// HandleGT applies the GT predicate on the "handle" field.
func HandleGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldHandle, v))
}

// HandleGTE applies the GTE predicate on the "handle" field.
func HandleGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldHandle, v))
}

// HandleLT applies the LT predicate on the "handle" field.
func HandleLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldHandle, v))
}

// HandleLTE applies the LTE predicate on the "handle" field.
func HandleLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldHandle, v))
}

// HandleContains applies the Contains predicate on the "handle" field.
func HandleContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldHandle, v))
}

// HandleHasPrefix applies the HasPrefix predicate on the "handle" field.
func HandleHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldHandle, v))
}

// HandleHasSuffix applies the HasSuffix predicate on the "handle" field.
func HandleHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldHandle, v))
}

// HandleEqualFold applies the EqualFold predicate on the "handle" field.
func HandleEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldHandle, v))
}

// HandleContainsFold applies the ContainsFold predicate on the "handle" field.
func HandleContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldHandle, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldName, v))
}

// BioEQ applies the EQ predicate on the "bio" field.
func BioEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldBio, v))
}

// BioNEQ applies the NEQ predicate on the "bio" field.
func BioNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldBio, v))
}

// BioIn applies the In predicate on the "bio" field.
func BioIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldBio, vs...))
}

// BioNotIn applies the NotIn predicate on the "bio" field.
func BioNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldBio, vs...))
}

// BioGT applies the GT predicate on the "bio" field.
func BioGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldBio, v))
}

// BioGTE applies the GTE predicate on the "bio" field.
func BioGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldBio, v))
}

// BioLT applies the LT predicate on the "bio" field.
func BioLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldBio, v))
}

// BioLTE applies the LTE predicate on the "bio" field.
func BioLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldBio, v))
}

// BioContains applies the Contains predicate on the "bio" field.
func BioContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldBio, v))
}

// BioHasPrefix applies the HasPrefix predicate on the "bio" field.
func BioHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldBio, v))
}

// BioHasSuffix applies the HasSuffix predicate on the "bio" field.
func BioHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldBio, v))
}

// BioIsNil applies the IsNil predicate on the "bio" field.
func BioIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldBio))
}

// BioNotNil applies the NotNil predicate on the "bio" field.
func BioNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldBio))
}

// BioEqualFold applies the EqualFold predicate on the "bio" field.
func BioEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldBio, v))
}

// BioContainsFold applies the ContainsFold predicate on the "bio" field.
func BioContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldBio, v))
}

// AdminEQ applies the EQ predicate on the "admin" field.
func AdminEQ(v bool) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAdmin, v))
}

// AdminNEQ applies the NEQ predicate on the "admin" field.
func AdminNEQ(v bool) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldAdmin, v))
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReacts applies the HasEdge predicate on the "reacts" edge.
func HasReacts() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReactsTable, ReactsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReactsWith applies the HasEdge predicate on the "reacts" edge with a given conditions (other predicates).
func HasReactsWith(preds ...predicate.React) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newReactsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newRolesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAuthentication applies the HasEdge predicate on the "authentication" edge.
func HasAuthentication() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AuthenticationTable, AuthenticationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthenticationWith applies the HasEdge predicate on the "authentication" edge with a given conditions (other predicates).
func HasAuthenticationWith(preds ...predicate.Authentication) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newAuthenticationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCollections applies the HasEdge predicate on the "collections" edge.
func HasCollections() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CollectionsTable, CollectionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCollectionsWith applies the HasEdge predicate on the "collections" edge with a given conditions (other predicates).
func HasCollectionsWith(preds ...predicate.Collection) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newCollectionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClusters applies the HasEdge predicate on the "clusters" edge.
func HasClusters() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClustersTable, ClustersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClustersWith applies the HasEdge predicate on the "clusters" edge with a given conditions (other predicates).
func HasClustersWith(preds ...predicate.Cluster) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newClustersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAssets applies the HasEdge predicate on the "assets" edge.
func HasAssets() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetsTable, AssetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetsWith applies the HasEdge predicate on the "assets" edge with a given conditions (other predicates).
func HasAssetsWith(preds ...predicate.Asset) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newAssetsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(sql.NotPredicates(p))
}
