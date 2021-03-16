package platformclientv2
import (
	"encoding/json"
)

// Bulkidsrequest
type Bulkidsrequest struct { 
	// Entities
	Entities *[]Entity `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkidsrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
