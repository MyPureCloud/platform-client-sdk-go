package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Participant
type Participant struct { 
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
	Calls *[]Call `json:"calls,omitempty"`


	// Callbacks
	Callbacks *[]Callback `json:"callbacks,omitempty"`


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

func (u *Participant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Participant

	
	StartTime := new(string)
	if u.StartTime != nil {
		
		*StartTime = timeutil.Strftime(u.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	EndTime := new(string)
	if u.EndTime != nil {
		
		*EndTime = timeutil.Strftime(u.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	ConnectedTime := new(string)
	if u.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(u.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	StartAcwTime := new(string)
	if u.StartAcwTime != nil {
		
		*StartAcwTime = timeutil.Strftime(u.StartAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAcwTime = nil
	}
	
	EndAcwTime := new(string)
	if u.EndAcwTime != nil {
		
		*EndAcwTime = timeutil.Strftime(u.EndAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		Calls *[]Call `json:"calls,omitempty"`
		
		Callbacks *[]Callback `json:"callbacks,omitempty"`
		
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
		Id: u.Id,
		
		StartTime: StartTime,
		
		EndTime: EndTime,
		
		ConnectedTime: ConnectedTime,
		
		Name: u.Name,
		
		UserUri: u.UserUri,
		
		UserId: u.UserId,
		
		ExternalContactId: u.ExternalContactId,
		
		ExternalOrganizationId: u.ExternalOrganizationId,
		
		QueueId: u.QueueId,
		
		GroupId: u.GroupId,
		
		TeamId: u.TeamId,
		
		QueueName: u.QueueName,
		
		Purpose: u.Purpose,
		
		ParticipantType: u.ParticipantType,
		
		ConsultParticipantId: u.ConsultParticipantId,
		
		Address: u.Address,
		
		Ani: u.Ani,
		
		AniName: u.AniName,
		
		Dnis: u.Dnis,
		
		Locale: u.Locale,
		
		WrapupRequired: u.WrapupRequired,
		
		WrapupPrompt: u.WrapupPrompt,
		
		WrapupTimeoutMs: u.WrapupTimeoutMs,
		
		WrapupSkipped: u.WrapupSkipped,
		
		Wrapup: u.Wrapup,
		
		ConversationRoutingData: u.ConversationRoutingData,
		
		AlertingTimeoutMs: u.AlertingTimeoutMs,
		
		MonitoredParticipantId: u.MonitoredParticipantId,
		
		CoachedParticipantId: u.CoachedParticipantId,
		
		Attributes: u.Attributes,
		
		Calls: u.Calls,
		
		Callbacks: u.Callbacks,
		
		Chats: u.Chats,
		
		Cobrowsesessions: u.Cobrowsesessions,
		
		Emails: u.Emails,
		
		Messages: u.Messages,
		
		Screenshares: u.Screenshares,
		
		SocialExpressions: u.SocialExpressions,
		
		Videos: u.Videos,
		
		Evaluations: u.Evaluations,
		
		ScreenRecordingState: u.ScreenRecordingState,
		
		FlaggedReason: u.FlaggedReason,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Participant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
