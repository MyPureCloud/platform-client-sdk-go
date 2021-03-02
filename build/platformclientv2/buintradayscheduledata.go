package platformclientv2
import (
	"encoding/json"
)

// Buintradayscheduledata
type Buintradayscheduledata struct { 
	// OnQueueTimeSeconds - The total on-queue time in seconds for all agents in this group
	OnQueueTimeSeconds *int `json:"onQueueTimeSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buintradayscheduledata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
