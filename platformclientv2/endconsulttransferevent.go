package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Endconsulttransferevent
type Endconsulttransferevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`

}

func (o *Endconsulttransferevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Endconsulttransferevent
	
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
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		Alias:    (*Alias)(o),
	})
}

func (o *Endconsulttransferevent) UnmarshalJSON(b []byte) error {
	var EndconsulttransfereventMap map[string]interface{}
	err := json.Unmarshal(b, &EndconsulttransfereventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := EndconsulttransfereventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := EndconsulttransfereventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := EndconsulttransfereventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Endconsulttransferevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
