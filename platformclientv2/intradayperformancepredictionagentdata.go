package platformclientv2
import (
	"encoding/json"
)

// Intradayperformancepredictionagentdata
type Intradayperformancepredictionagentdata struct { 
	// InteractingTimeSeconds - The total time spent interacting in seconds for all agents in this group
	InteractingTimeSeconds *float64 `json:"interactingTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayperformancepredictionagentdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
