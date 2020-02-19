package platformclientv2
import (
	"encoding/json"
)

// Shifttradematchessummaryresponse
type Shifttradematchessummaryresponse struct { 
	// Entities
	Entities *[]Weekshifttradematchessummaryresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Shifttradematchessummaryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
