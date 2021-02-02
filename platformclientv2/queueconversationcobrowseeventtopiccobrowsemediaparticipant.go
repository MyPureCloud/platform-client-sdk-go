package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationcobrowseeventtopiccobrowsemediaparticipant
type Queueconversationcobrowseeventtopiccobrowsemediaparticipant struct { 
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
	User *Queueconversationcobrowseeventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Queueconversationcobrowseeventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Queueconversationcobrowseeventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationcobrowseeventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Queueconversationcobrowseeventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Queueconversationcobrowseeventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Queueconversationcobrowseeventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Queueconversationcobrowseeventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Queueconversationcobrowseeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Queueconversationcobrowseeventtopicjourneycontext `json:"journeyContext,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// CobrowseSessionId
	CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`


	// CobrowseRole
	CobrowseRole *string `json:"cobrowseRole,omitempty"`


	// ViewerUrl
	ViewerUrl *string `json:"viewerUrl,omitempty"`


	// ProviderEventTime
	ProviderEventTime *time.Time `json:"providerEventTime,omitempty"`


	// Controlling
	Controlling *[]string `json:"controlling,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopiccobrowsemediaparticipant) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
