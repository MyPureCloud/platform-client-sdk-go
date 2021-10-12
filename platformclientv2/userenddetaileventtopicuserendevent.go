package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userenddetaileventtopicuserendevent
type Userenddetaileventtopicuserendevent struct { 
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


	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`


	// QueueId
	QueueId *string `json:"queueId,omitempty"`


	// InteractingDurationMs
	InteractingDurationMs *int `json:"interactingDurationMs,omitempty"`


	// HeldDurationMs
	HeldDurationMs *int `json:"heldDurationMs,omitempty"`


	// AlertingDurationMs
	AlertingDurationMs *int `json:"alertingDurationMs,omitempty"`


	// ContactingDurationMs
	ContactingDurationMs *int `json:"contactingDurationMs,omitempty"`


	// DialingDurationMs
	DialingDurationMs *int `json:"dialingDurationMs,omitempty"`

}

func (o *Userenddetaileventtopicuserendevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userenddetaileventtopicuserendevent
	
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
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackScheduledTime *int `json:"callbackScheduledTime,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		InteractingDurationMs *int `json:"interactingDurationMs,omitempty"`
		
		HeldDurationMs *int `json:"heldDurationMs,omitempty"`
		
		AlertingDurationMs *int `json:"alertingDurationMs,omitempty"`
		
		ContactingDurationMs *int `json:"contactingDurationMs,omitempty"`
		
		DialingDurationMs *int `json:"dialingDurationMs,omitempty"`
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
		
		CallbackUserName: o.CallbackUserName,
		
		CallbackNumbers: o.CallbackNumbers,
		
		CallbackScheduledTime: o.CallbackScheduledTime,
		
		Subject: o.Subject,
		
		MessageType: o.MessageType,
		
		UserId: o.UserId,
		
		DivisionId: o.DivisionId,
		
		QueueId: o.QueueId,
		
		InteractingDurationMs: o.InteractingDurationMs,
		
		HeldDurationMs: o.HeldDurationMs,
		
		AlertingDurationMs: o.AlertingDurationMs,
		
		ContactingDurationMs: o.ContactingDurationMs,
		
		DialingDurationMs: o.DialingDurationMs,
		Alias:    (*Alias)(o),
	})
}

func (o *Userenddetaileventtopicuserendevent) UnmarshalJSON(b []byte) error {
	var UserenddetaileventtopicuserendeventMap map[string]interface{}
	err := json.Unmarshal(b, &UserenddetaileventtopicuserendeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := UserenddetaileventtopicuserendeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := UserenddetaileventtopicuserendeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if ParticipantId, ok := UserenddetaileventtopicuserendeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
	
	if SessionId, ok := UserenddetaileventtopicuserendeventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if DisconnectType, ok := UserenddetaileventtopicuserendeventMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if MediaType, ok := UserenddetaileventtopicuserendeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if Provider, ok := UserenddetaileventtopicuserendeventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if Direction, ok := UserenddetaileventtopicuserendeventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if Ani, ok := UserenddetaileventtopicuserendeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
	
	if Dnis, ok := UserenddetaileventtopicuserendeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
	
	if AddressTo, ok := UserenddetaileventtopicuserendeventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
	
	if AddressFrom, ok := UserenddetaileventtopicuserendeventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
	
	if CallbackUserName, ok := UserenddetaileventtopicuserendeventMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
	
	if CallbackNumbers, ok := UserenddetaileventtopicuserendeventMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackScheduledTime, ok := UserenddetaileventtopicuserendeventMap["callbackScheduledTime"].(float64); ok {
		CallbackScheduledTimeInt := int(CallbackScheduledTime)
		o.CallbackScheduledTime = &CallbackScheduledTimeInt
	}
	
	if Subject, ok := UserenddetaileventtopicuserendeventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
	
	if MessageType, ok := UserenddetaileventtopicuserendeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
	
	if UserId, ok := UserenddetaileventtopicuserendeventMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if DivisionId, ok := UserenddetaileventtopicuserendeventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
	
	if QueueId, ok := UserenddetaileventtopicuserendeventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
	
	if InteractingDurationMs, ok := UserenddetaileventtopicuserendeventMap["interactingDurationMs"].(float64); ok {
		InteractingDurationMsInt := int(InteractingDurationMs)
		o.InteractingDurationMs = &InteractingDurationMsInt
	}
	
	if HeldDurationMs, ok := UserenddetaileventtopicuserendeventMap["heldDurationMs"].(float64); ok {
		HeldDurationMsInt := int(HeldDurationMs)
		o.HeldDurationMs = &HeldDurationMsInt
	}
	
	if AlertingDurationMs, ok := UserenddetaileventtopicuserendeventMap["alertingDurationMs"].(float64); ok {
		AlertingDurationMsInt := int(AlertingDurationMs)
		o.AlertingDurationMs = &AlertingDurationMsInt
	}
	
	if ContactingDurationMs, ok := UserenddetaileventtopicuserendeventMap["contactingDurationMs"].(float64); ok {
		ContactingDurationMsInt := int(ContactingDurationMs)
		o.ContactingDurationMs = &ContactingDurationMsInt
	}
	
	if DialingDurationMs, ok := UserenddetaileventtopicuserendeventMap["dialingDurationMs"].(float64); ok {
		DialingDurationMsInt := int(DialingDurationMs)
		o.DialingDurationMs = &DialingDurationMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userenddetaileventtopicuserendevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
