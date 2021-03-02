package platformclientv2
import (
	"encoding/json"
)

// Emailmediapolicy
type Emailmediapolicy struct { 
	// Actions - Actions applied when specified conditions are met
	Actions *Policyactions `json:"actions,omitempty"`


	// Conditions - Conditions for when actions should be applied
	Conditions *Emailmediapolicyconditions `json:"conditions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Emailmediapolicy) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
