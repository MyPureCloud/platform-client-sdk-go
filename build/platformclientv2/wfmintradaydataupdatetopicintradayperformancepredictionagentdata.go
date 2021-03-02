package platformclientv2
import (
	"encoding/json"
)

// Wfmintradaydataupdatetopicintradayperformancepredictionagentdata
type Wfmintradaydataupdatetopicintradayperformancepredictionagentdata struct { 
	// InteractingTimeSeconds
	InteractingTimeSeconds *float32 `json:"interactingTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayperformancepredictionagentdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
