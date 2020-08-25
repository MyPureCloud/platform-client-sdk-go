package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationvideoeventtopicaftercallwork
type Queueconversationvideoeventtopicaftercallwork struct { 
	// State
	State *string `json:"state,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime
	EndTime *time.Time `json:"endTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicaftercallwork) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
