package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
