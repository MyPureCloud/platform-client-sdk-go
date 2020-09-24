package platformclientv2
import (
	"encoding/json"
)

// Developmentactivityaggregatequeryresponsemetric
type Developmentactivityaggregatequeryresponsemetric struct { 
	// Metric - The metric this applies to
	Metric *string `json:"metric,omitempty"`


	// Stats - The aggregated values for this metric
	Stats *Developmentactivityaggregatequeryresponsestatistics `json:"stats,omitempty"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsemetric) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
