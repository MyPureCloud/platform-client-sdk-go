package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowstartdetaileventtopicflowstartevent
type Flowstartdetaileventtopicflowstartevent struct { 
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


	// AddressTo
	AddressTo *string `json:"addressTo,omitempty"`


	// AddressFrom
	AddressFrom *string `json:"addressFrom,omitempty"`


	// Subject
	Subject *string `json:"subject,omitempty"`


	// MessageType
	MessageType *string `json:"messageType,omitempty"`


	// FlowType
	FlowType *string `json:"flowType,omitempty"`


	// FlowId
	FlowId *string `json:"flowId,omitempty"`


	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`


	// FlowVersion
	FlowVersion *string `json:"flowVersion,omitempty"`

}

func (o *Flowstartdetaileventtopicflowstartevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowstartdetaileventtopicflowstartevent
	
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
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		FlowType *string `json:"flowType,omitempty"`
		
		FlowId *string `json:"flowId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		FlowVersion *string `json:"flowVersion,omitempty"`
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
		
		AddressTo: o.AddressTo,
		
		AddressFrom: o.AddressFrom,
		
		Subject: o.Subject,
		
		MessageType: o.MessageType,
		
		FlowType: o.FlowType,
		
		FlowId: o.FlowId,
		
		DivisionId: o.DivisionId,
		
		FlowVersion: o.FlowVersion,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowstartdetaileventtopicflowstartevent) UnmarshalJSON(b []byte) error {
	var FlowstartdetaileventtopicflowstarteventMap map[string]interface{}
	err := json.Unmarshal(b, &FlowstartdetaileventtopicflowstarteventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := FlowstartdetaileventtopicflowstarteventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := FlowstartdetaileventtopicflowstarteventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := FlowstartdetaileventtopicflowstarteventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SessionId, ok := FlowstartdetaileventtopicflowstarteventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if MediaType, ok := FlowstartdetaileventtopicflowstarteventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Provider, ok := FlowstartdetaileventtopicflowstarteventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Direction, ok := FlowstartdetaileventtopicflowstarteventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := FlowstartdetaileventtopicflowstarteventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := FlowstartdetaileventtopicflowstarteventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if AddressTo, ok := FlowstartdetaileventtopicflowstarteventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if AddressFrom, ok := FlowstartdetaileventtopicflowstarteventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if Subject, ok := FlowstartdetaileventtopicflowstarteventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if MessageType, ok := FlowstartdetaileventtopicflowstarteventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if FlowType, ok := FlowstartdetaileventtopicflowstarteventMap["flowType"].(string); ok {
		o.FlowType = &FlowType
	}
    
	if FlowId, ok := FlowstartdetaileventtopicflowstarteventMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
    
	if DivisionId, ok := FlowstartdetaileventtopicflowstarteventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if FlowVersion, ok := FlowstartdetaileventtopicflowstarteventMap["flowVersion"].(string); ok {
		o.FlowVersion = &FlowVersion
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flowstartdetaileventtopicflowstartevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
