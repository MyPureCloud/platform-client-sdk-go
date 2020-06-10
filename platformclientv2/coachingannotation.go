package platformclientv2
import (
	"time"
	"encoding/json"
)

// Coachingannotation
type Coachingannotation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// CreatedBy - The user who created the annotation.
	CreatedBy *Userreference `json:"createdBy,omitempty"`


	// DateCreated - The date/time the annotation was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ModifiedBy - The last user to modify the annotation.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`


	// DateModified - The date/time the annotation was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Text - The text of the annotation.
	Text *string `json:"text,omitempty"`


	// IsDeleted - Flag indicating whether the annotation is deleted.
	IsDeleted *bool `json:"isDeleted,omitempty"`


	// AccessType - Determines the permissions required to view this item.
	AccessType *string `json:"accessType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coachingannotation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
