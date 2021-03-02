package platformclientv2
import (
	"encoding/json"
)

// Buaveragespeedofanswer - Service goal average speed of answer configuration
type Buaveragespeedofanswer struct { 
	// Include - Whether to include average speed of answer (ASA) in the associated configuration
	Include *bool `json:"include,omitempty"`


	// Seconds - The target average speed of answer (ASA) in seconds. Required if include == true
	Seconds *int `json:"seconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buaveragespeedofanswer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
