package platformclientv2
import (
	"encoding/json"
)

// Shifttradelistresponse
type Shifttradelistresponse struct { 
	// Entities
	Entities *[]Shifttraderesponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shifttradelistresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
