package platformclientv2
import (
	"encoding/json"
)

// Validateaddressresponse
type Validateaddressresponse struct { 
	// Valid - Was the passed in address valid
	Valid *bool `json:"valid,omitempty"`


	// Response - Subscriber schema
	Response *Subscriberresponse `json:"response,omitempty"`

}

// String returns a JSON representation of the model
func (o *Validateaddressresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
