package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowenddetaileventtopicflowendevent
type Flowenddetaileventtopicflowendevent struct { 
	// EventTime
	EventTime *int `json:"eventTime,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// ParticipantId
	ParticipantId *string `json:"participantId,omitempty"`


	// SessionId
	SessionId *string `json:"sessionId,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


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


	// ConnectedDurationMs
	ConnectedDurationMs *int `json:"connectedDurationMs,omitempty"`

}

func (o *Flowenddetaileventtopicflowendevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowenddetaileventtopicflowendevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
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
		
		ConnectedDurationMs *int `json:"connectedDurationMs,omitempty"`
		*Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		SessionId: o.SessionId,
		
		DisconnectType: o.DisconnectType,
		
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
		
		ConnectedDurationMs: o.ConnectedDurationMs,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowenddetaileventtopicflowendevent) UnmarshalJSON(b []byte) error {
	var FlowenddetaileventtopicflowendeventMap map[string]interface{}
	err := json.Unmarshal(b, &FlowenddetaileventtopicflowendeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := FlowenddetaileventtopicflowendeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := FlowenddetaileventtopicflowendeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := FlowenddetaileventtopicflowendeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SessionId, ok := FlowenddetaileventtopicflowendeventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if DisconnectType, ok := FlowenddetaileventtopicflowendeventMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if MediaType, ok := FlowenddetaileventtopicflowendeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Provider, ok := FlowenddetaileventtopicflowendeventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Direction, ok := FlowenddetaileventtopicflowendeventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := FlowenddetaileventtopicflowendeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := FlowenddetaileventtopicflowendeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if AddressTo, ok := FlowenddetaileventtopicflowendeventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if AddressFrom, ok := FlowenddetaileventtopicflowendeventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if Subject, ok := FlowenddetaileventtopicflowendeventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if MessageType, ok := FlowenddetaileventtopicflowendeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if FlowType, ok := FlowenddetaileventtopicflowendeventMap["flowType"].(string); ok {
		o.FlowType = &FlowType
	}
    
	if FlowId, ok := FlowenddetaileventtopicflowendeventMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
    
	if DivisionId, ok := FlowenddetaileventtopicflowendeventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if FlowVersion, ok := FlowenddetaileventtopicflowendeventMap["flowVersion"].(string); ok {
		o.FlowVersion = &FlowVersion
	}
    
	if ConnectedDurationMs, ok := FlowenddetaileventtopicflowendeventMap["connectedDurationMs"].(float64); ok {
		ConnectedDurationMsInt := int(ConnectedDurationMs)
		o.ConnectedDurationMs = &ConnectedDurationMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowenddetaileventtopicflowendevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
