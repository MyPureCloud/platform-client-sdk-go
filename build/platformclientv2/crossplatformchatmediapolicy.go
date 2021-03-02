package platformclientv2
import (
	"encoding/json"
)

// Crossplatformchatmediapolicy
type Crossplatformchatmediapolicy struct { 
	// Actions - Actions applied when specified conditions are met
	Actions *Crossplatformpolicyactions `json:"actions,omitempty"`


	// Conditions - Conditions for when actions should be applied
	Conditions *Chatmediapolicyconditions `json:"conditions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Crossplatformchatmediapolicy) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
