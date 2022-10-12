package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Audioupdatedevent
type Audioupdatedevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId - A unique Id (V4 UUID) identifying this communication
	CommunicationId *string `json:"communicationId,omitempty"`


	// AudioState - The updated audioState for the target communication.
	AudioState *Audiostate `json:"audioState,omitempty"`

}

func (o *Audioupdatedevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Audioupdatedevent
	
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
		
		AudioState *Audiostate `json:"audioState,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		AudioState: o.AudioState,
		Alias:    (*Alias)(o),
	})
}

func (o *Audioupdatedevent) UnmarshalJSON(b []byte) error {
	var AudioupdatedeventMap map[string]interface{}
	err := json.Unmarshal(b, &AudioupdatedeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := AudioupdatedeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := AudioupdatedeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := AudioupdatedeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := AudioupdatedeventMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if AudioState, ok := AudioupdatedeventMap["audioState"].(map[string]interface{}); ok {
		AudioStateString, _ := json.Marshal(AudioState)
		json.Unmarshal(AudioStateString, &o.AudioState)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Audioupdatedevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
