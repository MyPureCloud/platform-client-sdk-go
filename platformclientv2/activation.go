package platformclientv2
import (
	"encoding/json"
)

// Activation
type Activation struct { 
	// VarType - Type of activation.
	VarType *string `json:"type,omitempty"`


	// DelayInSeconds - Activation delay time amount.
	DelayInSeconds *int `json:"delayInSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Activation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
