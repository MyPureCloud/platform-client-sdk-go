package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationscreenshareeventtopicscreensharemediaparticipant
type Queueconversationscreenshareeventtopicscreensharemediaparticipant struct { 
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
	User *Queueconversationscreenshareeventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Queueconversationscreenshareeventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Queueconversationscreenshareeventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationscreenshareeventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Queueconversationscreenshareeventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int32 `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int32 `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Queueconversationscreenshareeventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Queueconversationscreenshareeventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Queueconversationscreenshareeventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Queueconversationscreenshareeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Queueconversationscreenshareeventtopicjourneycontext `json:"journeyContext,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// Context
	Context *string `json:"context,omitempty"`


	// PeerCount
	PeerCount *int32 `json:"peerCount,omitempty"`


	// Sharing
	Sharing *bool `json:"sharing,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicscreensharemediaparticipant) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
