package platformclientv2
import (
	"encoding/json"
)

// Developmentactivityaggregatequeryresponsestatistics
type Developmentactivityaggregatequeryresponsestatistics struct { 
	// Count - The count for this metric
	Count *int32 `json:"count,omitempty"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsestatistics) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
