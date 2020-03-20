package platformclientv2
import (
	"time"
	"encoding/json"
)

// Routingstatus
type Routingstatus struct { 
	// UserId - The userId of the agent
	UserId *string `json:"userId,omitempty"`


	// Status - Indicates the Routing State of the agent.  A value of OFF_QUEUE will be returned if the specified user does not exist.
	Status *string `json:"status,omitempty"`


	// StartTime - The timestamp when the agent went into this state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartTime *time.Time `json:"startTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routingstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
