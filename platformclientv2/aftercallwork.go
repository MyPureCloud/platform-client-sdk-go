package platformclientv2
import (
	"time"
	"encoding/json"
)

// Aftercallwork
type Aftercallwork struct { 
	// StartTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndTime *time.Time `json:"endTime,omitempty"`


	// State
	State *string `json:"state,omitempty"`

}

// String returns a JSON representation of the model
func (o *Aftercallwork) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
