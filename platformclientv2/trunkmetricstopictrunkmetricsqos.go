package platformclientv2
import (
	"encoding/json"
)

// Trunkmetricstopictrunkmetricsqos
type Trunkmetricstopictrunkmetricsqos struct { 
	// MismatchCount
	MismatchCount *int `json:"mismatchCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetricsqos) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
