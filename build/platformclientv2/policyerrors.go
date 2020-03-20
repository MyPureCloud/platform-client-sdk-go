package platformclientv2
import (
	"encoding/json"
)

// Policyerrors
type Policyerrors struct { 
	// PolicyErrorMessages
	PolicyErrorMessages *[]Policyerrormessage `json:"policyErrorMessages,omitempty"`

}

// String returns a JSON representation of the model
func (o *Policyerrors) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
