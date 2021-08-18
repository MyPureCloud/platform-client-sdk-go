package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Integrationevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Integrationevent

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		Level *string `json:"level,omitempty"`
		
		EventCode *string `json:"eventCode,omitempty"`
		
		Message *Messageinfo `json:"message,omitempty"`
		
		Entities *[]Evententity `json:"entities,omitempty"`
		
		ContextAttributes *map[string]string `json:"contextAttributes,omitempty"`
		
		DetailMessage *Messageinfo `json:"detailMessage,omitempty"`
		
		User *User `json:"user,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		SelfUri: u.SelfUri,
		
		CorrelationId: u.CorrelationId,
		
		Timestamp: Timestamp,
		
		Level: u.Level,
		
		EventCode: u.EventCode,
		
		Message: u.Message,
		
		Entities: u.Entities,
		
		ContextAttributes: u.ContextAttributes,
		
		DetailMessage: u.DetailMessage,
		
		User: u.User,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Integrationevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
