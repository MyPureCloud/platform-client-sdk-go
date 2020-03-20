package platformclientv2
import (
	"time"
	"encoding/json"
)

// Agentactivitychangedtopicroutingstatus
type Agentactivitychangedtopicroutingstatus struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicroutingstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
