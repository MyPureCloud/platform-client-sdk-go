package platformclientv2
import (
	"encoding/json"
)

// Authzdivision
type Authzdivision struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description - A helpful description for the division.
	Description *string `json:"description,omitempty"`


	// HomeDivision - A flag indicating whether this division is the \"Home\" (default) division. Cannot be modified and any supplied value will be ignored on create or update.
	HomeDivision *bool `json:"homeDivision,omitempty"`


	// ObjectCounts - A count of objects in this division, grouped by type.
	ObjectCounts *map[string]int64 `json:"objectCounts,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Authzdivision) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
