package platformclientv2
import (
	"time"
	"encoding/json"
)

// Phonechangetopiclinestatus
type Phonechangetopiclinestatus struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Reachable
	Reachable *bool `json:"reachable,omitempty"`


	// AddressOfRecord
	AddressOfRecord *string `json:"addressOfRecord,omitempty"`


	// ContactAddresses
	ContactAddresses *[]string `json:"contactAddresses,omitempty"`


	// ReachableStateTime
	ReachableStateTime *time.Time `json:"reachableStateTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonechangetopiclinestatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
