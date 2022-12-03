package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ivrestablishedevent
type Ivrestablishedevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId - A unique Id (V4 UUID) identifying this communication
	CommunicationId *string `json:"communicationId,omitempty"`


	// IvrPhoneNumber - The phone number for this IVR, if any is known
	IvrPhoneNumber *string `json:"ivrPhoneNumber,omitempty"`


	// IvrName - A displayable name for this IVR, if any is known.
	IvrName *string `json:"ivrName,omitempty"`


	// Ani - The automatic number identification if it is available for this conversation.
	Ani *string `json:"ani,omitempty"`


	// Dnis - The dialed number identification if it is available for this conversation.
	Dnis *string `json:"dnis,omitempty"`


	// InitialConfiguration - Metadata about this communication.
	InitialConfiguration *Initialconfiguration `json:"initialConfiguration,omitempty"`


	// SourceConfiguration - Metadata about the source of this communication's interaction.
	SourceConfiguration *Sourceconfiguration `json:"sourceConfiguration,omitempty"`

}

func (o *Ivrestablishedevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ivrestablishedevent
	
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
		
		IvrPhoneNumber *string `json:"ivrPhoneNumber,omitempty"`
		
		IvrName *string `json:"ivrName,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		InitialConfiguration *Initialconfiguration `json:"initialConfiguration,omitempty"`
		
		SourceConfiguration *Sourceconfiguration `json:"sourceConfiguration,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		IvrPhoneNumber: o.IvrPhoneNumber,
		
		IvrName: o.IvrName,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		InitialConfiguration: o.InitialConfiguration,
		
		SourceConfiguration: o.SourceConfiguration,
		Alias:    (*Alias)(o),
	})
}

func (o *Ivrestablishedevent) UnmarshalJSON(b []byte) error {
	var IvrestablishedeventMap map[string]interface{}
	err := json.Unmarshal(b, &IvrestablishedeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := IvrestablishedeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := IvrestablishedeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := IvrestablishedeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := IvrestablishedeventMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if IvrPhoneNumber, ok := IvrestablishedeventMap["ivrPhoneNumber"].(string); ok {
		o.IvrPhoneNumber = &IvrPhoneNumber
	}
    
	if IvrName, ok := IvrestablishedeventMap["ivrName"].(string); ok {
		o.IvrName = &IvrName
	}
    
	if Ani, ok := IvrestablishedeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := IvrestablishedeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if InitialConfiguration, ok := IvrestablishedeventMap["initialConfiguration"].(map[string]interface{}); ok {
		InitialConfigurationString, _ := json.Marshal(InitialConfiguration)
		json.Unmarshal(InitialConfigurationString, &o.InitialConfiguration)
	}
	
	if SourceConfiguration, ok := IvrestablishedeventMap["sourceConfiguration"].(map[string]interface{}); ok {
		SourceConfigurationString, _ := json.Marshal(SourceConfiguration)
		json.Unmarshal(SourceConfigurationString, &o.SourceConfiguration)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ivrestablishedevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
