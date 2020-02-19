package platformclientv2
import (
	"encoding/json"
)

// Intradayforecastdata
type Intradayforecastdata struct { 
	// Offered - The number of interactions routed into the queue for the given media type(s) for an agent to answer
	Offered *float64 `json:"offered,omitempty"`


	// AverageTalkTimeSeconds - The average time in seconds an agent spends interacting with a customer
	AverageTalkTimeSeconds *float64 `json:"averageTalkTimeSeconds,omitempty"`


	// AverageAfterCallWorkSeconds - The average time in seconds spent in after-call work. After-call work is the work that an agent performs immediately following an interaction
	AverageAfterCallWorkSeconds *float64 `json:"averageAfterCallWorkSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayforecastdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
