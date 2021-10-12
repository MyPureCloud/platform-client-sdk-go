package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Acdenddetaileventtopicacdendevent
type Acdenddetaileventtopicacdendevent struct { 
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


	// QueueId
	QueueId *string `json:"queueId,omitempty"`


	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`


	// AcdOutcome
	AcdOutcome *string `json:"acdOutcome,omitempty"`


	// AnsweredUserId
	AnsweredUserId *string `json:"answeredUserId,omitempty"`


	// RequestedRoutings
	RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`


	// UsedRouting
	UsedRouting *string `json:"usedRouting,omitempty"`


	// RequestedRoutingSkillIds
	RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`


	// RequestedLanguageId
	RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`


	// RequestedRoutingUserIds
	RequestedRoutingUserIds *[]string `json:"requestedRoutingUserIds,omitempty"`


	// RoutingPriority
	RoutingPriority *int `json:"routingPriority,omitempty"`


	// ConnectedDurationMs
	ConnectedDurationMs *int `json:"connectedDurationMs,omitempty"`

}

func (o *Acdenddetaileventtopicacdendevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Acdenddetaileventtopicacdendevent
	
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
		
		QueueId *string `json:"queueId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		AcdOutcome *string `json:"acdOutcome,omitempty"`
		
		AnsweredUserId *string `json:"answeredUserId,omitempty"`
		
		RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`
		
		UsedRouting *string `json:"usedRouting,omitempty"`
		
		RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`
		
		RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`
		
		RequestedRoutingUserIds *[]string `json:"requestedRoutingUserIds,omitempty"`
		
		RoutingPriority *int `json:"routingPriority,omitempty"`
		
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
		
		CallbackUserName: o.CallbackUserName,
		
		CallbackNumbers: o.CallbackNumbers,
		
		CallbackScheduledTime: o.CallbackScheduledTime,
		
		Subject: o.Subject,
		
		MessageType: o.MessageType,
		
		QueueId: o.QueueId,
		
		DivisionId: o.DivisionId,
		
		AcdOutcome: o.AcdOutcome,
		
		AnsweredUserId: o.AnsweredUserId,
		
		RequestedRoutings: o.RequestedRoutings,
		
		UsedRouting: o.UsedRouting,
		
		RequestedRoutingSkillIds: o.RequestedRoutingSkillIds,
		
		RequestedLanguageId: o.RequestedLanguageId,
		
		RequestedRoutingUserIds: o.RequestedRoutingUserIds,
		
		RoutingPriority: o.RoutingPriority,
		
		ConnectedDurationMs: o.ConnectedDurationMs,
		Alias:    (*Alias)(o),
	})
}

func (o *Acdenddetaileventtopicacdendevent) UnmarshalJSON(b []byte) error {
	var AcdenddetaileventtopicacdendeventMap map[string]interface{}
	err := json.Unmarshal(b, &AcdenddetaileventtopicacdendeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := AcdenddetaileventtopicacdendeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := AcdenddetaileventtopicacdendeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if ParticipantId, ok := AcdenddetaileventtopicacdendeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
	
	if SessionId, ok := AcdenddetaileventtopicacdendeventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if DisconnectType, ok := AcdenddetaileventtopicacdendeventMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if MediaType, ok := AcdenddetaileventtopicacdendeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if Provider, ok := AcdenddetaileventtopicacdendeventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if Direction, ok := AcdenddetaileventtopicacdendeventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if Ani, ok := AcdenddetaileventtopicacdendeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
	
	if Dnis, ok := AcdenddetaileventtopicacdendeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
	
	if AddressTo, ok := AcdenddetaileventtopicacdendeventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
	
	if AddressFrom, ok := AcdenddetaileventtopicacdendeventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
	
	if CallbackUserName, ok := AcdenddetaileventtopicacdendeventMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
	
	if CallbackNumbers, ok := AcdenddetaileventtopicacdendeventMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackScheduledTime, ok := AcdenddetaileventtopicacdendeventMap["callbackScheduledTime"].(float64); ok {
		CallbackScheduledTimeInt := int(CallbackScheduledTime)
		o.CallbackScheduledTime = &CallbackScheduledTimeInt
	}
	
	if Subject, ok := AcdenddetaileventtopicacdendeventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
	
	if MessageType, ok := AcdenddetaileventtopicacdendeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
	
	if QueueId, ok := AcdenddetaileventtopicacdendeventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
	
	if DivisionId, ok := AcdenddetaileventtopicacdendeventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
	
	if AcdOutcome, ok := AcdenddetaileventtopicacdendeventMap["acdOutcome"].(string); ok {
		o.AcdOutcome = &AcdOutcome
	}
	
	if AnsweredUserId, ok := AcdenddetaileventtopicacdendeventMap["answeredUserId"].(string); ok {
		o.AnsweredUserId = &AnsweredUserId
	}
	
	if RequestedRoutings, ok := AcdenddetaileventtopicacdendeventMap["requestedRoutings"].([]interface{}); ok {
		RequestedRoutingsString, _ := json.Marshal(RequestedRoutings)
		json.Unmarshal(RequestedRoutingsString, &o.RequestedRoutings)
	}
	
	if UsedRouting, ok := AcdenddetaileventtopicacdendeventMap["usedRouting"].(string); ok {
		o.UsedRouting = &UsedRouting
	}
	
	if RequestedRoutingSkillIds, ok := AcdenddetaileventtopicacdendeventMap["requestedRoutingSkillIds"].([]interface{}); ok {
		RequestedRoutingSkillIdsString, _ := json.Marshal(RequestedRoutingSkillIds)
		json.Unmarshal(RequestedRoutingSkillIdsString, &o.RequestedRoutingSkillIds)
	}
	
	if RequestedLanguageId, ok := AcdenddetaileventtopicacdendeventMap["requestedLanguageId"].(string); ok {
		o.RequestedLanguageId = &RequestedLanguageId
	}
	
	if RequestedRoutingUserIds, ok := AcdenddetaileventtopicacdendeventMap["requestedRoutingUserIds"].([]interface{}); ok {
		RequestedRoutingUserIdsString, _ := json.Marshal(RequestedRoutingUserIds)
		json.Unmarshal(RequestedRoutingUserIdsString, &o.RequestedRoutingUserIds)
	}
	
	if RoutingPriority, ok := AcdenddetaileventtopicacdendeventMap["routingPriority"].(float64); ok {
		RoutingPriorityInt := int(RoutingPriority)
		o.RoutingPriority = &RoutingPriorityInt
	}
	
	if ConnectedDurationMs, ok := AcdenddetaileventtopicacdendeventMap["connectedDurationMs"].(float64); ok {
		ConnectedDurationMsInt := int(ConnectedDurationMs)
		o.ConnectedDurationMs = &ConnectedDurationMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Acdenddetaileventtopicacdendevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
