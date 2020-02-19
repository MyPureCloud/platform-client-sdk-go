package platformclientv2
import (
	"time"
	"encoding/json"
)

// Agentactivitychangedtopicpresence
type Agentactivitychangedtopicpresence struct { 
	// PresenceDefinition
	PresenceDefinition *Agentactivitychangedtopicorganizationpresence `json:"presenceDefinition,omitempty"`


	// PresenceMessage
	PresenceMessage *string `json:"presenceMessage,omitempty"`


	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicpresence) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
