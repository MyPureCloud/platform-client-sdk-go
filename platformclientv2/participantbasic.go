package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Participantbasic
type Participantbasic struct { 
	// Id - A globally unique identifier for this conversation.
	Id *string `json:"id,omitempty"`


	// StartTime - The timestamp when this participant joined the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The timestamp when this participant disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// ConnectedTime - The timestamp when this participant was connected to the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// Name - A human readable name identifying the participant.
	Name *string `json:"name,omitempty"`


	// UserUri - If this participant represents a user, then this will be an URI that can be used to fetch the user.
	UserUri *string `json:"userUri,omitempty"`


	// UserId - If this participant represents a user, then this will be the globally unique identifier for the user.
	UserId *string `json:"userId,omitempty"`


	// ExternalContactId - If this participant represents an external contact, then this will be the globally unique identifier for the external contact.
	ExternalContactId *string `json:"externalContactId,omitempty"`


	// ExternalOrganizationId - If this participant represents an external org, then this will be the globally unique identifier for the external org.
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`


	// QueueId - If present, the queue id that the communication channel came in on.
	QueueId *string `json:"queueId,omitempty"`


	// GroupId - If present, group of users the participant represents.
	GroupId *string `json:"groupId,omitempty"`


	// TeamId - The team id that this participant is a member of when added to the conversation.
	TeamId *string `json:"teamId,omitempty"`


	// QueueName - If present, the queue name that the communication channel came in on.
	QueueName *string `json:"queueName,omitempty"`


	// Purpose - A well known string that specifies the purpose of this participant.
	Purpose *string `json:"purpose,omitempty"`


	// ParticipantType - A well known string that specifies the type of this participant.
	ParticipantType *string `json:"participantType,omitempty"`


	// ConsultParticipantId - If this participant is part of a consult transfer, then this will be the participant id of the participant being transferred.
	ConsultParticipantId *string `json:"consultParticipantId,omitempty"`


	// Address - The address for the this participant. For a phone call this will be the ANI.
	Address *string `json:"address,omitempty"`


	// Ani - The address for the this participant. For a phone call this will be the ANI.
	Ani *string `json:"ani,omitempty"`


	// AniName - The ani-based name for this participant.
	AniName *string `json:"aniName,omitempty"`


	// Dnis - The address for the this participant. For a phone call this will be the ANI.
	Dnis *string `json:"dnis,omitempty"`


	// Locale - An ISO 639 language code specifying the locale for this participant
	Locale *string `json:"locale,omitempty"`


	// WrapupRequired - True iff this participant is required to enter wrapup for this conversation.
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`


	// WrapupPrompt - This field controls how the UI prompts the agent for a wrapup.
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`


	// WrapupTimeoutMs - Specifies how long a timed ACW session will last.
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped - The UI sets this field when the agent chooses to skip entering a wrapup for this participant.
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData - Information on how a communication should be routed to an agent.
	ConversationRoutingData *Conversationroutingdata `json:"conversationRoutingData,omitempty"`


	// AlertingTimeoutMs - Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// MonitoredParticipantId - If this participant is a monitor, then this will be the id of the participant that is being monitored.
	MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`


	// CoachedParticipantId - If this participant is a coach, then this will be the id of the participant that is being coached.
	CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`


	// Attributes - Additional participant attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// Calls
	Calls *[]Callbasic `json:"calls,omitempty"`


	// Callbacks
	Callbacks *[]Callbackbasic `json:"callbacks,omitempty"`


	// Chats
	Chats *[]Conversationchat `json:"chats,omitempty"`


	// Cobrowsesessions
	Cobrowsesessions *[]Cobrowsesession `json:"cobrowsesessions,omitempty"`


	// Emails
	Emails *[]Email `json:"emails,omitempty"`


	// Messages
	Messages *[]Message `json:"messages,omitempty"`


	// Screenshares
	Screenshares *[]Screenshare `json:"screenshares,omitempty"`


	// SocialExpressions
	SocialExpressions *[]Socialexpression `json:"socialExpressions,omitempty"`


	// Videos
	Videos *[]Video `json:"videos,omitempty"`


	// Evaluations
	Evaluations *[]Evaluation `json:"evaluations,omitempty"`


	// ScreenRecordingState - The current screen recording state for this participant.
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason - The reason specifying why participant flagged the conversation.
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// StartAcwTime - The timestamp when this participant started after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime - The timestamp when this participant ended after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`

}

func (o *Participantbasic) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Participantbasic
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
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
		
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UserUri *string `json:"userUri,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		GroupId *string `json:"groupId,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		QueueName *string `json:"queueName,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		ParticipantType *string `json:"participantType,omitempty"`
		
		ConsultParticipantId *string `json:"consultParticipantId,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		AniName *string `json:"aniName,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		Locale *string `json:"locale,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		Wrapup *Wrapup `json:"wrapup,omitempty"`
		
		ConversationRoutingData *Conversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`
		
		CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		Calls *[]Callbasic `json:"calls,omitempty"`
		
		Callbacks *[]Callbackbasic `json:"callbacks,omitempty"`
		
		Chats *[]Conversationchat `json:"chats,omitempty"`
		
		Cobrowsesessions *[]Cobrowsesession `json:"cobrowsesessions,omitempty"`
		
		Emails *[]Email `json:"emails,omitempty"`
		
		Messages *[]Message `json:"messages,omitempty"`
		
		Screenshares *[]Screenshare `json:"screenshares,omitempty"`
		
		SocialExpressions *[]Socialexpression `json:"socialExpressions,omitempty"`
		
		Videos *[]Video `json:"videos,omitempty"`
		
		Evaluations *[]Evaluation `json:"evaluations,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		StartTime: StartTime,
		
		EndTime: EndTime,
		
		ConnectedTime: ConnectedTime,
		
		Name: o.Name,
		
		UserUri: o.UserUri,
		
		UserId: o.UserId,
		
		ExternalContactId: o.ExternalContactId,
		
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		QueueId: o.QueueId,
		
		GroupId: o.GroupId,
		
		TeamId: o.TeamId,
		
		QueueName: o.QueueName,
		
		Purpose: o.Purpose,
		
		ParticipantType: o.ParticipantType,
		
		ConsultParticipantId: o.ConsultParticipantId,
		
		Address: o.Address,
		
		Ani: o.Ani,
		
		AniName: o.AniName,
		
		Dnis: o.Dnis,
		
		Locale: o.Locale,
		
		WrapupRequired: o.WrapupRequired,
		
		WrapupPrompt: o.WrapupPrompt,
		
		WrapupTimeoutMs: o.WrapupTimeoutMs,
		
		WrapupSkipped: o.WrapupSkipped,
		
		Wrapup: o.Wrapup,
		
		ConversationRoutingData: o.ConversationRoutingData,
		
		AlertingTimeoutMs: o.AlertingTimeoutMs,
		
		MonitoredParticipantId: o.MonitoredParticipantId,
		
		CoachedParticipantId: o.CoachedParticipantId,
		
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
		
		Evaluations: o.Evaluations,
		
		ScreenRecordingState: o.ScreenRecordingState,
		
		FlaggedReason: o.FlaggedReason,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Participantbasic) UnmarshalJSON(b []byte) error {
	var ParticipantbasicMap map[string]interface{}
	err := json.Unmarshal(b, &ParticipantbasicMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ParticipantbasicMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if startTimeString, ok := ParticipantbasicMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if endTimeString, ok := ParticipantbasicMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if connectedTimeString, ok := ParticipantbasicMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if Name, ok := ParticipantbasicMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if UserUri, ok := ParticipantbasicMap["userUri"].(string); ok {
		o.UserUri = &UserUri
	}
	
	if UserId, ok := ParticipantbasicMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if ExternalContactId, ok := ParticipantbasicMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
	
	if ExternalOrganizationId, ok := ParticipantbasicMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
	
	if QueueId, ok := ParticipantbasicMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
	
	if GroupId, ok := ParticipantbasicMap["groupId"].(string); ok {
		o.GroupId = &GroupId
	}
	
	if TeamId, ok := ParticipantbasicMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
	
	if QueueName, ok := ParticipantbasicMap["queueName"].(string); ok {
		o.QueueName = &QueueName
	}
	
	if Purpose, ok := ParticipantbasicMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
	
	if ParticipantType, ok := ParticipantbasicMap["participantType"].(string); ok {
		o.ParticipantType = &ParticipantType
	}
	
	if ConsultParticipantId, ok := ParticipantbasicMap["consultParticipantId"].(string); ok {
		o.ConsultParticipantId = &ConsultParticipantId
	}
	
	if Address, ok := ParticipantbasicMap["address"].(string); ok {
		o.Address = &Address
	}
	
	if Ani, ok := ParticipantbasicMap["ani"].(string); ok {
		o.Ani = &Ani
	}
	
	if AniName, ok := ParticipantbasicMap["aniName"].(string); ok {
		o.AniName = &AniName
	}
	
	if Dnis, ok := ParticipantbasicMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
	
	if Locale, ok := ParticipantbasicMap["locale"].(string); ok {
		o.Locale = &Locale
	}
	
	if WrapupRequired, ok := ParticipantbasicMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
	
	if WrapupPrompt, ok := ParticipantbasicMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
	
	if WrapupTimeoutMs, ok := ParticipantbasicMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if WrapupSkipped, ok := ParticipantbasicMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
	
	if Wrapup, ok := ParticipantbasicMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if ConversationRoutingData, ok := ParticipantbasicMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if AlertingTimeoutMs, ok := ParticipantbasicMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if MonitoredParticipantId, ok := ParticipantbasicMap["monitoredParticipantId"].(string); ok {
		o.MonitoredParticipantId = &MonitoredParticipantId
	}
	
	if CoachedParticipantId, ok := ParticipantbasicMap["coachedParticipantId"].(string); ok {
		o.CoachedParticipantId = &CoachedParticipantId
	}
	
	if Attributes, ok := ParticipantbasicMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Calls, ok := ParticipantbasicMap["calls"].([]interface{}); ok {
		CallsString, _ := json.Marshal(Calls)
		json.Unmarshal(CallsString, &o.Calls)
	}
	
	if Callbacks, ok := ParticipantbasicMap["callbacks"].([]interface{}); ok {
		CallbacksString, _ := json.Marshal(Callbacks)
		json.Unmarshal(CallbacksString, &o.Callbacks)
	}
	
	if Chats, ok := ParticipantbasicMap["chats"].([]interface{}); ok {
		ChatsString, _ := json.Marshal(Chats)
		json.Unmarshal(ChatsString, &o.Chats)
	}
	
	if Cobrowsesessions, ok := ParticipantbasicMap["cobrowsesessions"].([]interface{}); ok {
		CobrowsesessionsString, _ := json.Marshal(Cobrowsesessions)
		json.Unmarshal(CobrowsesessionsString, &o.Cobrowsesessions)
	}
	
	if Emails, ok := ParticipantbasicMap["emails"].([]interface{}); ok {
		EmailsString, _ := json.Marshal(Emails)
		json.Unmarshal(EmailsString, &o.Emails)
	}
	
	if Messages, ok := ParticipantbasicMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	
	if Screenshares, ok := ParticipantbasicMap["screenshares"].([]interface{}); ok {
		ScreensharesString, _ := json.Marshal(Screenshares)
		json.Unmarshal(ScreensharesString, &o.Screenshares)
	}
	
	if SocialExpressions, ok := ParticipantbasicMap["socialExpressions"].([]interface{}); ok {
		SocialExpressionsString, _ := json.Marshal(SocialExpressions)
		json.Unmarshal(SocialExpressionsString, &o.SocialExpressions)
	}
	
	if Videos, ok := ParticipantbasicMap["videos"].([]interface{}); ok {
		VideosString, _ := json.Marshal(Videos)
		json.Unmarshal(VideosString, &o.Videos)
	}
	
	if Evaluations, ok := ParticipantbasicMap["evaluations"].([]interface{}); ok {
		EvaluationsString, _ := json.Marshal(Evaluations)
		json.Unmarshal(EvaluationsString, &o.Evaluations)
	}
	
	if ScreenRecordingState, ok := ParticipantbasicMap["screenRecordingState"].(string); ok {
		o.ScreenRecordingState = &ScreenRecordingState
	}
	
	if FlaggedReason, ok := ParticipantbasicMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
	
	if startAcwTimeString, ok := ParticipantbasicMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := ParticipantbasicMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Participantbasic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
