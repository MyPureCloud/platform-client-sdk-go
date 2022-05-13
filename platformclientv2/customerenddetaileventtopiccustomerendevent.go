package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Customerenddetaileventtopiccustomerendevent
type Customerenddetaileventtopiccustomerendevent struct { 
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


	// ExternalOrganizationId
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`


	// ExternalContactId
	ExternalContactId *string `json:"externalContactId,omitempty"`


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


	// CallbackUserName
	CallbackUserName *string `json:"callbackUserName,omitempty"`


	// CallbackNumbers
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`


	// CallbackScheduledTime
	CallbackScheduledTime *int `json:"callbackScheduledTime,omitempty"`


	// Subject
	Subject *string `json:"subject,omitempty"`


	// MessageType
	MessageType *string `json:"messageType,omitempty"`


	// InteractingDurationMs
	InteractingDurationMs *int `json:"interactingDurationMs,omitempty"`

}

func (o *Customerenddetaileventtopiccustomerendevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Customerenddetaileventtopiccustomerendevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackScheduledTime *int `json:"callbackScheduledTime,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		InteractingDurationMs *int `json:"interactingDurationMs,omitempty"`
		*Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		SessionId: o.SessionId,
		
		DisconnectType: o.DisconnectType,
		
		MediaType: o.MediaType,
		
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		ExternalContactId: o.ExternalContactId,
		
		Provider: o.Provider,
		
		Direction: o.Direction,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		AddressTo: o.AddressTo,
		
		AddressFrom: o.AddressFrom,
		
		CallbackUserName: o.CallbackUserName,
		
		CallbackNumbers: o.CallbackNumbers,
		
		CallbackScheduledTime: o.CallbackScheduledTime,
		
		Subject: o.Subject,
		
		MessageType: o.MessageType,
		
		InteractingDurationMs: o.InteractingDurationMs,
		Alias:    (*Alias)(o),
	})
}

func (o *Customerenddetaileventtopiccustomerendevent) UnmarshalJSON(b []byte) error {
	var CustomerenddetaileventtopiccustomerendeventMap map[string]interface{}
	err := json.Unmarshal(b, &CustomerenddetaileventtopiccustomerendeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := CustomerenddetaileventtopiccustomerendeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := CustomerenddetaileventtopiccustomerendeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := CustomerenddetaileventtopiccustomerendeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SessionId, ok := CustomerenddetaileventtopiccustomerendeventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if DisconnectType, ok := CustomerenddetaileventtopiccustomerendeventMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if MediaType, ok := CustomerenddetaileventtopiccustomerendeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ExternalOrganizationId, ok := CustomerenddetaileventtopiccustomerendeventMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
    
	if ExternalContactId, ok := CustomerenddetaileventtopiccustomerendeventMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if Provider, ok := CustomerenddetaileventtopiccustomerendeventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Direction, ok := CustomerenddetaileventtopiccustomerendeventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := CustomerenddetaileventtopiccustomerendeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := CustomerenddetaileventtopiccustomerendeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if AddressTo, ok := CustomerenddetaileventtopiccustomerendeventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if AddressFrom, ok := CustomerenddetaileventtopiccustomerendeventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if CallbackUserName, ok := CustomerenddetaileventtopiccustomerendeventMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
    
	if CallbackNumbers, ok := CustomerenddetaileventtopiccustomerendeventMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackScheduledTime, ok := CustomerenddetaileventtopiccustomerendeventMap["callbackScheduledTime"].(float64); ok {
		CallbackScheduledTimeInt := int(CallbackScheduledTime)
		o.CallbackScheduledTime = &CallbackScheduledTimeInt
	}
	
	if Subject, ok := CustomerenddetaileventtopiccustomerendeventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if MessageType, ok := CustomerenddetaileventtopiccustomerendeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if InteractingDurationMs, ok := CustomerenddetaileventtopiccustomerendeventMap["interactingDurationMs"].(float64); ok {
		InteractingDurationMsInt := int(InteractingDurationMs)
		o.InteractingDurationMs = &InteractingDurationMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Customerenddetaileventtopiccustomerendevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
