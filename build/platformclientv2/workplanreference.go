package platformclientv2
import (
	"encoding/json"
)

// Workplanreference - Work plan information
type Workplanreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// ManagementUnit - The management unit to which this work plan belongs.  Nullable in some routes
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanreference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
