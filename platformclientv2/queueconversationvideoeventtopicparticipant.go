package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationvideoeventtopicparticipant
type Queueconversationvideoeventtopicparticipant struct { 
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
	WrapupTimeoutMs *int32 `json:"wrapupTimeoutMs,omitempty"`


	// Wrapup
	Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Queueconversationvideoeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int32 `json:"alertingTimeoutMs,omitempty"`


	// MonitoredParticipantId
	MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// Calls
	Calls *[]Queueconversationvideoeventtopiccall `json:"calls,omitempty"`


	// Callbacks
	Callbacks *[]Queueconversationvideoeventtopiccallback `json:"callbacks,omitempty"`


	// Chats
	Chats *[]Queueconversationvideoeventtopicchat `json:"chats,omitempty"`


	// Cobrowsesessions
	Cobrowsesessions *[]Queueconversationvideoeventtopiccobrowse `json:"cobrowsesessions,omitempty"`


	// Emails
	Emails *[]Queueconversationvideoeventtopicemail `json:"emails,omitempty"`


	// Messages
	Messages *[]Queueconversationvideoeventtopicmessage `json:"messages,omitempty"`


	// Screenshares
	Screenshares *[]Queueconversationvideoeventtopicscreenshare `json:"screenshares,omitempty"`


	// SocialExpressions
	SocialExpressions *[]Queueconversationvideoeventtopicsocialexpression `json:"socialExpressions,omitempty"`


	// Videos
	Videos *[]Queueconversationvideoeventtopicvideo `json:"videos,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicparticipant) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
