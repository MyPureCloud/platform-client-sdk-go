package platformclientv2
import (
	"encoding/json"
)

// Wfmbuintradaydataupdatetopicbuintradayscheduledata
type Wfmbuintradaydataupdatetopicbuintradayscheduledata struct { 
	// OnQueueTimeSeconds
	OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
