package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Acwdetaileventtopicaftercallworkevent
type Acwdetaileventtopicaftercallworkevent struct { 
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


	// UserId
	UserId *string `json:"userId,omitempty"`


	// QueueId
	QueueId *string `json:"queueId,omitempty"`


	// WrapupCode
	WrapupCode *string `json:"wrapupCode,omitempty"`


	// WrapupNotes
	WrapupNotes *string `json:"wrapupNotes,omitempty"`


	// WrapupDurationMs
	WrapupDurationMs *int `json:"wrapupDurationMs,omitempty"`

}

func (o *Acwdetaileventtopicaftercallworkevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Acwdetaileventtopicaftercallworkevent
	
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
		
		UserId *string `json:"userId,omitempty"`
		
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
		
		UserId: o.UserId,
		
		QueueId: o.QueueId,
		
		WrapupCode: o.WrapupCode,
		
		WrapupNotes: o.WrapupNotes,
		
		WrapupDurationMs: o.WrapupDurationMs,
		Alias:    (*Alias)(o),
	})
}

func (o *Acwdetaileventtopicaftercallworkevent) UnmarshalJSON(b []byte) error {
	var AcwdetaileventtopicaftercallworkeventMap map[string]interface{}
	err := json.Unmarshal(b, &AcwdetaileventtopicaftercallworkeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := AcwdetaileventtopicaftercallworkeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := AcwdetaileventtopicaftercallworkeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := AcwdetaileventtopicaftercallworkeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SessionId, ok := AcwdetaileventtopicaftercallworkeventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if MediaType, ok := AcwdetaileventtopicaftercallworkeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Provider, ok := AcwdetaileventtopicaftercallworkeventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Direction, ok := AcwdetaileventtopicaftercallworkeventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := AcwdetaileventtopicaftercallworkeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := AcwdetaileventtopicaftercallworkeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if AddressTo, ok := AcwdetaileventtopicaftercallworkeventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if AddressFrom, ok := AcwdetaileventtopicaftercallworkeventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if CallbackUserName, ok := AcwdetaileventtopicaftercallworkeventMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
    
	if CallbackNumbers, ok := AcwdetaileventtopicaftercallworkeventMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackScheduledTime, ok := AcwdetaileventtopicaftercallworkeventMap["callbackScheduledTime"].(float64); ok {
		CallbackScheduledTimeInt := int(CallbackScheduledTime)
		o.CallbackScheduledTime = &CallbackScheduledTimeInt
	}
	
	if Subject, ok := AcwdetaileventtopicaftercallworkeventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if MessageType, ok := AcwdetaileventtopicaftercallworkeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if UserId, ok := AcwdetaileventtopicaftercallworkeventMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if QueueId, ok := AcwdetaileventtopicaftercallworkeventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if WrapupCode, ok := AcwdetaileventtopicaftercallworkeventMap["wrapupCode"].(string); ok {
		o.WrapupCode = &WrapupCode
	}
    
	if WrapupNotes, ok := AcwdetaileventtopicaftercallworkeventMap["wrapupNotes"].(string); ok {
		o.WrapupNotes = &WrapupNotes
	}
    
	if WrapupDurationMs, ok := AcwdetaileventtopicaftercallworkeventMap["wrapupDurationMs"].(float64); ok {
		WrapupDurationMsInt := int(WrapupDurationMs)
		o.WrapupDurationMs = &WrapupDurationMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Acwdetaileventtopicaftercallworkevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
