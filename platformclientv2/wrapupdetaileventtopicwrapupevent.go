package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wrapupdetaileventtopicwrapupevent
type Wrapupdetaileventtopicwrapupevent struct { 
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


	// WrapupCode
	WrapupCode *string `json:"wrapupCode,omitempty"`


	// WrapupNotes
	WrapupNotes *string `json:"wrapupNotes,omitempty"`


	// WrapupDurationMs
	WrapupDurationMs *int `json:"wrapupDurationMs,omitempty"`

}

func (o *Wrapupdetaileventtopicwrapupevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wrapupdetaileventtopicwrapupevent
	
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
		
		WrapupCode *string `json:"wrapupCode,omitempty"`
		
		WrapupNotes *string `json:"wrapupNotes,omitempty"`
		
		WrapupDurationMs *int `json:"wrapupDurationMs,omitempty"`
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
		
		WrapupCode: o.WrapupCode,
		
		WrapupNotes: o.WrapupNotes,
		
		WrapupDurationMs: o.WrapupDurationMs,
		Alias:    (*Alias)(o),
	})
}

func (o *Wrapupdetaileventtopicwrapupevent) UnmarshalJSON(b []byte) error {
	var WrapupdetaileventtopicwrapupeventMap map[string]interface{}
	err := json.Unmarshal(b, &WrapupdetaileventtopicwrapupeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := WrapupdetaileventtopicwrapupeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := WrapupdetaileventtopicwrapupeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if ParticipantId, ok := WrapupdetaileventtopicwrapupeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
	
	if SessionId, ok := WrapupdetaileventtopicwrapupeventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if MediaType, ok := WrapupdetaileventtopicwrapupeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if Provider, ok := WrapupdetaileventtopicwrapupeventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if Direction, ok := WrapupdetaileventtopicwrapupeventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if Ani, ok := WrapupdetaileventtopicwrapupeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
	
	if Dnis, ok := WrapupdetaileventtopicwrapupeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
	
	if AddressTo, ok := WrapupdetaileventtopicwrapupeventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
	
	if AddressFrom, ok := WrapupdetaileventtopicwrapupeventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
	
	if CallbackUserName, ok := WrapupdetaileventtopicwrapupeventMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
	
	if CallbackNumbers, ok := WrapupdetaileventtopicwrapupeventMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackScheduledTime, ok := WrapupdetaileventtopicwrapupeventMap["callbackScheduledTime"].(float64); ok {
		CallbackScheduledTimeInt := int(CallbackScheduledTime)
		o.CallbackScheduledTime = &CallbackScheduledTimeInt
	}
	
	if Subject, ok := WrapupdetaileventtopicwrapupeventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
	
	if MessageType, ok := WrapupdetaileventtopicwrapupeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
	
	if QueueId, ok := WrapupdetaileventtopicwrapupeventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
	
	if WrapupCode, ok := WrapupdetaileventtopicwrapupeventMap["wrapupCode"].(string); ok {
		o.WrapupCode = &WrapupCode
	}
	
	if WrapupNotes, ok := WrapupdetaileventtopicwrapupeventMap["wrapupNotes"].(string); ok {
		o.WrapupNotes = &WrapupNotes
	}
	
	if WrapupDurationMs, ok := WrapupdetaileventtopicwrapupeventMap["wrapupDurationMs"].(float64); ok {
		WrapupDurationMsInt := int(WrapupDurationMs)
		o.WrapupDurationMs = &WrapupDurationMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wrapupdetaileventtopicwrapupevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
