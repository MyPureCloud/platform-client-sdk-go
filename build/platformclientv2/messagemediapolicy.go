package platformclientv2
import (
	"encoding/json"
)

// Messagemediapolicy
type Messagemediapolicy struct { 
	// Actions - Actions applied when specified conditions are met
	Actions *Policyactions `json:"actions,omitempty"`


	// Conditions - Conditions for when actions should be applied
	Conditions *Messagemediapolicyconditions `json:"conditions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messagemediapolicy) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
