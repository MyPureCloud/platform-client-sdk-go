package platformclientv2
import (
	"encoding/json"
)

// Buabandonrate - Service goal abandon rate configuration
type Buabandonrate struct { 
	// Include - Whether to include abandon rate in the associated configuration
	Include *bool `json:"include,omitempty"`


	// Percent - Abandon rate percent goal. Required if include == true
	Percent *int `json:"percent,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buabandonrate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
