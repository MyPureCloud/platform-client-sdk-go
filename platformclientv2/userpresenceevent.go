package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userpresenceevent
type Userpresenceevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// UserId - The User ID of the user associated with this UserPresence
	UserId *string `json:"userId,omitempty"`


	// SourceId - The id (V4 UUID) of the presence source being updated
	SourceId *string `json:"sourceId,omitempty"`


	// PresenceDefinitionId - The id (UUID) of the presence definition that the user presence is associated with
	PresenceDefinitionId *string `json:"presenceDefinitionId,omitempty"`


	// Message - The message associated with the presence
	Message *string `json:"message,omitempty"`

}

func (o *Userpresenceevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userpresenceevent
	
	EventDateTime := new(string)
	if o.EventDateTime != nil {
		
		*EventDateTime = timeutil.Strftime(o.EventDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventDateTime = nil
	}
	
	return json.Marshal(&struct { 
		EventId *string `json:"eventId,omitempty"`
		
		EventDateTime *string `json:"eventDateTime,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		SourceId *string `json:"sourceId,omitempty"`
		
		PresenceDefinitionId *string `json:"presenceDefinitionId,omitempty"`
		
		Message *string `json:"message,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		UserId: o.UserId,
		
		SourceId: o.SourceId,
		
		PresenceDefinitionId: o.PresenceDefinitionId,
		
		Message: o.Message,
		Alias:    (*Alias)(o),
	})
}

func (o *Userpresenceevent) UnmarshalJSON(b []byte) error {
	var UserpresenceeventMap map[string]interface{}
	err := json.Unmarshal(b, &UserpresenceeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := UserpresenceeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := UserpresenceeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if UserId, ok := UserpresenceeventMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if SourceId, ok := UserpresenceeventMap["sourceId"].(string); ok {
		o.SourceId = &SourceId
	}
    
	if PresenceDefinitionId, ok := UserpresenceeventMap["presenceDefinitionId"].(string); ok {
		o.PresenceDefinitionId = &PresenceDefinitionId
	}
    
	if Message, ok := UserpresenceeventMap["message"].(string); ok {
		o.Message = &Message
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userpresenceevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
