package platformclientv2
import (
	"encoding/json"
)

// Workplanlistresponse
type Workplanlistresponse struct { 
	// Entities
	Entities *[]Workplanlistitemresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanlistresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
