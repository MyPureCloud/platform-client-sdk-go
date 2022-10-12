package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Endtransferevent
type Endtransferevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// CommandId - The id (V4 UUID) used to identify the transfer already started by the external platform.
	CommandId *string `json:"commandId,omitempty"`


	// FinalState - Indicates whether the transfer completed successfully, was cancelled, or failed for some reason.
	FinalState *string `json:"finalState,omitempty"`


	// ObjectCommunicationId - The id (V4 UUID) of the communication that was being transferred.
	ObjectCommunicationId *string `json:"objectCommunicationId,omitempty"`

}

func (o *Endtransferevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Endtransferevent
	
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
		
		FinalState *string `json:"finalState,omitempty"`
		
		ObjectCommunicationId *string `json:"objectCommunicationId,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommandId: o.CommandId,
		
		FinalState: o.FinalState,
		
		ObjectCommunicationId: o.ObjectCommunicationId,
		Alias:    (*Alias)(o),
	})
}

func (o *Endtransferevent) UnmarshalJSON(b []byte) error {
	var EndtransfereventMap map[string]interface{}
	err := json.Unmarshal(b, &EndtransfereventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := EndtransfereventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := EndtransfereventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := EndtransfereventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommandId, ok := EndtransfereventMap["commandId"].(string); ok {
		o.CommandId = &CommandId
	}
    
	if FinalState, ok := EndtransfereventMap["finalState"].(string); ok {
		o.FinalState = &FinalState
	}
    
	if ObjectCommunicationId, ok := EndtransfereventMap["objectCommunicationId"].(string); ok {
		o.ObjectCommunicationId = &ObjectCommunicationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Endtransferevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
