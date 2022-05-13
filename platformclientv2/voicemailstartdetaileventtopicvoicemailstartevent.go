package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailstartdetaileventtopicvoicemailstartevent
type Voicemailstartdetaileventtopicvoicemailstartevent struct { 
	// EventTime
	EventTime *int `json:"eventTime,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// ParticipantId
	ParticipantId *string `json:"participantId,omitempty"`


	// SessionId
	SessionId *string `json:"sessionId,omitempty"`


	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Ani
	Ani *string `json:"ani,omitempty"`


	// Dnis
	Dnis *string `json:"dnis,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// QueueId
	QueueId *string `json:"queueId,omitempty"`


	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`

}

func (o *Voicemailstartdetaileventtopicvoicemailstartevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailstartdetaileventtopicvoicemailstartevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		*Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		SessionId: o.SessionId,
		
		MediaType: o.MediaType,
		
		Provider: o.Provider,
		
		Direction: o.Direction,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		UserId: o.UserId,
		
		QueueId: o.QueueId,
		
		DivisionId: o.DivisionId,
		Alias:    (*Alias)(o),
	})
}

func (o *Voicemailstartdetaileventtopicvoicemailstartevent) UnmarshalJSON(b []byte) error {
	var VoicemailstartdetaileventtopicvoicemailstarteventMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailstartdetaileventtopicvoicemailstarteventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SessionId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if MediaType, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Provider, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Direction, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if UserId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if QueueId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if DivisionId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailstartdetaileventtopicvoicemailstartevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
