package platformclientv2
import (
	"encoding/json"
)

// Trunkmetricsqos
type Trunkmetricsqos struct { 
	// MismatchCount - Total number of QoS mismatches over the course of the last 24-hour period (sliding window).
	MismatchCount *int32 `json:"mismatchCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricsqos) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
