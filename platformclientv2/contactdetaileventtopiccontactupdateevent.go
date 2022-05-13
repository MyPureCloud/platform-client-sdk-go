package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactdetaileventtopiccontactupdateevent
type Contactdetaileventtopiccontactupdateevent struct { 
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

}

func (o *Contactdetaileventtopiccontactupdateevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactdetaileventtopiccontactupdateevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
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
		*Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		SessionId: o.SessionId,
		
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
		Alias:    (*Alias)(o),
	})
}

func (o *Contactdetaileventtopiccontactupdateevent) UnmarshalJSON(b []byte) error {
	var ContactdetaileventtopiccontactupdateeventMap map[string]interface{}
	err := json.Unmarshal(b, &ContactdetaileventtopiccontactupdateeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := ContactdetaileventtopiccontactupdateeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := ContactdetaileventtopiccontactupdateeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := ContactdetaileventtopiccontactupdateeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SessionId, ok := ContactdetaileventtopiccontactupdateeventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if MediaType, ok := ContactdetaileventtopiccontactupdateeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ExternalOrganizationId, ok := ContactdetaileventtopiccontactupdateeventMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
    
	if ExternalContactId, ok := ContactdetaileventtopiccontactupdateeventMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if Provider, ok := ContactdetaileventtopiccontactupdateeventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Direction, ok := ContactdetaileventtopiccontactupdateeventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := ContactdetaileventtopiccontactupdateeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := ContactdetaileventtopiccontactupdateeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if AddressTo, ok := ContactdetaileventtopiccontactupdateeventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if AddressFrom, ok := ContactdetaileventtopiccontactupdateeventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if CallbackUserName, ok := ContactdetaileventtopiccontactupdateeventMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
    
	if CallbackNumbers, ok := ContactdetaileventtopiccontactupdateeventMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackScheduledTime, ok := ContactdetaileventtopiccontactupdateeventMap["callbackScheduledTime"].(float64); ok {
		CallbackScheduledTimeInt := int(CallbackScheduledTime)
		o.CallbackScheduledTime = &CallbackScheduledTimeInt
	}
	
	if Subject, ok := ContactdetaileventtopiccontactupdateeventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if MessageType, ok := ContactdetaileventtopiccontactupdateeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactdetaileventtopiccontactupdateevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
