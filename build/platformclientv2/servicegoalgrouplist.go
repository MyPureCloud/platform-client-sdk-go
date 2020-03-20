package platformclientv2
import (
	"encoding/json"
)

// Servicegoalgrouplist - List of service goal groups
type Servicegoalgrouplist struct { 
	// Entities
	Entities *[]Servicegoalgroup `json:"entities,omitempty"`


	// Metadata - Version metadata for the list of service goal groups for the associated management unit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Servicegoalgrouplist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
