package platformclientv2
import (
	"encoding/json"
)

// Workplanpatternresponse
type Workplanpatternresponse struct { 
	// WorkPlans - List of work plans in order of rotation on a weekly basis
	WorkPlans *[]Workplanreference `json:"workPlans,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workplanpatternresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
