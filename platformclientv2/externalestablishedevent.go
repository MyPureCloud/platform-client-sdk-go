package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalestablishedevent
type Externalestablishedevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId - A unique Id (V4 UUID) identifying this communication
	CommunicationId *string `json:"communicationId,omitempty"`


	// Ani - The automatic number identification if it is available for this conversation.
	Ani *string `json:"ani,omitempty"`


	// AniName - The automatic number identification name if it is available for this conversation.
	AniName *string `json:"aniName,omitempty"`


	// Dnis - The dialed number identification if it is available for this conversation.
	Dnis *string `json:"dnis,omitempty"`


	// DnisName - The dialed number identification name if it is available for this conversation.
	DnisName *string `json:"dnisName,omitempty"`


	// InitialConfiguration - Metadata about this communication.
	InitialConfiguration *Initialconfiguration `json:"initialConfiguration,omitempty"`


	// SourceConfiguration - Metadata about the source of this communication's interaction.
	SourceConfiguration *Sourceconfiguration `json:"sourceConfiguration,omitempty"`

}

func (o *Externalestablishedevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalestablishedevent
	
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
		
		Ani *string `json:"ani,omitempty"`
		
		AniName *string `json:"aniName,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		DnisName *string `json:"dnisName,omitempty"`
		
		InitialConfiguration *Initialconfiguration `json:"initialConfiguration,omitempty"`
		
		SourceConfiguration *Sourceconfiguration `json:"sourceConfiguration,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		Ani: o.Ani,
		
		AniName: o.AniName,
		
		Dnis: o.Dnis,
		
		DnisName: o.DnisName,
		
		InitialConfiguration: o.InitialConfiguration,
		
		SourceConfiguration: o.SourceConfiguration,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalestablishedevent) UnmarshalJSON(b []byte) error {
	var ExternalestablishedeventMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalestablishedeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := ExternalestablishedeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := ExternalestablishedeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := ExternalestablishedeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := ExternalestablishedeventMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if Ani, ok := ExternalestablishedeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if AniName, ok := ExternalestablishedeventMap["aniName"].(string); ok {
		o.AniName = &AniName
	}
    
	if Dnis, ok := ExternalestablishedeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if DnisName, ok := ExternalestablishedeventMap["dnisName"].(string); ok {
		o.DnisName = &DnisName
	}
    
	if InitialConfiguration, ok := ExternalestablishedeventMap["initialConfiguration"].(map[string]interface{}); ok {
		InitialConfigurationString, _ := json.Marshal(InitialConfiguration)
		json.Unmarshal(InitialConfigurationString, &o.InitialConfiguration)
	}
	
	if SourceConfiguration, ok := ExternalestablishedeventMap["sourceConfiguration"].(map[string]interface{}); ok {
		SourceConfigurationString, _ := json.Marshal(SourceConfiguration)
		json.Unmarshal(SourceConfigurationString, &o.SourceConfiguration)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalestablishedevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
