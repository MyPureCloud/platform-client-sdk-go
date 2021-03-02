package platformclientv2
import (
	"encoding/json"
)

// Workplanrotationlistresponse
type Workplanrotationlistresponse struct { 
	// Entities
	Entities *[]Workplanrotationresponse `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanrotationlistresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
