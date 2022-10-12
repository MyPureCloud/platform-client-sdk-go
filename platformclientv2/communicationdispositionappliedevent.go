package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Communicationdispositionappliedevent
type Communicationdispositionappliedevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId - A unique Id (V4 UUID) identifying this communication
	CommunicationId *string `json:"communicationId,omitempty"`


	// Code - The wrapup-code (V4 UUID) used to disposition this interaction. If this value is not provided the disposition is considered skipped.
	Code *string `json:"code,omitempty"`


	// Notes - Text entered by the agent to describe the interaction or disposition. Ignored if the disposition is considered skipped.
	Notes *string `json:"notes,omitempty"`


	// Tags - The list of tags selected by the agent to describe the interaction or disposition. Ignored if the disposition is considered skipped.
	Tags *[]string `json:"tags,omitempty"`

}

func (o *Communicationdispositionappliedevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Communicationdispositionappliedevent
	
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
		
		Code *string `json:"code,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		Code: o.Code,
		
		Notes: o.Notes,
		
		Tags: o.Tags,
		Alias:    (*Alias)(o),
	})
}

func (o *Communicationdispositionappliedevent) UnmarshalJSON(b []byte) error {
	var CommunicationdispositionappliedeventMap map[string]interface{}
	err := json.Unmarshal(b, &CommunicationdispositionappliedeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := CommunicationdispositionappliedeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := CommunicationdispositionappliedeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := CommunicationdispositionappliedeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := CommunicationdispositionappliedeventMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if Code, ok := CommunicationdispositionappliedeventMap["code"].(string); ok {
		o.Code = &Code
	}
    
	if Notes, ok := CommunicationdispositionappliedeventMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if Tags, ok := CommunicationdispositionappliedeventMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Communicationdispositionappliedevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
