package platformclientv2
import (
	"time"
	"encoding/json"
)

// Routingskill
type Routingskill struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the skill.
	Name *string `json:"name,omitempty"`


	// DateModified - Date last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// State - The current state for this skill.
	State *string `json:"state,omitempty"`


	// Version - Required when updating. Version must be the current version. Only the system can assign version.
	Version *string `json:"version,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routingskill) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
