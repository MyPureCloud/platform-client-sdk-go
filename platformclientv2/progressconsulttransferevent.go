package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Progressconsulttransferevent
type Progressconsulttransferevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// InitiatingCommunicationId - The id (V4 UUID) of the communication representing the participant that is initiating the transfer.
	InitiatingCommunicationId *string `json:"initiatingCommunicationId,omitempty"`


	// DestinationCommunicationId - The id (V4 UUID) of the communication that is being transferred to.
	DestinationCommunicationId *string `json:"destinationCommunicationId,omitempty"`


	// ObjectCommunicationId - The id (V4 UUID) of the communication that is being transferred.
	ObjectCommunicationId *string `json:"objectCommunicationId,omitempty"`

}

func (o *Progressconsulttransferevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Progressconsulttransferevent
	
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
		
		InitiatingCommunicationId *string `json:"initiatingCommunicationId,omitempty"`
		
		DestinationCommunicationId *string `json:"destinationCommunicationId,omitempty"`
		
		ObjectCommunicationId *string `json:"objectCommunicationId,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		InitiatingCommunicationId: o.InitiatingCommunicationId,
		
		DestinationCommunicationId: o.DestinationCommunicationId,
		
		ObjectCommunicationId: o.ObjectCommunicationId,
		Alias:    (*Alias)(o),
	})
}

func (o *Progressconsulttransferevent) UnmarshalJSON(b []byte) error {
	var ProgressconsulttransfereventMap map[string]interface{}
	err := json.Unmarshal(b, &ProgressconsulttransfereventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := ProgressconsulttransfereventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := ProgressconsulttransfereventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := ProgressconsulttransfereventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if InitiatingCommunicationId, ok := ProgressconsulttransfereventMap["initiatingCommunicationId"].(string); ok {
		o.InitiatingCommunicationId = &InitiatingCommunicationId
	}
    
	if DestinationCommunicationId, ok := ProgressconsulttransfereventMap["destinationCommunicationId"].(string); ok {
		o.DestinationCommunicationId = &DestinationCommunicationId
	}
    
	if ObjectCommunicationId, ok := ProgressconsulttransfereventMap["objectCommunicationId"].(string); ok {
		o.ObjectCommunicationId = &ObjectCommunicationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Progressconsulttransferevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
