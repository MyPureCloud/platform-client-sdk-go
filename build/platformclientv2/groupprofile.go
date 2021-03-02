package platformclientv2
import (
	"time"
	"encoding/json"
)

// Groupprofile
type Groupprofile struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// State - The state of the user resource
	State *string `json:"state,omitempty"`


	// DateModified - Datetime of the last modification. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - The version of the group resource
	Version *int `json:"version,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Groupprofile) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
