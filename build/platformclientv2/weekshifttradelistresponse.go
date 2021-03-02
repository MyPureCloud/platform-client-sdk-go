package platformclientv2
import (
	"encoding/json"
)

// Weekshifttradelistresponse
type Weekshifttradelistresponse struct { 
	// Entities
	Entities *[]Weekshifttraderesponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Weekshifttradelistresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
