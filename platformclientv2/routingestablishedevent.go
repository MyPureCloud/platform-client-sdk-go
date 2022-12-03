package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingestablishedevent
type Routingestablishedevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId - A unique Id (V4 UUID) identifying this communication
	CommunicationId *string `json:"communicationId,omitempty"`


	// PhoneNumber - Identifies the phone number used to reach this queue if it is different from the information that would be accessed by queueId.
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// QueueId - The id (V4 UUID) of the queue that is routing this conversation.
	QueueId *string `json:"queueId,omitempty"`


	// Ani - The automatic number identification if it is available for this conversation.
	Ani *string `json:"ani,omitempty"`


	// Dnis - The dialed number identification if it is available for this conversation.
	Dnis *string `json:"dnis,omitempty"`


	// SkillIds - The unique identifiers (V4 UUID) for the skills that should be used to determine the destination for the conversation.
	SkillIds *[]string `json:"skillIds,omitempty"`


	// LanguageId - The unique identifier (V4 UUID) for the language that should be used to determine the destination for the conversation.
	LanguageId *string `json:"languageId,omitempty"`


	// InitialConfiguration - Metadata about this communication.
	InitialConfiguration *Initialconfiguration `json:"initialConfiguration,omitempty"`


	// SourceConfiguration - Metadata about the source of this communication's interaction.
	SourceConfiguration *Sourceconfiguration `json:"sourceConfiguration,omitempty"`

}

func (o *Routingestablishedevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingestablishedevent
	
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
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		InitialConfiguration *Initialconfiguration `json:"initialConfiguration,omitempty"`
		
		SourceConfiguration *Sourceconfiguration `json:"sourceConfiguration,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		PhoneNumber: o.PhoneNumber,
		
		QueueId: o.QueueId,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		SkillIds: o.SkillIds,
		
		LanguageId: o.LanguageId,
		
		InitialConfiguration: o.InitialConfiguration,
		
		SourceConfiguration: o.SourceConfiguration,
		Alias:    (*Alias)(o),
	})
}

func (o *Routingestablishedevent) UnmarshalJSON(b []byte) error {
	var RoutingestablishedeventMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingestablishedeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := RoutingestablishedeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := RoutingestablishedeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := RoutingestablishedeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := RoutingestablishedeventMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if PhoneNumber, ok := RoutingestablishedeventMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if QueueId, ok := RoutingestablishedeventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if Ani, ok := RoutingestablishedeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := RoutingestablishedeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if SkillIds, ok := RoutingestablishedeventMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if LanguageId, ok := RoutingestablishedeventMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if InitialConfiguration, ok := RoutingestablishedeventMap["initialConfiguration"].(map[string]interface{}); ok {
		InitialConfigurationString, _ := json.Marshal(InitialConfiguration)
		json.Unmarshal(InitialConfigurationString, &o.InitialConfiguration)
	}
	
	if SourceConfiguration, ok := RoutingestablishedeventMap["sourceConfiguration"].(map[string]interface{}); ok {
		SourceConfigurationString, _ := json.Marshal(SourceConfiguration)
		json.Unmarshal(SourceConfigurationString, &o.SourceConfiguration)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routingestablishedevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
