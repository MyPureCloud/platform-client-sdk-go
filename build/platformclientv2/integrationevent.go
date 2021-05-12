package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationevent - Describes an event that has happened related to an integration
type Integrationevent struct { 
	// Id - Unique ID for this event
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CorrelationId - Correlation ID for the event
	CorrelationId *string `json:"correlationId,omitempty"`


	// Timestamp - Time the event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// Level - Indicates the severity of the event.
	Level *string `json:"level,omitempty"`


	// EventCode - A classification for the event. Suitable for programmatic searching, sorting, or filtering
	EventCode *string `json:"eventCode,omitempty"`


	// Message - Message indicating what happened
	Message *Messageinfo `json:"message,omitempty"`


	// Entities - Collection of entities affected by or pertaining to the event (e.g. a list of Integrations or Bridge connectors)
	Entities *[]Evententity `json:"entities,omitempty"`


	// ContextAttributes - Map of context attributes specific to this event.
	ContextAttributes *map[string]string `json:"contextAttributes,omitempty"`


	// DetailMessage - Message with additional details about the event. (e.g. an exception cause.)
	DetailMessage *Messageinfo `json:"detailMessage,omitempty"`


	// User - User that took an action that resulted in the event.
	User *User `json:"user,omitempty"`

}

// String returns a JSON representation of the model
func (o *Integrationevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
