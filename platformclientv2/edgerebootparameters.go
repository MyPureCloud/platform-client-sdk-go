package platformclientv2
import (
	"encoding/json"
)

// Edgerebootparameters
type Edgerebootparameters struct { 
	// CallDrainingWaitTimeSeconds - The number of seconds to wait for call draining to complete before initiating the reboot. A value of 0 will prevent call draining and all calls will disconnect immediately.
	CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgerebootparameters) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
