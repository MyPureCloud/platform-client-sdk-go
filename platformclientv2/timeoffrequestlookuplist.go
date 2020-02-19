package platformclientv2
import (
	"encoding/json"
)

// Timeoffrequestlookuplist
type Timeoffrequestlookuplist struct { 
	// Entities - List of time off request look up objects
	Entities *[]Timeoffrequestlookup `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestlookuplist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
