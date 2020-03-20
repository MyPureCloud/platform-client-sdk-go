package platformclientv2
import (
	"encoding/json"
)

// Domainphysicalcapabilities
type Domainphysicalcapabilities struct { 
	// Vlan
	Vlan *bool `json:"vlan,omitempty"`


	// Team
	Team *bool `json:"team,omitempty"`

}

// String returns a JSON representation of the model
func (o *Domainphysicalcapabilities) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
