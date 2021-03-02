package platformclientv2
import (
	"time"
	"encoding/json"
)

// Linestatus
type Linestatus struct { 
	// Id - The id of this line
	Id *string `json:"id,omitempty"`


	// Reachable - Indicates whether the edge can reach the line.
	Reachable *bool `json:"reachable,omitempty"`


	// AddressOfRecord - The line's address of record.
	AddressOfRecord *string `json:"addressOfRecord,omitempty"`


	// ContactAddresses - The addresses used to contact the line.
	ContactAddresses *[]string `json:"contactAddresses,omitempty"`


	// ReachableStateTime - The time the line entered its current reachable state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReachableStateTime *time.Time `json:"reachableStateTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Linestatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
