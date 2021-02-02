package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationcallbackeventtopiccallbackmediaparticipant
type Queueconversationcallbackeventtopiccallbackmediaparticipant struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// EndTime
	EndTime *time.Time `json:"endTime,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// Purpose
	Purpose *string `json:"purpose,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// WrapupRequired
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`


	// WrapupPrompt
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`


	// User
	User *Queueconversationcallbackeventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Queueconversationcallbackeventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Queueconversationcallbackeventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationcallbackeventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Queueconversationcallbackeventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Queueconversationcallbackeventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Queueconversationcallbackeventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Queueconversationcallbackeventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Queueconversationcallbackeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Queueconversationcallbackeventtopicjourneycontext `json:"journeyContext,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// OutboundPreview
	OutboundPreview *Queueconversationcallbackeventtopicdialerpreview `json:"outboundPreview,omitempty"`


	// Voicemail
	Voicemail *Queueconversationcallbackeventtopicvoicemail `json:"voicemail,omitempty"`


	// CallbackNumbers
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`


	// CallbackUserName
	CallbackUserName *string `json:"callbackUserName,omitempty"`


	// SkipEnabled
	SkipEnabled *bool `json:"skipEnabled,omitempty"`


	// ExternalCampaign
	ExternalCampaign *bool `json:"externalCampaign,omitempty"`


	// TimeoutSeconds
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`


	// CallbackScheduledTime
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`


	// AutomatedCallbackConfigId
	AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopiccallbackmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
