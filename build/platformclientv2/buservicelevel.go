package platformclientv2
import (
	"encoding/json"
)

// Buservicelevel - Service goal service level configuration
type Buservicelevel struct { 
	// Include - Whether to include service level targets in the associated configuration
	Include *bool `json:"include,omitempty"`


	// Percent - Service level target percent answered. Required if include == true
	Percent *int `json:"percent,omitempty"`


	// Seconds - Service level target answer time. Required if include == true
	Seconds *int `json:"seconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buservicelevel) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
