package platformclientv2
import (
	"time"
	"encoding/json"
)

// Agentactivitychangedtopicagentactivity
type Agentactivitychangedtopicagentactivity struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// RoutingStatus
	RoutingStatus *Agentactivitychangedtopicroutingstatus `json:"routingStatus,omitempty"`


	// Presence
	Presence *Agentactivitychangedtopicpresence `json:"presence,omitempty"`


	// OutOfOffice
	OutOfOffice *Agentactivitychangedtopicoutofoffice `json:"outOfOffice,omitempty"`


	// ActiveQueueIds
	ActiveQueueIds *[]string `json:"activeQueueIds,omitempty"`


	// DateActiveQueuesChanged
	DateActiveQueuesChanged *time.Time `json:"dateActiveQueuesChanged,omitempty"`

}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicagentactivity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
