package platformclientv2
import (
	"encoding/json"
)

// Wfmaveragespeedofanswer - Service goal average speed of answer configuration
type Wfmaveragespeedofanswer struct { 
	// Include - Whether to include average speed of answer (ASA) in the associated configuration
	Include *bool `json:"include,omitempty"`


	// Seconds - The target average speed of answer (ASA) in seconds. Required if include == true
	Seconds *int32 `json:"seconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmaveragespeedofanswer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
