package platformclientv2
import (
	"time"
	"encoding/json"
)

// Participant
type Participant struct { 
	// Id - A globally unique identifier for this conversation.
	Id *string `json:"id,omitempty"`


	// StartTime - The timestamp when this participant joined the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The timestamp when this participant disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndTime *time.Time `json:"endTime,omitempty"`


	// ConnectedTime - The timestamp when this participant was connected to the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
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
	WrapupTimeoutMs *int32 `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped - The UI sets this field when the agent chooses to skip entering a wrapup for this participant.
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData - Information on how a communication should be routed to an agent.
	ConversationRoutingData *Conversationroutingdata `json:"conversationRoutingData,omitempty"`


	// AlertingTimeoutMs - Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs *int32 `json:"alertingTimeoutMs,omitempty"`


	// MonitoredParticipantId - If this participant is a monitor, then this will be the id of the participant that is being monitored.
	MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`


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


	// StartAcwTime - The timestamp when this participant started after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime - The timestamp when this participant ended after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Participant) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
