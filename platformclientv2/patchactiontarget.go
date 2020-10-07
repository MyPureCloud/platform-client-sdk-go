package platformclientv2
import (
	"encoding/json"
)

// Patchactiontarget
type Patchactiontarget struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ServiceLevel - Service Level of the action target. Chat offers for the target will be throttled with the aim of achieving this service level.
	ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchactiontarget) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
