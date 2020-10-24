package platformclientv2
import (
	"time"
	"encoding/json"
)

// Contactlistfilter
type Contactlistfilter struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the list.
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int32 `json:"version,omitempty"`


	// ContactList - The contact list the filter is based on.
	ContactList *Domainentityref `json:"contactList,omitempty"`


	// Clauses - Groups of conditions to filter the contacts by.
	Clauses *[]Contactlistfilterclause `json:"clauses,omitempty"`


	// FilterType - How to join clauses together.
	FilterType *string `json:"filterType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactlistfilter) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
