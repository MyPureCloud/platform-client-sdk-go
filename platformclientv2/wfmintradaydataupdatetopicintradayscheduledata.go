package platformclientv2
import (
	"encoding/json"
)

// Wfmintradaydataupdatetopicintradayscheduledata
type Wfmintradaydataupdatetopicintradayscheduledata struct { 
	// OnQueueTimeSeconds
	OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`


	// ScheduledTimeSeconds
	ScheduledTimeSeconds *int `json:"scheduledTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
