package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicparticipant
type Conversationeventtopicparticipant struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// EndTime
	EndTime *time.Time `json:"endTime,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// ExternalContactId
	ExternalContactId *string `json:"externalContactId,omitempty"`


	// ExternalOrganizationId
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// QueueId
	QueueId *string `json:"queueId,omitempty"`


	// GroupId
	GroupId *string `json:"groupId,omitempty"`


	// TeamId
	TeamId *string `json:"teamId,omitempty"`


	// Purpose
	Purpose *string `json:"purpose,omitempty"`


	// ConsultParticipantId
	ConsultParticipantId *string `json:"consultParticipantId,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`


	// WrapupRequired
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`


	// WrapupExpected
	WrapupExpected *bool `json:"wrapupExpected,omitempty"`


	// WrapupPrompt
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// Wrapup
	Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Conversationeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// MonitoredParticipantId
	MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`


	// CoachedParticipantId
	CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// Calls
	Calls *[]Conversationeventtopiccall `json:"calls,omitempty"`


	// Callbacks
	Callbacks *[]Conversationeventtopiccallback `json:"callbacks,omitempty"`


	// Chats
	Chats *[]Conversationeventtopicchat `json:"chats,omitempty"`


	// Cobrowsesessions
	Cobrowsesessions *[]Conversationeventtopiccobrowse `json:"cobrowsesessions,omitempty"`


	// Emails
	Emails *[]Conversationeventtopicemail `json:"emails,omitempty"`


	// Messages
	Messages *[]Conversationeventtopicmessage `json:"messages,omitempty"`


	// Screenshares
	Screenshares *[]Conversationeventtopicscreenshare `json:"screenshares,omitempty"`


	// SocialExpressions
	SocialExpressions *[]Conversationeventtopicsocialexpression `json:"socialExpressions,omitempty"`


	// Videos
	Videos *[]Conversationeventtopicvideo `json:"videos,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Conversationeventtopicparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicparticipant
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	StartAcwTime := new(string)
	if o.StartAcwTime != nil {
		
		*StartAcwTime = timeutil.Strftime(o.StartAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAcwTime = nil
	}
	
	EndAcwTime := new(string)
	if o.EndAcwTime != nil {
		
		*EndAcwTime = timeutil.Strftime(o.EndAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndAcwTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		GroupId *string `json:"groupId,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		ConsultParticipantId *string `json:"consultParticipantId,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupExpected *bool `json:"wrapupExpected,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		ConversationRoutingData *Conversationeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`
		
		CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		Calls *[]Conversationeventtopiccall `json:"calls,omitempty"`
		
		Callbacks *[]Conversationeventtopiccallback `json:"callbacks,omitempty"`
		
		Chats *[]Conversationeventtopicchat `json:"chats,omitempty"`
		
		Cobrowsesessions *[]Conversationeventtopiccobrowse `json:"cobrowsesessions,omitempty"`
		
		Emails *[]Conversationeventtopicemail `json:"emails,omitempty"`
		
		Messages *[]Conversationeventtopicmessage `json:"messages,omitempty"`
		
		Screenshares *[]Conversationeventtopicscreenshare `json:"screenshares,omitempty"`
		
		SocialExpressions *[]Conversationeventtopicsocialexpression `json:"socialExpressions,omitempty"`
		
		Videos *[]Conversationeventtopicvideo `json:"videos,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ConnectedTime: ConnectedTime,
		
		EndTime: EndTime,
		
		UserId: o.UserId,
		
		ExternalContactId: o.ExternalContactId,
		
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		Name: o.Name,
		
		QueueId: o.QueueId,
		
		GroupId: o.GroupId,
		
		TeamId: o.TeamId,
		
		Purpose: o.Purpose,
		
		ConsultParticipantId: o.ConsultParticipantId,
		
		Address: o.Address,
		
		WrapupRequired: o.WrapupRequired,
		
		WrapupExpected: o.WrapupExpected,
		
		WrapupPrompt: o.WrapupPrompt,
		
		WrapupTimeoutMs: o.WrapupTimeoutMs,
		
		Wrapup: o.Wrapup,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		ConversationRoutingData: o.ConversationRoutingData,
		
		AlertingTimeoutMs: o.AlertingTimeoutMs,
		
		MonitoredParticipantId: o.MonitoredParticipantId,
		
		CoachedParticipantId: o.CoachedParticipantId,
		
		ScreenRecordingState: o.ScreenRecordingState,
		
		FlaggedReason: o.FlaggedReason,
		
		Attributes: o.Attributes,
		
		Calls: o.Calls,
		
		Callbacks: o.Callbacks,
		
		Chats: o.Chats,
		
		Cobrowsesessions: o.Cobrowsesessions,
		
		Emails: o.Emails,
		
		Messages: o.Messages,
		
		Screenshares: o.Screenshares,
		
		SocialExpressions: o.SocialExpressions,
		
		Videos: o.Videos,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventtopicparticipant) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationeventtopicparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if connectedTimeString, ok := ConversationeventtopicparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := ConversationeventtopicparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if UserId, ok := ConversationeventtopicparticipantMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if ExternalContactId, ok := ConversationeventtopicparticipantMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
	
	if ExternalOrganizationId, ok := ConversationeventtopicparticipantMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
	
	if Name, ok := ConversationeventtopicparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if QueueId, ok := ConversationeventtopicparticipantMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
	
	if GroupId, ok := ConversationeventtopicparticipantMap["groupId"].(string); ok {
		o.GroupId = &GroupId
	}
	
	if TeamId, ok := ConversationeventtopicparticipantMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
	
	if Purpose, ok := ConversationeventtopicparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
	
	if ConsultParticipantId, ok := ConversationeventtopicparticipantMap["consultParticipantId"].(string); ok {
		o.ConsultParticipantId = &ConsultParticipantId
	}
	
	if Address, ok := ConversationeventtopicparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
	
	if WrapupRequired, ok := ConversationeventtopicparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
	
	if WrapupExpected, ok := ConversationeventtopicparticipantMap["wrapupExpected"].(bool); ok {
		o.WrapupExpected = &WrapupExpected
	}
	
	if WrapupPrompt, ok := ConversationeventtopicparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
	
	if WrapupTimeoutMs, ok := ConversationeventtopicparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if Wrapup, ok := ConversationeventtopicparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if startAcwTimeString, ok := ConversationeventtopicparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := ConversationeventtopicparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if ConversationRoutingData, ok := ConversationeventtopicparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if AlertingTimeoutMs, ok := ConversationeventtopicparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if MonitoredParticipantId, ok := ConversationeventtopicparticipantMap["monitoredParticipantId"].(string); ok {
		o.MonitoredParticipantId = &MonitoredParticipantId
	}
	
	if CoachedParticipantId, ok := ConversationeventtopicparticipantMap["coachedParticipantId"].(string); ok {
		o.CoachedParticipantId = &CoachedParticipantId
	}
	
	if ScreenRecordingState, ok := ConversationeventtopicparticipantMap["screenRecordingState"].(string); ok {
		o.ScreenRecordingState = &ScreenRecordingState
	}
	
	if FlaggedReason, ok := ConversationeventtopicparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
	
	if Attributes, ok := ConversationeventtopicparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Calls, ok := ConversationeventtopicparticipantMap["calls"].([]interface{}); ok {
		CallsString, _ := json.Marshal(Calls)
		json.Unmarshal(CallsString, &o.Calls)
	}
	
	if Callbacks, ok := ConversationeventtopicparticipantMap["callbacks"].([]interface{}); ok {
		CallbacksString, _ := json.Marshal(Callbacks)
		json.Unmarshal(CallbacksString, &o.Callbacks)
	}
	
	if Chats, ok := ConversationeventtopicparticipantMap["chats"].([]interface{}); ok {
		ChatsString, _ := json.Marshal(Chats)
		json.Unmarshal(ChatsString, &o.Chats)
	}
	
	if Cobrowsesessions, ok := ConversationeventtopicparticipantMap["cobrowsesessions"].([]interface{}); ok {
		CobrowsesessionsString, _ := json.Marshal(Cobrowsesessions)
		json.Unmarshal(CobrowsesessionsString, &o.Cobrowsesessions)
	}
	
	if Emails, ok := ConversationeventtopicparticipantMap["emails"].([]interface{}); ok {
		EmailsString, _ := json.Marshal(Emails)
		json.Unmarshal(EmailsString, &o.Emails)
	}
	
	if Messages, ok := ConversationeventtopicparticipantMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	
	if Screenshares, ok := ConversationeventtopicparticipantMap["screenshares"].([]interface{}); ok {
		ScreensharesString, _ := json.Marshal(Screenshares)
		json.Unmarshal(ScreensharesString, &o.Screenshares)
	}
	
	if SocialExpressions, ok := ConversationeventtopicparticipantMap["socialExpressions"].([]interface{}); ok {
		SocialExpressionsString, _ := json.Marshal(SocialExpressions)
		json.Unmarshal(SocialExpressionsString, &o.SocialExpressions)
	}
	
	if Videos, ok := ConversationeventtopicparticipantMap["videos"].([]interface{}); ok {
		VideosString, _ := json.Marshal(Videos)
		json.Unmarshal(VideosString, &o.Videos)
	}
	
	if AdditionalProperties, ok := ConversationeventtopicparticipantMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
