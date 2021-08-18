package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// Wrapup
	Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Queueconversationvideoeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


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
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Queueconversationvideoeventtopicparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicparticipant

	
	ConnectedTime := new(string)
	if u.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(u.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	EndTime := new(string)
	if u.EndTime != nil {
		
		*EndTime = timeutil.Strftime(u.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
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
		
		Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		ConversationRoutingData *Queueconversationvideoeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`
		
		CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		Calls *[]Queueconversationvideoeventtopiccall `json:"calls,omitempty"`
		
		Callbacks *[]Queueconversationvideoeventtopiccallback `json:"callbacks,omitempty"`
		
		Chats *[]Queueconversationvideoeventtopicchat `json:"chats,omitempty"`
		
		Cobrowsesessions *[]Queueconversationvideoeventtopiccobrowse `json:"cobrowsesessions,omitempty"`
		
		Emails *[]Queueconversationvideoeventtopicemail `json:"emails,omitempty"`
		
		Messages *[]Queueconversationvideoeventtopicmessage `json:"messages,omitempty"`
		
		Screenshares *[]Queueconversationvideoeventtopicscreenshare `json:"screenshares,omitempty"`
		
		SocialExpressions *[]Queueconversationvideoeventtopicsocialexpression `json:"socialExpressions,omitempty"`
		
		Videos *[]Queueconversationvideoeventtopicvideo `json:"videos,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		ConnectedTime: ConnectedTime,
		
		EndTime: EndTime,
		
		UserId: u.UserId,
		
		ExternalContactId: u.ExternalContactId,
		
		ExternalOrganizationId: u.ExternalOrganizationId,
		
		Name: u.Name,
		
		QueueId: u.QueueId,
		
		GroupId: u.GroupId,
		
		TeamId: u.TeamId,
		
		Purpose: u.Purpose,
		
		ConsultParticipantId: u.ConsultParticipantId,
		
		Address: u.Address,
		
		WrapupRequired: u.WrapupRequired,
		
		WrapupExpected: u.WrapupExpected,
		
		WrapupPrompt: u.WrapupPrompt,
		
		WrapupTimeoutMs: u.WrapupTimeoutMs,
		
		Wrapup: u.Wrapup,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		ConversationRoutingData: u.ConversationRoutingData,
		
		AlertingTimeoutMs: u.AlertingTimeoutMs,
		
		MonitoredParticipantId: u.MonitoredParticipantId,
		
		CoachedParticipantId: u.CoachedParticipantId,
		
		ScreenRecordingState: u.ScreenRecordingState,
		
		FlaggedReason: u.FlaggedReason,
		
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
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
