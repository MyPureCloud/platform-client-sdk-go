package platformclientv2
import (
	"encoding/json"
)

// Dialersequencescheduleconfigchangescheduleinterval
type Dialersequencescheduleconfigchangescheduleinterval struct { 
	// Start
	Start *string `json:"start,omitempty"`


	// End
	End *string `json:"end,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialersequencescheduleconfigchangescheduleinterval) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
