package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Progresstransferevent
type Progresstransferevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// CommandId - The id (V4 UUID) used to identify the transfer already started by the external platform.
	CommandId *string `json:"commandId,omitempty"`


	// ObjectCommunicationId - The id (V4 UUID) of the communication that is being transferred.
	ObjectCommunicationId *string `json:"objectCommunicationId,omitempty"`


	// DestinationCommunicationId - The id (V4 UUID) of the communication that is being transferred to.
	DestinationCommunicationId *string `json:"destinationCommunicationId,omitempty"`

}

func (o *Progresstransferevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Progresstransferevent
	
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
		
		CommandId *string `json:"commandId,omitempty"`
		
		ObjectCommunicationId *string `json:"objectCommunicationId,omitempty"`
		
		DestinationCommunicationId *string `json:"destinationCommunicationId,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommandId: o.CommandId,
		
		ObjectCommunicationId: o.ObjectCommunicationId,
		
		DestinationCommunicationId: o.DestinationCommunicationId,
		Alias:    (*Alias)(o),
	})
}

func (o *Progresstransferevent) UnmarshalJSON(b []byte) error {
	var ProgresstransfereventMap map[string]interface{}
	err := json.Unmarshal(b, &ProgresstransfereventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := ProgresstransfereventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := ProgresstransfereventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := ProgresstransfereventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommandId, ok := ProgresstransfereventMap["commandId"].(string); ok {
		o.CommandId = &CommandId
	}
    
	if ObjectCommunicationId, ok := ProgresstransfereventMap["objectCommunicationId"].(string); ok {
		o.ObjectCommunicationId = &ObjectCommunicationId
	}
    
	if DestinationCommunicationId, ok := ProgresstransfereventMap["destinationCommunicationId"].(string); ok {
		o.DestinationCommunicationId = &DestinationCommunicationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Progresstransferevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
