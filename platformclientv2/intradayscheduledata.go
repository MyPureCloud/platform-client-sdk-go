package platformclientv2
import (
	"encoding/json"
)

// Intradayscheduledata
type Intradayscheduledata struct { 
	// OnQueueTimeSeconds - The total scheduled on-queue time in seconds for all agents in this group
	OnQueueTimeSeconds *int32 `json:"onQueueTimeSeconds,omitempty"`


	// ScheduledTimeSeconds - The total scheduled time in seconds for all agents in this group
	ScheduledTimeSeconds *int32 `json:"scheduledTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
