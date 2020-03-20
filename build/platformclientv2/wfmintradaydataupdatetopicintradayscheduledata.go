package platformclientv2
import (
	"encoding/json"
)

// Wfmintradaydataupdatetopicintradayscheduledata
type Wfmintradaydataupdatetopicintradayscheduledata struct { 
	// OnQueueTimeSeconds
	OnQueueTimeSeconds *int32 `json:"onQueueTimeSeconds,omitempty"`


	// ScheduledTimeSeconds
	ScheduledTimeSeconds *int32 `json:"scheduledTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
