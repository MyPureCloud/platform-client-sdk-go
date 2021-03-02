package platformclientv2
import (
	"encoding/json"
)

// Crossplatformemailmediapolicy
type Crossplatformemailmediapolicy struct { 
	// Actions - Actions applied when specified conditions are met
	Actions *Crossplatformpolicyactions `json:"actions,omitempty"`


	// Conditions - Conditions for when actions should be applied
	Conditions *Emailmediapolicyconditions `json:"conditions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Crossplatformemailmediapolicy) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
