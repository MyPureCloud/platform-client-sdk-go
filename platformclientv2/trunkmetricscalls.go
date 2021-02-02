package platformclientv2
import (
	"encoding/json"
)

// Trunkmetricscalls
type Trunkmetricscalls struct { 
	// InboundCallCount
	InboundCallCount *int `json:"inboundCallCount,omitempty"`


	// OutboundCallCount
	OutboundCallCount *int `json:"outboundCallCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricscalls) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
