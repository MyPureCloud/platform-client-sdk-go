package platformclientv2
import (
	"encoding/json"
)

// Edgeservicestaterequest
type Edgeservicestaterequest struct { 
	// InService - A boolean that sets the Edge in-service or out-of-service.
	InService *bool `json:"inService,omitempty"`


	// CallDrainingWaitTimeSeconds - The number of seconds to wait for call draining to complete before initiating the reboot. A value of 0 will prevent call draining and all calls will disconnect immediately.
	CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgeservicestaterequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
