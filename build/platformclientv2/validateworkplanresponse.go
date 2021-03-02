package platformclientv2
import (
	"encoding/json"
)

// Validateworkplanresponse
type Validateworkplanresponse struct { 
	// WorkPlan - The work plan reference associated with this response
	WorkPlan *Workplanreference `json:"workPlan,omitempty"`


	// Valid - Whether the work plan is valid or not
	Valid *bool `json:"valid,omitempty"`


	// Messages - Validation messages for this work plan
	Messages *Validateworkplanmessages `json:"messages,omitempty"`

}

// String returns a JSON representation of the model
func (o *Validateworkplanresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
