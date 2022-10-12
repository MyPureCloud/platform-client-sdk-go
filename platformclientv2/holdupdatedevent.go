package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Holdupdatedevent
type Holdupdatedevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId - A unique Id (V4 UUID) identifying this communication
	CommunicationId *string `json:"communicationId,omitempty"`


	// Held - Indicates whether this communication is held.
	Held *bool `json:"held,omitempty"`

}

func (o *Holdupdatedevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Holdupdatedevent
	
	EventDateTime := new(string)
	if o.EventDateTime != nil {
		
		*EventDateTime = timeutil.Strftime(o.EventDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventDateTime = nil
	}
	
	return json.Marshal(&struct { 
		EventId *string `json:"eventId,omitempty"`
		
		EventDateTime *string `json:"eventDateTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		Held: o.Held,
		Alias:    (*Alias)(o),
	})
}

func (o *Holdupdatedevent) UnmarshalJSON(b []byte) error {
	var HoldupdatedeventMap map[string]interface{}
	err := json.Unmarshal(b, &HoldupdatedeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := HoldupdatedeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := HoldupdatedeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := HoldupdatedeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := HoldupdatedeventMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if Held, ok := HoldupdatedeventMap["held"].(bool); ok {
		o.Held = &Held
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Holdupdatedevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
