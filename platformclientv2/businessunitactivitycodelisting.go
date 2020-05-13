package platformclientv2
import (
	"encoding/json"
)

// Businessunitactivitycodelisting - List of BusinessUnitActivityCode
type Businessunitactivitycodelisting struct { 
	// Entities
	Entities *[]Businessunitactivitycode `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Businessunitactivitycodelisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
