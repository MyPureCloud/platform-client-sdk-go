package platformclientv2
import (
	"time"
	"encoding/json"
)

// Attribute
type Attribute struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The attribute name.
	Name *string `json:"name,omitempty"`


	// Version
	Version *int32 `json:"version,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// CreatedBy
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ModifiedBy
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Attribute) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
