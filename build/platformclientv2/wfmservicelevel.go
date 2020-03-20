package platformclientv2
import (
	"encoding/json"
)

// Wfmservicelevel - Service goal service level configuration
type Wfmservicelevel struct { 
	// Include - Whether to include service level targets in the associated configuration
	Include *bool `json:"include,omitempty"`


	// Percent - Service level target percent answered. Required if include == true
	Percent *int32 `json:"percent,omitempty"`


	// Seconds - Service level target answer time. Required if include == true
	Seconds *int32 `json:"seconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmservicelevel) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
