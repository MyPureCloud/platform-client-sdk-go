package platformclientv2
import (
	"encoding/json"
)

// Trunkbaseassignment
type Trunkbaseassignment struct { 
	// Family - The address family to use with the trunk base settings. 2=IPv4, 23=IPv6
	Family *int `json:"family,omitempty"`


	// TrunkBase - A trunk base settings reference.
	TrunkBase *Trunkbase `json:"trunkBase,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkbaseassignment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
