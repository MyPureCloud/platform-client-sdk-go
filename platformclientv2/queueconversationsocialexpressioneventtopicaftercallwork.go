package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationsocialexpressioneventtopicaftercallwork
type Queueconversationsocialexpressioneventtopicaftercallwork struct { 
	// State
	State *string `json:"state,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime
	EndTime *time.Time `json:"endTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicaftercallwork) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
