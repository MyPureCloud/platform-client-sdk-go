package platformclientv2
import (
	"encoding/json"
)

// Wfmintradaydataupdatetopicintradayhistoricalagentdata
type Wfmintradaydataupdatetopicintradayhistoricalagentdata struct { 
	// OnQueueTimeSeconds
	OnQueueTimeSeconds *float32 `json:"onQueueTimeSeconds,omitempty"`


	// InteractingTimeSeconds
	InteractingTimeSeconds *float32 `json:"interactingTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayhistoricalagentdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
