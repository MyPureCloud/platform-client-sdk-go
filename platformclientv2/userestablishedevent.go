package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userestablishedevent
type Userestablishedevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// CommunicationId - A unique Id (V4 UUID) identifying this communication
	CommunicationId *string `json:"communicationId,omitempty"`


	// PhoneNumber - Identifies the phone number used to reach this user if it is different from the information that would be accessed by userId.
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// UserId - The userId (V4 UUID) for the user this communication belongs to.
	UserId *string `json:"userId,omitempty"`


	// StationId - A Station ID (V4 UUID) that identifies the station being used if the user is using a station and the stationId is known.
	StationId *string `json:"stationId,omitempty"`


	// Ani - The automatic number identification if it is available for this conversation.
	Ani *string `json:"ani,omitempty"`


	// Dnis - The dialed number identification if it is available for this conversation.
	Dnis *string `json:"dnis,omitempty"`


	// AfterCallWorkRequired - Indicates whether or not this user will be required to complete after call work.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// QueueId - The id (V4 UUID) of the queue that the user is calling on behalf of. Applies to outbound calls only.
	QueueId *string `json:"queueId,omitempty"`


	// InitialConfiguration - Metadata about this communication.
	InitialConfiguration *Initialconfiguration `json:"initialConfiguration,omitempty"`


	// SourceConfiguration - Metadata about the source of this communication's interaction.
	SourceConfiguration *Sourceconfiguration `json:"sourceConfiguration,omitempty"`

}

func (o *Userestablishedevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userestablishedevent
	
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
		
		UserId *string `json:"userId,omitempty"`
		
		StationId *string `json:"stationId,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		InitialConfiguration *Initialconfiguration `json:"initialConfiguration,omitempty"`
		
		SourceConfiguration *Sourceconfiguration `json:"sourceConfiguration,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		PhoneNumber: o.PhoneNumber,
		
		UserId: o.UserId,
		
		StationId: o.StationId,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		QueueId: o.QueueId,
		
		InitialConfiguration: o.InitialConfiguration,
		
		SourceConfiguration: o.SourceConfiguration,
		Alias:    (*Alias)(o),
	})
}

func (o *Userestablishedevent) UnmarshalJSON(b []byte) error {
	var UserestablishedeventMap map[string]interface{}
	err := json.Unmarshal(b, &UserestablishedeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := UserestablishedeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := UserestablishedeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := UserestablishedeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := UserestablishedeventMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if PhoneNumber, ok := UserestablishedeventMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if UserId, ok := UserestablishedeventMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if StationId, ok := UserestablishedeventMap["stationId"].(string); ok {
		o.StationId = &StationId
	}
    
	if Ani, ok := UserestablishedeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := UserestablishedeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if AfterCallWorkRequired, ok := UserestablishedeventMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    
	if QueueId, ok := UserestablishedeventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if InitialConfiguration, ok := UserestablishedeventMap["initialConfiguration"].(map[string]interface{}); ok {
		InitialConfigurationString, _ := json.Marshal(InitialConfiguration)
		json.Unmarshal(InitialConfigurationString, &o.InitialConfiguration)
	}
	
	if SourceConfiguration, ok := UserestablishedeventMap["sourceConfiguration"].(map[string]interface{}); ok {
		SourceConfigurationString, _ := json.Marshal(SourceConfiguration)
		json.Unmarshal(SourceConfigurationString, &o.SourceConfiguration)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userestablishedevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
