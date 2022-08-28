/*
 * storyden
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Thread struct {

	Id string `json:"id"`
	// The time the resource was created.
	CreatedAt time.Time `json:"createdAt"`
	// The time the resource was updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// The time the resource was soft-deleted.
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	// The title of the thread.
	Title string `json:"title"`
	// A URL friendly slug which is prepended with the post ID for uniqueness and sortability. 
	Slug string `json:"slug"`
	// A short version of the thread's body text for use in previews. 
	Short string `json:"short"`
	// Whether the thread is pinned in this category.
	Pinned bool `json:"pinned"`

	Author *ProfileReference `json:"author"`
	// A list of tags associated with the thread.
	Tags []string `json:"tags"`
	// The number of posts under this thread.
	Posts int32 `json:"posts"`

	Category *Category `json:"category"`
	// A list of reactions this post has had from people.
	Reacts []React `json:"reacts"`
}
