package platformclientv2
import (
	"encoding/json"
)

// Intradayhistoricalagentdata
type Intradayhistoricalagentdata struct { 
	// OnQueueTimeSeconds - The total on-queue time in seconds for all agents in this group
	OnQueueTimeSeconds *float64 `json:"onQueueTimeSeconds,omitempty"`


	// InteractingTimeSeconds - The total time spent interacting in seconds for all agents in this group
	InteractingTimeSeconds *float64 `json:"interactingTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayhistoricalagentdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
