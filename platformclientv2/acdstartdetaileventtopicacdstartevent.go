package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Acdstartdetaileventtopicacdstartevent
type Acdstartdetaileventtopicacdstartevent struct { 
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


	// QueueId
	QueueId *string `json:"queueId,omitempty"`


	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`

}

func (o *Acdstartdetaileventtopicacdstartevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Acdstartdetaileventtopicacdstartevent
	
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
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackScheduledTime *int `json:"callbackScheduledTime,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
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
		
		AddressTo: o.AddressTo,
		
		AddressFrom: o.AddressFrom,
		
		CallbackUserName: o.CallbackUserName,
		
		CallbackNumbers: o.CallbackNumbers,
		
		CallbackScheduledTime: o.CallbackScheduledTime,
		
		Subject: o.Subject,
		
		MessageType: o.MessageType,
		
		QueueId: o.QueueId,
		
		DivisionId: o.DivisionId,
		Alias:    (*Alias)(o),
	})
}

func (o *Acdstartdetaileventtopicacdstartevent) UnmarshalJSON(b []byte) error {
	var AcdstartdetaileventtopicacdstarteventMap map[string]interface{}
	err := json.Unmarshal(b, &AcdstartdetaileventtopicacdstarteventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := AcdstartdetaileventtopicacdstarteventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := AcdstartdetaileventtopicacdstarteventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if ParticipantId, ok := AcdstartdetaileventtopicacdstarteventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
	
	if SessionId, ok := AcdstartdetaileventtopicacdstarteventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if MediaType, ok := AcdstartdetaileventtopicacdstarteventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if Provider, ok := AcdstartdetaileventtopicacdstarteventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if Direction, ok := AcdstartdetaileventtopicacdstarteventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if Ani, ok := AcdstartdetaileventtopicacdstarteventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
	
	if Dnis, ok := AcdstartdetaileventtopicacdstarteventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
	
	if AddressTo, ok := AcdstartdetaileventtopicacdstarteventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
	
	if AddressFrom, ok := AcdstartdetaileventtopicacdstarteventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
	
	if CallbackUserName, ok := AcdstartdetaileventtopicacdstarteventMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
	
	if CallbackNumbers, ok := AcdstartdetaileventtopicacdstarteventMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackScheduledTime, ok := AcdstartdetaileventtopicacdstarteventMap["callbackScheduledTime"].(float64); ok {
		CallbackScheduledTimeInt := int(CallbackScheduledTime)
		o.CallbackScheduledTime = &CallbackScheduledTimeInt
	}
	
	if Subject, ok := AcdstartdetaileventtopicacdstarteventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
	
	if MessageType, ok := AcdstartdetaileventtopicacdstarteventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
	
	if QueueId, ok := AcdstartdetaileventtopicacdstarteventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
	
	if DivisionId, ok := AcdstartdetaileventtopicacdstarteventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Acdstartdetaileventtopicacdstartevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
