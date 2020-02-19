package platformclientv2
import (
	"encoding/json"
)

// Wfmabandonrate - Service goal abandon rate configuration
type Wfmabandonrate struct { 
	// Include - Whether to include abandon rate in the associated configuration
	Include *bool `json:"include,omitempty"`


	// Percent - Abandon rate percent goal. Required if include == true
	Percent *int32 `json:"percent,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmabandonrate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
