package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Communicationendedevent
type Communicationendedevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId - A unique Id (V4 UUID) identifying this communication
	CommunicationId *string `json:"communicationId,omitempty"`


	// DisconnectType - Indicates how this communication was ended.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// DestinationConversationId - The id (V4 UUID) of the conversation that the communication is being moved to when conversations are merged.
	DestinationConversationId *string `json:"destinationConversationId,omitempty"`

}

func (o *Communicationendedevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Communicationendedevent
	
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
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		DestinationConversationId *string `json:"destinationConversationId,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		DisconnectType: o.DisconnectType,
		
		DestinationConversationId: o.DestinationConversationId,
		Alias:    (*Alias)(o),
	})
}

func (o *Communicationendedevent) UnmarshalJSON(b []byte) error {
	var CommunicationendedeventMap map[string]interface{}
	err := json.Unmarshal(b, &CommunicationendedeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := CommunicationendedeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := CommunicationendedeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := CommunicationendedeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := CommunicationendedeventMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if DisconnectType, ok := CommunicationendedeventMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if DestinationConversationId, ok := CommunicationendedeventMap["destinationConversationId"].(string); ok {
		o.DestinationConversationId = &DestinationConversationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Communicationendedevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
